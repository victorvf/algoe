package main


type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func inOrder(root *BinaryTree, sum int, heights *[]int, visited map[int]bool){
    if root != nil {
		if _, prs := visited[root.Value]; !prs {
			sum++
			visited[root.Value] = true
		}

        *heights = append(*heights, sum)

		inOrder(root.Left, sum, heights, visited)
		inOrder(root.Right, sum, heights, visited)
    }
}

func NodeDepths(root *BinaryTree) int {
	var heights []int
	visited := map[int]bool {root.Value: true}

	inOrder(root, 0, &heights, visited)

	depth := 0
	for _, val := range heights { depth += val }
	return depth
}
