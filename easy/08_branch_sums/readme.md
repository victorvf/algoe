## Branch Sums

Write a function that takes in a Binary Tree and returns a list of its branch sums ordered from leftmost branch sum to rightmost branch sum.

A branch sum is the sum of all values in a Binary Tree branch. A Binary Tree branch is a path of nodes in a tree that starts at the root node and ends at any leaf node.

Each `BinaryTree` node has an integer `value`, a `left` child node, and a `right` child node. Children nodes can either be `BinaryTree` nodes themselves or `None` / `null`.

### Sample Input
```
tree =      1
          /    \
        2        3
      /   \    /   \
    4      5  6     7
   / \    /
  8   9  10
```

### Sample Output
```
[15, 16, 18, 10, 11]
// 15 == 1 + 2 + 4 + 8
// 16 == 1 + 2 + 4 + 9
// 18 == 1 + 2 + 5 + 10
// 10 == 1 + 3 + 6
// 11 == 1 + 3 + 7
```

### Tests
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "4", "left": "8", "right": "9", "value": 4},
      {"id": "5", "left": "10", "right": null, "value": 5},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "8", "left": null, "right": null, "value": 8},
      {"id": "9", "left": null, "right": null, "value": 9},
      {"id": "10", "left": null, "right": null, "value": 10}
    ],
    "root": "1"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "1", "left": null, "right": null, "value": 1}
    ],
    "root": "1"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": null, "value": 1},
      {"id": "2", "left": null, "right": null, "value": 2}
    ],
    "root": "1"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": null, "right": null, "value": 2},
      {"id": "3", "left": null, "right": null, "value": 3}
    ],
    "root": "1"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": null, "right": null, "value": 3},
      {"id": "4", "left": null, "right": null, "value": 4},
      {"id": "5", "left": null, "right": null, "value": 5}
    ],
    "root": "1"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "4", "left": "8", "right": "9", "value": 4},
      {"id": "5", "left": "10", "right": "1-2", "value": 5},
      {"id": "6", "left": "1-3", "right": "1-4", "value": 6},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "8", "left": null, "right": null, "value": 8},
      {"id": "9", "left": null, "right": null, "value": 9},
      {"id": "10", "left": null, "right": null, "value": 10},
      {"id": "1-2", "left": null, "right": null, "value": 1},
      {"id": "1-3", "left": null, "right": null, "value": 1},
      {"id": "1-4", "left": null, "right": null, "value": 1}
    ],
    "root": "1"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "0", "left": "1", "right": null, "value": 0},
      {"id": "1", "left": "10", "right": null, "value": 1},
      {"id": "10", "left": "100", "right": null, "value": 10},
      {"id": "100", "left": null, "right": null, "value": 100}
    ],
    "root": "0"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "0", "left": null, "right": "1", "value": 0},
      {"id": "1", "left": null, "right": "10", "value": 1},
      {"id": "10", "left": null, "right": "100", "value": 10},
      {"id": "100", "left": null, "right": null, "value": 100}
    ],
    "root": "0"
  }
}
{
  "tree": {
    "nodes": [
      {"id": "0", "left": "9", "right": "1", "value": 0},
      {"id": "9", "left": null, "right": null, "value": 9},
      {"id": "1", "left": "15", "right": "10", "value": 1},
      {"id": "15", "left": null, "right": null, "value": 15},
      {"id": "10", "left": "100", "right": "200", "value": 10},
      {"id": "100", "left": null, "right": null, "value": 100},
      {"id": "200", "left": null, "right": null, "value": 200}
    ],
    "root": "0"
  }
}
