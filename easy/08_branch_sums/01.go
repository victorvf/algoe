package main


type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func inOrder(root *BinaryTree, sum int, results *[]int) {
	if root != nil {
        sum += root.Value

    	inOrder(root.Left, sum, results)
    	inOrder(root.Right, sum, results)

        if root.Left == nil && root.Right == nil {
            *results = append(*results, sum)
        }
    }
}

func BranchSums(root *BinaryTree) []int {
	var results []int
	inOrder(root, 0, &results)
    return results
}
