package bookstore

import (
	"math"
	"sort"
)

// The cost of single book
// in cents
const bookCost = 800

// discount of the books based on
// number of unique books in
// the group
var countDiscount = map[int]float64{
	1: 1.0,
	2: 0.95,
	3: 0.9,
	4: 0.8,
	5: 0.75,
}

// Cost calculates the best possible
// discount for a given basket of
// books
func Cost(basket []int) int {
	// this makes grouping
	// easier/possible
	sort.Ints(basket)

	basketCount := len(basket)
	discount := countDiscount[1] // 1.0

	// this first check is an optimization
	// since only baskets with 5 or less
	// items can all be unique, we can do
	// a much simpler check first
	if basketCount < 6 && allUnique(basket) {
		discount = countDiscount[basketCount]
	} else {
		// If we ever come up with other
		// strategies to try, add them
		// here
		discountGroups := [][]grouping{
			broadGrouping(basket),
			deepGrouping(basket),
		}

		// try all the strategies to see which
		// gives the largest discount
		for _, groups := range discountGroups {
			tmpDiscount := 0.0
			for _, group := range groups {
				// calculate the discount per group of books.
				// Note the discount is per book, not per group
				// so we have to multiple the dicount times
				// the number of books in the group.
				// Example: for groups, 1,2,3,4 and 1,2
				// the discount would be 20% off all the books
				// in group 1 and 5% off all the books in group
				// 2 resulting in a discount of 15%. Not
				// 12.5% off the total (20 + 5 / 2)
				tmpDiscount += countDiscount[len(group)] *
					float64(len(group)) / float64(basketCount)
			}

			// if we have a better discount,
			// use it
			if tmpDiscount < discount {
				discount = tmpDiscount
			}
		}
	}

	baseBookCost := float64(basketCount * bookCost)
	discountedBookCost := int(discount * baseBookCost)
	return discountedBookCost
}

// for baskets len 5 or less we can
// do a quick check for uniqueness
func allUnique(basket []int) bool {
	for i := 1; i < len(basket); i++ {
		if basket[i] == basket[i-1] {
			return false
		}
	}
	return true
}

// keeps track of books seen
type grouping map[int]bool

// This strategy tries to fill up one group
// before starting a new group.
// Example: 1,1,2,2,3,4,5 would result in
// 1,2,3,4,5 and 1,2
func deepGrouping(basket []int) []grouping {
	groups := []grouping{}

	for _, v := range basket {
		added := false
		// try to push to any group
		// starting with the first
		for _, g := range groups {
			// if it doesn't exist in the
			// group add it and move on to
			// the next number
			if !g[v] {
				g[v] = true
				added = true
				break
			}
		}

		// if all groups had this num
		// create a new group and add
		// it
		if !added {
			groups = append(groups, grouping{
				v: true,
			})
		}

	}
	return groups
}

// This strategy tries to group the books
// evenly.
// Example: 1,1,2,2,3,4,5 would result in
// 1, 2, 3, 5 and 1, 2, 4
func broadGrouping(basket []int) []grouping {
	groups := []grouping{}
	// Guestimate number of groups based
	// on the square root of the basket length
	numGroups := int(math.Floor(math.Sqrt((float64(len(basket))))))

	for i, v := range basket {
		// this is the target group
		// we wnt to push to
		groupIndex := i % numGroups

		// make a grouping it exists
		if groupIndex > len(groups)-1 {
			groups = append(groups, grouping{})
		}

		// if value already exists in a group
		// create a new group and increase
		// This will likely occur if we have
		// length of basket < 8
		if groups[groupIndex][v] {
			groups = append(groups, grouping{})
			numGroups++
			groupIndex++
		}

		// mark as exists
		groups[groupIndex][v] = true
	}
	return groups
}
