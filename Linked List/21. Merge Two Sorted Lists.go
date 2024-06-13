package main

import (
	"errors"
	"fmt"
	"sort"
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

func mergeTwoLists(list1 *Chain, list2 *Chain) *Chain {
	var res []int

	// 直接append到slice後再對slice排序
	for list1.head != nil {
		res = append(res, list1.head.data)
		list1.head = list1.head.next
	}
	for list2.head != nil {
		res = append(res, list2.head.data)
		list2.head = list2.head.next
	}

	sort.Ints(res)

	resChain := NewChain()
	for i := 0; i < len(res); i++ {
		resChain.Push(res[i])
	}

	return resChain
}

func createTestCase(arr []int) *Chain {
	testChain := NewChain()
	for _, v := range arr {
		testChain.Push(v)
	}
	return testChain
}

func main() {
	testChain_1_A := createTestCase([]int{1, 2, 4})
	testChain_1_B := createTestCase([]int{1, 3, 4})
	testChain_2_A := createTestCase([]int{})
	testChain_2_B := createTestCase([]int{})
	testChain_3_A := createTestCase([]int{})
	testChain_3_B := createTestCase([]int{0})

	res, _ := mergeTwoLists(testChain_1_A, testChain_1_B).PrintChain()
	fmt.Printf("Case 1: %v\n", res)

	res, _ = mergeTwoLists(testChain_2_A, testChain_2_B).PrintChain()
	fmt.Printf("Case 2: %d\n", res)

	res, _ = mergeTwoLists(testChain_3_A, testChain_3_B).PrintChain()
	fmt.Printf("Case 3: %d\n", res)
}
