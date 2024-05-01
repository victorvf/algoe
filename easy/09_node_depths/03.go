package main


type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func SumDepths(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}

	return depth + SumDepths(root.Left, depth + 1) + SumDepths(root.Right, depth + 1)
}

func NodeDepths(root *BinaryTree) int {
	return SumDepths(root, 0)
}
