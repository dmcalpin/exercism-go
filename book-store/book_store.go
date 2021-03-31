package bookstore

import (
	"math"
	"sort"
)

const bookCost = 800

var countDiscount = map[int]float64{
	1: 1.0,
	2: 0.95,
	3: 0.9,
	4: 0.8,
	5: 0.75,
}

func Cost(basket []int) int {
	sort.Ints(basket)

	l := len(basket)
	numUnique := countUnique(basket)
	discount := countDiscount[1]

	if allUnique(basket) {
		discount = countDiscount[numUnique]
	}

	if l >= 5 {
		lenFloat := float64(l)

		// 4-based discount greedy (4 + remainder)
		if numUnique >= 4 {
			num4discounts := float64(l / 4)
			remainder := float64(l % 4)

			disc4 := countDiscount[4] * (4.0 * num4discounts / lenFloat)
			remainderDisc := countDiscount[int(remainder)] * (remainder / lenFloat)
			discount = disc4 + remainderDisc

		}

		if numUnique == 5 { // 5-based discount greedy (5 + remainder)
			num5discounts := float64(l / 5)
			remainder := float64(l % 5)

			disc5 := countDiscount[5] * (5.0 * num5discounts / lenFloat)
			remainderDisc := countDiscount[int(remainder)] * (remainder / lenFloat)
			discount5 := disc5 + remainderDisc

			if discount5 < discount {
				discount = discount5
			}
		}

	}

	return int(math.Round((discount*float64(l)*float64(bookCost))*100) / 100)
}

func allUnique(basket []int) bool {
	return countUnique(basket) == len(basket)
}

func countUnique(basket []int) int {
	seen := map[int]bool{}
	for i := 0; i < len(basket); i++ {
		seen[basket[i]] = true
	}
	return len(seen)
}
