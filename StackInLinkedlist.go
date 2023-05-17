package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(item int) {
	newNode := &Node{data: item}
	if s.top == nil {
		s.top = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	item := s.top.data
	s.top = s.top.next
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Println("Data popped from the stack:")
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println(item)
	}
}
