## Transpose Matrix

You're given a 2D array of integers matrix. Write a function that returns the transpose of the matrix.

The transpose of a matrix is a flipped version of the original matrix across its main diagonal(which runs form top-left to bottom-right); it switches the row and column indices of the original matrix.

You can assume that input matrix always has at least 1 value; however its width and height are not necessarily the same.

### Sample Input
matrix = [
    [1, 2]
]

### Sample Output
[
    [1],
    [2]
]

### Sample Input 2
matrix = [
    [1, 2],
    [3, 4],
    [5, 6]
]

### Sample Output 2
[
    [1, 3, 5],
    [2, 4, 6]
]

### Tests
{"matrix": [[1]]}
{"matrix": [[1], [-1]]}
{"matrix": [[1, 2, 3]]}
{"matrix": [[1], [2], [3]]}
{"matrix": [[1, 2, 3], [4, 5, 6]]}
{"matrix": [[0, 0, 0], [1, 1, 1]]}
{"matrix": [[0, 1], [0, 1], [0, 1]]}
{"matrix": [[0, 0, 0], [0, 0, 0]]}
{"matrix": [[1, 4], [2, 5], [3, 6]]}
{"matrix": [[-7, -7], [100, 12], [-33, 17]]}
{"matrix": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]}
{"matrix": [[1, 4, 7], [2, 5, 8], [3, 6, 9]]}
{"matrix": [[5, 6, 3, -3, 12], [-3, 6, 5, 2, -1], [0, 0, 3, 12, 3]]}
{"matrix": [[0, -1, -2, -3], [4, 5, 6, 7], [2, 3, -2, -3], [42, 100, 30, -42]]}
{"matrix": [[1234, 6935, 4205], [-23459, 314159, 0], [100, 3, 987654]]}
