package bookstore

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
	basketCount := len(basket)
	if basketCount == 0 {
		return 0
	}

	// count each type of book
	// {1, 1, 2, 2, 3} => {2, 2, 1}
	bookCounts := make([]int, 5)
	largestGroup := 0
	for _, v := range basket {
		if len(bookCounts) == v-1 {
			bookCounts = append(bookCounts, 0)
		}

		bookCounts[v-1]++

		if bookCounts[v-1] > bookCounts[largestGroup] {
			largestGroup = v - 1
		}
	}

	// Count books per unique set
	// 1,1,2,2,3 => 3, 2 meaning
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
	for _, d := range bookSets {
		numBooks := float64(d)
		discount += countDiscount[d] * float64(numBooks)

		// covers 5 & 3 vs 4 & 4 comparison
		if bookSets[0] >= 4 && basketCount%4 == 0 {
			altDiscount += countDiscount[4] * float64(numBooks)
		}
	}

	if altDiscount != 0 && altDiscount < discount {
		discount = altDiscount
	}

	// cost befor discount applied
	baseBasketCost := bookCost * basketCount

	// return discounted basket
	return int(discount * float64(baseBasketCost) / float64(basketCount))
}
