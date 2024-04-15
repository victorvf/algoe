package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func abs(val int) int {
	if val < 0 { return -val } else { return val }
}

func getClosest(node *BST, target int) int {
    currentNode := node
    closest := currentNode.Value

    for currentNode != nil {
        if currentNode.Value == target { return target }

        if abs(currentNode.Value - target) < abs(closest - target) {
    		closest = currentNode.Value
    	}
    
    	if currentNode.Value < target {
    		currentNode = currentNode.Right
    	} else {
            currentNode = currentNode.Left
        }
    }

    return closest
}

func (tree *BST) FindClosestValue(target int) int {
	return getClosest(tree, target)
}
