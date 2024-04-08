package main

func TwoNumberSum(array []int, target int) []int {
    // space O(n)
    // time O(n)

	m := make(map[int]bool)

    for _, num := range array {
        y := target - num

        if _, prs := m[y]; prs { return []int{num, y} }

        m[num] = true
    }

    return []int{}
}
