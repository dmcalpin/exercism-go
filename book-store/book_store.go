package bookstore

// The cost of single book
// in cents
const bookCost = 800

// discount of the books based on
// number of unique books in
// the group. The index matches
// the book number in the basket
var countDiscount = []float64{
	0.0, // to bump the indices up
	1.0,
	0.95,
	0.9,
	0.8,
	0.75,
}

// Cost calculates the best possible
// discount for a given basket of
// books
func Cost(basket []int) int {
	n := len(basket)
	if n == 0 {
		return 0
	}

	// count each type of book
	// {1, 1, 2, 2, 3} => {2, 2, 1}
	bookCounts := [5]int{}
	largestGroup := 0
	for _, v := range basket {
		bookCounts[v-1]++

		if bookCounts[v-1] > bookCounts[largestGroup] {
			largestGroup = v - 1
		}
	}

	// Count books per unique set
	// {2, 2, 1} => {3, 2} meaning
	// the first set has 3 books,
	// and the second set has 2
	max := bookCounts[largestGroup]
	bookSets := make([]int, max)
	for i := range bookSets {
		for k := range bookCounts {
			if bookCounts[k] > i {
				bookSets[i]++
			}
		}
	}

	// apply a discount either by num of books per set
	// or try an alternet discount using a grouping of 4
	discount := 0.0
	altDiscount := 0.0
	for _, s := range bookSets {
		numBooks := float64(s)
		discount += countDiscount[s] * float64(numBooks)

		// covers 5 & 3 vs 4 & 4 comparison
		if bookSets[0] >= 4 && n%4 == 0 {
			altDiscount += countDiscount[4] * float64(numBooks)
		}
	}

	if altDiscount != 0 && altDiscount < discount {
		discount = altDiscount
	}

	// cost befor discount applied
	baseBasketCost := bookCost * n

	// return discounted basket
	return int(discount * float64(baseBasketCost) / float64(n))
}
