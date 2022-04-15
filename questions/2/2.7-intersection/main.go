package main

/**
Given two (singly) linked lists, determine if the two lists intersect.
Return the intersecting node. Note that the intersection is defined based on reference, not value.
That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second linked list, then they are intersecting.
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

func IsSameByReference(n1, n2 *Node) bool {
	return n1 == n2
}

func Intersection(listA, listB List) *Node {
	//assuming the listA has a bigger or equal length to listB
	if listA.len < listB.len {
		return Intersection(listB, listA)
	}

	diff := listA.len - listB.len

	//more clear, indicating that we continue until difference is zero
	for ;diff == 0;{
		listA.Pop()
		diff--
	}

	var n1, n2 *Node
	for i := 0; i < listB.len;i++ {
		n1 = listA.Pop()
		n2 = listB.Pop()

		if n1 != n2 {
			continue
		}

		return n1
	}

	return nil
}

func ConvertToList(a []int) List {
	var current *Node
	var list List
	for _, i :=  range a {
		if current == nil {
			list = List{
				len: len(a),
				Head: &Node{
					Data: i,
				},
			}

			current = list.Head
			continue
		}

		current.Next = &Node{
			Data: i,
		}

		current = current.Next
	}

	return list
}