package main

/**
Imagine a (literal) stack of plates. If the stack gets too high, it might topple.
Therefore, in real life, we would likely start a new stack when the 
previous stack exceeds some thresholds.
Implement a data structure SetOfStacks that mimics this. SetOfStacks should be composed of several
stacks and should created a new stack once the previous one exceeds capacity.
SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack.
(that is, pop() should return the same values as it would if there were just a signle stack)

FOLLOW UP
implement a function popAt(int index) which performs a pop operation on a specific sub-stack.
*/

const THRESHOLD = 3

type SetOfStacks struct {
	Top *Stack
}

type Stack struct {
	Top *Node
	Next *Stack
	Size int
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
}

func (s *Stack) Pop() *Node {
	if s.Top == nil {
		return nil
	}

	node := s.Top
	s.Top = node.Next
	return node
}

func (s *SetOfStacks) Push(node *Node) {
	if s.Top == nil {
		s.Top = &Stack{
			Top: node,
		}
		return
	}

	if s.Top.Size == THRESHOLD {
		s.Top = &Stack{
			Top: node,
			Next: s.Top,
		}

		return
	}
	
	s.Top.Push(node)
}

func (s *SetOfStacks) Pop() *Node {
	if s.Top == nil {
		return nil
	}

	node := s.Top.Pop()
	if s.Top.Top == nil {
		s.Top = nil
	}

	return node
}