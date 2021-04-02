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
	baseBasketCost := bookCost * basketCount

	discount := 0.0
	altDiscount := 0.0

	// count each type of book
	bookCounts := make(map[int]int, 5)
	largestGroup := 0
	for _, v := range basket {
		bookCounts[v]++
		if bookCounts[v] > bookCounts[largestGroup] {
			largestGroup = v
		}
	}

	// create slices of unique books
	// 1,1,2,2,3 => 1,2,3 and 1,2
	max := bookCounts[largestGroup]
	bookSets := make([]int, max)
	for i, _ := range bookSets {
		for k, _ := range bookCounts {
			if bookCounts[k] > 0 {
				bookSets[i]++
				bookCounts[k]--
			}
		}
	}

	// apply a discount either by num of books per set
	// or try an alternet discount using a grouping of 4
	for _, d := range bookSets {
		discountPart := float64(d) / float64(basketCount)
		discount += countDiscount[d] * discountPart

		// covers 5 & 3 vs 4 & 4 comparison
		if bookSets[0] >= 4 && basketCount%4 == 0 {
			altDiscount += countDiscount[4] * discountPart
		}
	}

	if altDiscount != 0 && altDiscount < discount {
		discount = altDiscount
	}

	return int(discount * float64(baseBasketCost))
}
