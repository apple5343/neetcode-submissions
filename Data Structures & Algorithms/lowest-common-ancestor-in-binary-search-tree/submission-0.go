/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    var result *TreeNode
	var dfs func(root *TreeNode) (bool, bool)
	dfs = func(root *TreeNode) (bool, bool) {
		var pFlag, qFlag bool
		if root == nil {
			return pFlag, qFlag
		}
		if root == p {
			pFlag = true
		} else if root == q {
			qFlag = true
		}
		leftP, leftQ := dfs(root.Left)
		rightP, rightQ := dfs(root.Right)
		pFlag = pFlag || leftP || rightP
		qFlag = qFlag || leftQ || rightQ
		if pFlag && qFlag && result == nil {
			result = root
		}
		return pFlag, qFlag
	}
	dfs(root)
	return result
}
