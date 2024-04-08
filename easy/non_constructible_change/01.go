package main

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	sum := 0
	for _, num := range coins {
		if num > sum + 1 {
			return sum + 1;
		}

		sum += num;
	}

	return sum + 1
}
