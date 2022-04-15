package main

/**
Write code to partition a linked list around a value x, 
such that all nodes less than x come before all nodes greater than or equal to x.
If x is contained within the list, 
the values of x only need to be after the elements less than x (see below).
The partition element x can appear anywhere in the "right partition";
it does not need to appear between the left and right partitions.
EXAMPLE
Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition=5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
*/
type List struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func (l *List) Push(node *Node) {
	if l.Head == nil {
		l.Head = node
		return
	}

	current := l.Head
	for ;current.Next != nil; {
		if current.Next == nil {
			current.Next = node
			return
		}

		current = current.Next
	}
}

func Partition1(list List, x int) List {
	current := list.Head
	var smaller, bigger List
	for ;current != nil; {
		if current.Data >= x {
			bigger.Push(current)
		} else {
			smaller.Push(current)
		}

		current = current.Next
	}

	return merge(smaller, bigger)
}

func Partition2(list List, x int) List {
	head := list.Head
	current := list.Head

	for ;current != nil; {
		if current.Next != nil && current.Next.Data < x {
			//append to head
			d := current.Next
			d.Next = head
			head = d

			//remove from current position
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return list
}

func merge(l1, l2 List) List {
	var result List
	current := l1.Head

	var last *Node
	for ;current != nil; {
		result.Push(current)
		if current.Next == nil {
			last = current
		}
		current = current.Next
	}

	head := l2.Head

	last.Next = head

	return result
}