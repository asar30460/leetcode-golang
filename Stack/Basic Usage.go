package Stack

import (
	"errors"
	"fmt"
)

// Stack represents a stack data structure
type Stack struct {
	data []int
	top  int
	size int
}

// NewStack initializes a new stack with a given size
func NewStack(size int) *Stack {
	return &Stack{
		data: make([]int, size),
		top:  -1,
		size: size,
	}
}

// Pop removes the top element from the stack and returns it
func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	value := s.data[s.top]
	s.top--
	return value, nil
}

// Push adds an element to the top of the stack
func (s *Stack) Push(v int) error {
	if s.top == s.size-1 {
		return errors.New("stack is full")
	}
	s.top++
	s.data[s.top] = v
	return nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

// IsFull checks if the stack is full
func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (int, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	return s.data[s.top], nil
}

func main() {
	stack := NewStack(10)

	// Pop from an empty stack
	if _, err := stack.Pop(); err != nil {
		fmt.Println(err)
	}

	// Push elements to the stack
	if err := stack.Push(4); err != nil {
		fmt.Println(err)
	}

	// Peek the top element
	if top, err := stack.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Top element: %d\n", top)
	}

	// Pop the top element
	if top, err := stack.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Popped element: %d\n", top)
	}
}