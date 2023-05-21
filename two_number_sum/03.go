package main

import "sort"

func TwoNumberSum(array []int, target int) []int {
	// space O(1)
	// time O(nlog(n))

	left := 0
    right := len(array) - 1

    sort.Ints(array)

    for i := 0; i < len(array); i++ {
        y := array[left] + array[right]
        if y == target {
            return []int{array[left], array[right]}
        } else if  y < target {
            left++;
        } else {
            right--;
        }
    }

    return []int{}
}
