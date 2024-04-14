## Non-Constructible Change

Given an array of positive integers representing the values of coins in your possession, write a function that returns the minimum amount of change (the minimum sum of money) that you **cannot** create. The given coins can have any positive integer value and aren't necessarily unique(i.e., you can have multiple coins of the same value).

For example, if you're given coins = [1, 2, 5], the minimum amount of change that you can't create is 4. If you're given no coins, the minimum amount of change that you can't create is 1.

### Sample Input
coins = [5, 7, 1, 1, 2, 3, 22]

### Sample Output
20

### Tests
{"coins": [5, 7, 1, 1, 2, 3, 22]}
{"coins": [1, 1, 1, 1, 1]}
{"coins": [1, 5, 1, 1, 1, 10, 15, 20, 100]}
{"coins": [6, 4, 5, 1, 1, 8, 9]}
{"coins": []}
{"coins": [87]}
{"coins": [5, 6, 1, 1, 2, 3, 4, 9]}
{"coins": [5, 6, 1, 1, 2, 3, 43]}
{"coins": [1, 1]}
{"coins": [2]}
{"coins": [1]}
{"coins": [109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4]}
{"coins": [1, 2, 3, 4, 5, 6, 7]}
