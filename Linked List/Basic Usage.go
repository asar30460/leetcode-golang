package LinkedList

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
	NewNode := NewChainNode(data)

	if c.head == nil {
		c.head = NewNode
		return
	} else {
		var current *ChainNode = c.head
		for current.next != nil {
			current = current.next
		}

		current.next = NewNode
	}
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

func main() {
	var gradeChain = NewChain()

	// Check if the chain is empty
	if gradeChain.head == nil {
		fmt.Println(errors.New("chain is empty"))
	} else {
		// Since the chain is not empty, print the data
		fmt.Printf("%d", gradeChain.head.data)
	}

	// For demonstration, add some data and print the chain
	gradeChain.Push(90)
	gradeChain.Push(80)
	gradeChain.Push(70)

	result, err := gradeChain.PrintChain()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
