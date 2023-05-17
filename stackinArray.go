package main

import "fmt"

type Stack struct {
	data [10]int
	top  int
}

func (s *Stack) Push(item int) {
	if s.top == len(s.data)-1 {
		fmt.Println("Stack overflow")
		return
	}
	s.top++
	s.data[s.top] = item
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	item := s.data[s.top]
	s.top--
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println(item)
	}
}
