package main

func TransposeMatrix(matrix [][]int) [][]int {
	len_column := len(matrix[0]);
	len_row := len(matrix);

	result := make([][]int, len_column);

	for col :=0; col < len_column; col++ {
		nrow := make([]int, len_row);

		for row := 0; row < len_row; row++ {
			nrow[row] = matrix[row][col];
		}

		result[col] = nrow;
	}

	return result
}
