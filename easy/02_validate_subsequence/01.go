package main

func IsValidSubsequence(array []int, sequence []int) bool {
    is_valid := false
    m := make(map[int] []int)

    for i, num := range array { m[num] = []int{i, 0} }

    index := -1
    for _, num := range sequence {
        if val, prs := m[num]; prs {
            if len(m) == 1 {
                return true
            }
            
            if val[1] == 1 {
                return false
            } else {
                val[1] = 1
            }

            if index == -1 {
                index = val[0]
            } else if val[0] < index {
                return false
            } else {
                index = val[0]
            }

            is_valid = true
        } else {
            return false
        }
    }

	return is_valid
}
