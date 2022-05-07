package main

func preorderTraversal(root *TreeNode) []int {
	data := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		data = append(data, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return data
}
