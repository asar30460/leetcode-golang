package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if (root == nil) {
		return 0
	}

	rightNode, leftNode, maxNode := 0, 0, 0
	leftNode = maxDepth(root.Left)
	rightNode = maxDepth(root.Right)
	if leftNode > rightNode {
		maxNode = leftNode
	} else {
		maxNode = rightNode
	}

	return maxNode + 1
}

func (t *TreeNode) Insert(val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}

	var visitor *TreeNode = t // 尋訪樹節點
	var parent * TreeNode // 當visitor尋訪到nil時，紀錄此時visitor的父節點

	for visitor != nil {
		parent = visitor
		if val < visitor.Val {
			visitor = visitor.Left
		} else {
			visitor = visitor.Right
		}
	}

	newNode := &TreeNode{Val: val}
	if val <= parent.Val {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	
	return t
}

func main() {
	testCase_1 := [...]int{3, 9, 20, 15, 7}	// This case is not a balanced BST, you won't get the same ans as leetcode.
	testCase_2 := [...]int{1, 2}

	var testTree_1 *TreeNode
	var testTree_2 *TreeNode

	for _, v := range(testCase_1) {
		testTree_1 = testTree_1.Insert(v)
	}

	for _, v := range(testCase_2) {
		testTree_2 = testTree_2.Insert(v)
	}

	fmt.Printf("Case1: %d\n", maxDepth(testTree_1))
	fmt.Printf("Case2: %d\n", maxDepth(testTree_2))
}