/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var result int
	var dfs func(node *TreeNode, i int) int
	dfs = func(node *TreeNode, i int) int {
		if node == nil {
			return i
		}
		i = dfs(node.Left, i) + 1
		if i == k {
			result = node.Val
		}
		return dfs(node.Right, i)
	}
	dfs(root, 0)
	return result
}