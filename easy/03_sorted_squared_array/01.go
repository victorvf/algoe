package main

import "sort"

func SortedSquaredArray(array []int) []int {
	for index, num := range array {
        array[index] = num * num
    }

    sort.Ints(array)

    return array
}
