package main


func SortedSquaredArray(array []int) []int {
    output := make([]int, len(array))

    start := 0
    end := len(array) - 1

    for i := end; i >= 0; i-- {
        smallest := array[start] * array[start]
        largest := array[end] * array[end]

        if smallest < largest {
            output[i] = largest
            end--
        } else {
            output[i] = smallest
            start++
        }
    }

    return output
}
