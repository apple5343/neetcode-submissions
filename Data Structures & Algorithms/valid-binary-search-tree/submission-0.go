/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt, math.MaxInt)
}

func check(node *TreeNode, minValue, maxValue int) bool {
	if node == nil {
		return true
	}
	return minValue < node.Val && node.Val < maxValue && check(node.Left, minValue, node.Val) && check(node.Right, node.Val, maxValue)
}
