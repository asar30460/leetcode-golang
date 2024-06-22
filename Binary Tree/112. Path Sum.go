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
 func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }

    leftVal := targetSum - root.Val    

    return hasPathSum(root.Left, leftVal) || hasPathSum(root.Right, leftVal)
}