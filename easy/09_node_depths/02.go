package main


type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

type Leaf struct {
	Node  *BinaryTree
	Depth int
}

func NodeDepths(root *BinaryTree) int {
	sum := 0
	stack := []*Leaf {{Node: root, Depth: 0}}
	var leaf *Leaf

	for len(stack) > 0 {
		leaf, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if leaf.Node == nil {
			continue
		}

		sum += leaf.Depth
		stack = append(stack, &Leaf{Node: leaf.Node.Left, Depth: leaf.Depth + 1})
		stack = append(stack, &Leaf{Node: leaf.Node.Right, Depth: leaf.Depth + 1})
	}

	return sum
}
