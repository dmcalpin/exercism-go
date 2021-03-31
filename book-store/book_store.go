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

	discount := countDiscount[1]
	if allUnique(basket) {
		discount = countDiscount[len(basket)]
	}
	if l >= 5 {
		lenFloat := float64(l)

		// 5-based discount
		num5discounts := float64(l / 5)
		remainder := float64(l % 5)

		disc5 := countDiscount[5] * (5.0 * num5discounts / lenFloat)
		remainderDisc := countDiscount[int(remainder)] * (remainder / lenFloat)
		discount = disc5 + remainderDisc

		// 4-based discount
		num4discounts := float64(l / 4)
		remainder = float64(l % 4)

		disc4 := countDiscount[4] * (4.0 * num4discounts / lenFloat)
		remainderDisc = countDiscount[int(remainder)] * (remainder / lenFloat)
		discount4 := disc4 + remainderDisc

		if discount4 < discount {
			discount = discount4
		}
	}

	if l%6 == 0 {
		disc4 := countDiscount[4] * (4.0 / 6.0)
		disc2 := countDiscount[2] * (2.0 / 6.0)
		discount = disc4 + disc2
	} else if l%4 == 0 {
		discount = countDiscount[4]
	}

	return int(math.Round((discount*float64(l)*float64(bookCost))*100) / 100)
}

func allUnique(basket []int) bool {
	for i := 1; i < len(basket); i++ {
		if basket[i] == basket[i-1] {
			return false
		}
	}
	return true
}
