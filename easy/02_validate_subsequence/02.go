package main

func IsValidSubsequence(array []int, sequence []int) bool {
    index := 0

    for _, num := range array {
        if index < len(sequence) {
            if num == sequence[index] { index++ }
        }
    }

    return index != 0 && index == len(sequence)
}
