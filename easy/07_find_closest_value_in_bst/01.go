package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func abs(val int) int {
	if val < 0 { return -val } else { return val }
}

func getClosest(node *BST, target int, closest int) int {
	if node == nil {
		return closest
	} else if node.Value == target {
		return target
	}

	if closest == -1 { closest = node.Value }

	if abs(node.Value - target) < abs(closest - target) {
		closest = node.Value
	}

	if node.Value < target {
		return getClosest(node.Right, target, closest)
	}

	return getClosest(node.Left, target, closest)
}

func (tree *BST) FindClosestValue(target int) int {
	return getClosest(tree, target, -1)
}
