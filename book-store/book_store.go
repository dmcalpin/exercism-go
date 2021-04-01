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
		var groups []grouping
		for i := 0; i < 2; i++ {
			if i == 0 {
				groups = broadGrouping(basket)
			} else {
				groups = deepGrouping(basket)
			}

			tmpDiscount := 0.0
			for _, group := range groups {
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

type grouping map[int]bool

func deepGrouping(basket []int) []grouping {
	groups := []grouping{}

	for _, v := range basket {
		added := false
		for _, g := range groups {
			if !g[v] {
				g[v] = true
				added = true
				break
			}
		}

		if !added {
			groups = append(groups, grouping{
				v: true,
			})
		}

	}
	return groups
}

func broadGrouping(basket []int) []grouping {
	groups := []grouping{}
	numGroups := int(math.Floor(math.Sqrt((float64(len(basket))))))

	for i := 0; i < numGroups; i++ {
		groups = append(groups, grouping{})
	}

	for i, v := range basket {
		groupIndex := i % numGroups
		groups[groupIndex][v] = true
	}
	return groups
}
