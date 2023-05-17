package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
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
