package main

func TwoNumberSum(array []int, target int) []int {
    // space O(1)
    // time O(n^2) -> n * n-1

	for i := 0; i < len(array); i++ {
        for j := i + 1; j < len(array); j++ {
            if array[i] + array[j] == target {
                return []int{array[i], array[j]}
            }
        }
    }

	return []int{}
}
