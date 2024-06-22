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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 一直走訪下去（遞迴），交換左右子樹
	tempTree := root.Left
	root.Left = root.Right
	root.Right = tempTree

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func (t *TreeNode) Insert(val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}

	var visitor *TreeNode = t // 尋訪樹節點
	var parent *TreeNode      // 當visitor尋訪到nil時，紀錄此時visitor的父節點

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

func Inorder(t *TreeNode) {
	if (t != nil) {
		Inorder(t.Left)
		fmt.Printf("%d, ", t.Val)
		Inorder(t.Right)
	}
}



func main() {
	testCase_1 := [...]int{2, 1, 3} 
	testCase_2 := [...]int{2, 3, 1}
	testCase_3 := [...]int{4,2,7,1,3,6,9} 

	var testTree_1 *TreeNode
	var testTree_2 *TreeNode
	var testTree_3 *TreeNode

	for _, v := range testCase_1 {
		testTree_1 = testTree_1.Insert(v)
	}

	for _, v := range testCase_2 {
		testTree_2 = testTree_2.Insert(v)
	}

	for _, v := range testCase_3 {
		testTree_3 = testTree_3.Insert(v)
	}

	testTree_1 = invertTree(testTree_1)
	testTree_2 = invertTree(testTree_2)
	testTree_3 = invertTree(testTree_3)

	Inorder(testTree_1)
	println()
	Inorder(testTree_2)
	println()
	Inorder(testTree_3)
}
