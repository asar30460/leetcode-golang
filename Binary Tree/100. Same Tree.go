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
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil || q == nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
        return true
    } 

    return false
}