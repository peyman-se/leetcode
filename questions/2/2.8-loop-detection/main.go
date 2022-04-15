package main

/**
Given a circular linked list, implement an algorithm which returns 
the node at the beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer points
to an earlier node, so as to make a loop in the linked list.
EXAMPLE
Input: A -> B -> C -> D -> E -> C [the same C as earlier]
Output: C
*/

type List struct {
	Head *Node
	len int
}

type Node struct {
	Data int
	Next *Node
}

func (l *List) Pop() *Node {
	head := l.Head
	l.Head = head.Next
	return head
}

func LoopStart(list List) *Node {
	current := list.Head
	var temp []*Node
	for ;current != nil;{
		for _, t := range temp {
			if t == current {
				return current
			}
		}
		
		temp = append(temp, current)
		current = current.Next
	}

	return nil
}
