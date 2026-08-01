/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	layer := []*TreeNode{root}
	for len(layer) != 0 {
		nextLayer := []*TreeNode{}
		nums := []int{}
		for _, n := range layer {
			nums = append(nums, n.Val)
			if n.Left != nil {
				nextLayer = append(nextLayer, n.Left)
			}
			if n.Right != nil {
				nextLayer = append(nextLayer, n.Right)
			}
		}
		result = append(result, nums)
		layer = nextLayer
	}
	return result
}
