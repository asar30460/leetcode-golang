package main

import (
	"errors"
	"fmt"
)

type ChainNode struct {
	data int
	next *ChainNode
}

func NewChainNode(data int) *ChainNode {
	return &ChainNode{
		data: data,
		next: nil,
	}
}

type Chain struct {
	head *ChainNode
}

func NewChain() *Chain {
	return &Chain{
		head: nil,
	}
}

func (c *Chain) Push(data int) {
	newNode := NewChainNode(data)

	if c.head == nil {
		c.head = newNode
		return
	}

	current := c.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (c *Chain) PrintChain() ([]int, error) {
	if c.head == nil {
		return nil, errors.New("chain is empty")
	}
	var list []int
	node := c.head
	for node != nil {
		list = append(list, node.data)
		node = node.next
	}
	return list, nil
}

func hasCycle(c *Chain) bool {
	fast, slow := c.head, c.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}

// pos表示最後一個節點指向回鏈中的指定節點以形成環
// 如果pos=0表示指回第一個節點，以此類推
// 但若是pos=-1，則表示沒有指回去
// 在Leetcode上面，pos不會丟入hasCycle函數
// 所以local端的測資我們需要自己將鏈兜成環
func createTestCase(arr []int, pos int) *Chain {
	testChain := NewChain()
	for _, v := range arr {
		testChain.Push(v)
	}

	// fmt.Println(testChain.PrintChain())

	if pos != -1 {
		var posNode, tailNode *ChainNode
		current := testChain.head

		for i := 0; current != nil; i++ {
			if i == pos {
				posNode = current
			}
			if current.next == nil {
				tailNode = current
			}
			current = current.next
		}
		tailNode.next = posNode
	}

	return testChain
}

func main() {
	testChain_1 := createTestCase([]int{3, 2, 0, -4}, 1)
	testChain_2 := createTestCase([]int{1, 2}, 0)
	testChain_3 := createTestCase([]int{1}, -1)

	fmt.Printf("Case 1: %t\n", hasCycle(testChain_1)) // Expected true
	fmt.Printf("Case 2: %t\n", hasCycle(testChain_2)) // Expected true
	fmt.Printf("Case 3: %t\n", hasCycle(testChain_3)) // Expected false
}
