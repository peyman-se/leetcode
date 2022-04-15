package main

/**
How would you design a stack which, in addition to push and pop has a function min
which returns the minimum element? Push, pop and min should all operate in O(1) time.
*/

type Stack struct {
	Top *Node
	Minimum *Node
}

type Node struct {
	Data int
	Next *Node
}

func (s *Stack) Push(node *Node) {
	if s.Top == nil {
		s.Top = node
		return
	}

	node.Next = s.Top
	s.Top = node

	if s.Minimum == nil || s.Minimum.Data > node.Data {
		s.Minimum = node
	}
}

func (s *Stack) Min() *Node {
	return s.Minimum
}

func (s *Stack) Pop() *Node {
	if s.Top == nil {
		return nil
	}

	node := s.Top
	s.Top = node.Next

	return node
}