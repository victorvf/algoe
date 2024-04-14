package main

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	possibles := make([][]int, 0)

	for i := 0; i < len(array); i++ {
		left := i + 1
		right := len(array) - 1

		for left < right {
			sum := array[i] + array[left] + array[right]

			if sum == target {
				possible := []int{array[i], array[left], array[right]}
				possibles = append(possibles, possible)

				left++
				right = len(array) - 1
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return possibles
}
