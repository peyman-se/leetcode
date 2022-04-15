package main

import (
	"fmt"
)

/**
Given a binary tree, design an algorithm which creates a linked list of all nodes at each depth
(e.g. if you have a tree with depth D, you'll have D linked lists).

*/

type Tree struct {
	Val int
	Left *Tree
	Right *Tree
}

type List struct {
	Val int
	Next *List
	Left *Tree
	Right *Tree
}

func (l *List) Add (list *List) {
	// if l == nil {
	// 	l = list
	// } else {
		l.Next = list
	// }
}


func BreakTreeToLists(t *Tree) []*List {
	var lists []*List
	list := &List{
		Val: t.Val,
		Left: t.Left,
		Right: t.Right,
	}

	lists = append(lists, list)
	i := 1
	for {
		list = GetLinkedListFromPrevious(list)
		if list == nil {
			fmt.Println("list is nil")
			break
		}

		i++
		fmt.Println(i)
		lists = append(lists, list)
	}
	
	return lists
}

func GetLinkedListFromPrevious(d *List) *List {
	if d == nil {
		return nil
	}

	l := &List{
		Val: d.Val,
		Left: d.Left.Left,
		Right: d.Left.Right,
	}

	l.Next = &List{
		Val: d.Val,
		Left: d.Right.Left,
		Right: d.Right.Right,
	}
	current := d.Next
	for {
		fmt.Println("current is not nil")
		if current.Left != nil {
			fmt.Println("current.Left is not nil")
			l.Add(&List{
				Val: current.Left.Val,
				Left: current.Left.Left,
				Right: current.Left.Right,
			})
			l = l.Next
		}
		
		if current.Right != nil {
			fmt.Println("current.Right is not nil")
			l.Add(&List{
				Val: current.Right.Val,
				Left: current.Right.Left,
				Right: current.Right.Right,
			})

			fmt.Printf("%v \n", l)
		}

		current = current.Next
		if current == nil {
			break
		}
		
		l = l.Next
	}

	return l
}

