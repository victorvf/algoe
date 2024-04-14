package main

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	possibles := make([][]int, 0)

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			for k := j + 1; k < len(array); k++ {
				sum := array[i] + array[j] + array[k]
				if sum == target {
					possible := []int{array[i], array[j], array[k]}
					possibles = append(possibles, possible)
				}
			}
		}
	}

	return possibles
}
