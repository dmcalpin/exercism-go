package bookstore

import (
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

	baseBookCost := bookCost * basketCount
	discount := 0.0
	altDiscount := 0.0

	grouping := groupBooks(basket)
	discounts := getDiscounts(grouping)

	for _, d := range discounts {
		perBookDiscount := countDiscount[d]
		discount += float64(d) * perBookDiscount / float64(basketCount)

		if discounts[0] >= 4 && basketCount%4 == 0 {
			perBookDiscount = countDiscount[4]
			altDiscount += float64(d) * perBookDiscount / float64(basketCount)
		}

	}
	if altDiscount != 0 && altDiscount < discount {
		discount = altDiscount
	}

	return int(discount * float64(baseBookCost))
}

func groupBooks(basket []int) map[int]int {
	grouping := make(map[int]int)
	for _, v := range basket {
		grouping[v]++
	}
	return grouping
}

func getDiscounts(grouping map[int]int) []int {
	groupCount := largestCount(grouping)
	discounts := make([]int, groupCount)
	for i := 0; i < groupCount; i++ {
		for k, _ := range grouping {
			if grouping[k] > 0 {
				discounts[i]++
				grouping[k]--
			}
		}
	}
	return discounts
}

func numUnique(basket []int) int {
	return len(groupBooks(basket))
}

func largestCount(grouping map[int]int) int {
	max := 0
	for _, v := range grouping {
		if v > max {
			max = v
		}
	}
	return max
}
