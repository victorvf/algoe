package main

func TransposeMatrix(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]));

	for _, row := range matrix {
		for i, val := range row {
			result[i] = append(result[i], val);
		}
	}

	return result
}
