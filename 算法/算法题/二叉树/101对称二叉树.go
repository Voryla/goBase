package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	leftData, rightData := make([]int, 0), make([]int, 0)
	var preorder func(node *TreeNode, left bool)
	preorder = func(node *TreeNode, left bool) {
		if node == nil {
			if left {
				leftData = append(leftData, 0)
			} else {
				rightData = append(rightData, 0)
			}
			return
		}
		if left {
			leftData = append(leftData, node.Val)
			preorder(node.Left, left)
			preorder(node.Right, left)
		} else {
			rightData = append(rightData, node.Val)
			preorder(node.Right, left)
			preorder(node.Left, left)
		}
	}
	preorder(root.Left, true)
	preorder(root.Right, false)
	if len(leftData) != len(rightData) {
		return false
	}
	for i := 0; i < len(leftData); i++ {
		if leftData[i] != rightData[i] {
			return false
		}
	}
	return true
}
func main() {
	tree := TreeNode{1,
		&TreeNode{2, nil, &TreeNode{3, nil, nil}},
		&TreeNode{2, nil, &TreeNode{3, nil, nil}}}

	isSymmetric(&tree)
}
