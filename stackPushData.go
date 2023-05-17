package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Println("Data pushed into the stack:")
	for _, item := range stack.data {
		fmt.Println(item)
	}
}
