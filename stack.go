package main

import "fmt"

type Stack struct {
	Len   int
	Stack []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.Stack) == 0
}

func (s *Stack) Push(data int) {
	s.Stack = append(s.Stack, data)
	s.Len++
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("Stack is Empty!")
	}
	res := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	s.Len--
	return res, nil
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	return s.Stack[len(s.Stack)-1]
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for !s.IsEmpty() { // Loop until the stack is empty
		fmt.Println("Top ->", s.Top(), "Total", s.Len)
		v, _ := s.Pop()
		fmt.Println("Pop->", v)
	}
}
