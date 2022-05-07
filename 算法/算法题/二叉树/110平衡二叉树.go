package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dept func(node *TreeNode) int
	dept = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max1(dept(node.Left), dept(node.Right)) + 1
	}
	return abs(dept(root.Left)-dept(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
