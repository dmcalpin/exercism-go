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
	basket = orderByMode(basket)

	l := len(basket)
	discount := countDiscount[1]

	if allUnique(basket) {
		discount = countDiscount[l]
	} else if l >= 5 {
		for i := 3; i < 6; i++ {
			groupedByI := groupUnique(basket, i)

			tmpDiscount := 0.0
			for _, group := range groupedByI {
				lenGroup := len(group)
				tmpDiscount += countDiscount[lenGroup] * float64(lenGroup) / float64(l)
			}

			if tmpDiscount < discount {
				discount = tmpDiscount
			}
		}
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

func countUnique(groupedBooks [][]int) int {
	return len(groupedBooks[0])
}

func orderByMode(basket []int) []int {
	counts := [][]int{
		{}, {}, {}, {}, {},
	}
	for _, elem := range basket {
		counts[elem-1] = append(counts[elem-1], elem)
	}
	sort.SliceStable(counts, func(i, j int) bool {
		return len(counts[j]) < len(counts[i])
	})
	orderedBasket := []int{}
	for i := 0; i < 5; i++ {
		for _, elem := range counts[i] {
			orderedBasket = append(orderedBasket, elem)
		}
	}
	return orderedBasket
}

func groupUnique(basket []int, targetCount int) [][]int {
	lastNum := -1
	groupIndex := 0

	groups := [][]int{
		{},
	}

	baseGroupIndex := 0
	for i := 0; i < len(basket); i++ {
		groupIndex = baseGroupIndex

		if basket[i] == lastNum {
			groupIndex++
			if groupIndex == len(groups) {
				groups = append(groups, []int{})
			}
		}

		groups[groupIndex] = append(groups[groupIndex], basket[i])
		lastNum = basket[i]

		if len(groups[groupIndex])%targetCount == 0 && i != len(basket)-1 {
			baseGroupIndex++
			if len(groups) == baseGroupIndex {
				groups = append(groups, []int{})
			}
			lastNum = -1
		}

	}
	return groups
}
