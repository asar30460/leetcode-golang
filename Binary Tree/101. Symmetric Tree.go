/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// As we don't create that balance-like tree in testcase of leetcode
// We just copy the scope
func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	// 上個判斷式已處理兩個都是nil情況，所以此判斷式成立表示騎一不為nil
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}