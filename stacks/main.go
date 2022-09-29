package main

import (
	"fmt"
	"math"
)

type Node struct {
	value any
	prev  *Node
}

type Stack struct {
	length int
	head   *Node
}

func (s *Stack) Pop() any {
	s.length = int(math.Max(0, float64(s.length-1)))
	if s.length == 0 {
		head := s.head
		s.head = nil
		if head == nil {
			return nil
		}
		return head.value
	}
	head := s.head
	s.head = head.prev
	head.prev = nil
	return head.value
}

func (s *Stack) Push(v any) {
	node := Node{value: v}
	if s.length == 0 {
		s.head = &node
		s.length++
		return
	}
	head := s.head
	node.prev = head
	s.head = &node
	s.length++
}

func (s *Stack) Peek() any {
	return s.head.value
}

func main() {

	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())

}
