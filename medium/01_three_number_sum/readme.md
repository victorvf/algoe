## Three Number Sum

Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum. The function should find all triplets in the array that sum up to the target sum and return a two-dimensional array of all these triplets. The numbers in each triplet should be ordered in ascending order, and the triplets themselves should be ordered in ascending order with respect to the numbers they hold.

If no three numbers sum up to the target sum, the function should return an empty array.

### Sample Input
array = [12, 3, 1, 2, -6, 5, -8, 6]
targetSum = 0

### Sample Output
[[-8, 2, 6], [-8, 3, 5], [-6, 1, 5]]

# Tests
{"array": [1, 2, 3], "targetSum": 6}
{"array": [1, 2, 3], "targetSum": 7}
{"array": [8, 10, -2, 49, 14], "targetSum": 57}
{"array": [12, 3, 1, 2, -6, 5, 0, -8, -1], "targetSum": 0}
{"array": [12, 3, 1, 2, -6, 5, 0, -8, -1, 6], "targetSum": 0}
{"array": [12, 3, 1, 2, -6, 5, 0, -8, -1, 6, -5], "targetSum": 0}
{"array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15], "targetSum": 18}
{"array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15], "targetSum": 32}
{"array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15], "targetSum": 33}
{"array": [1, 2, 3, 4, 5, 6, 7, 8, 9, 15], "targetSum": 5}
