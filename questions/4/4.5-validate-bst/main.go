package main

import "fmt"

/**
implement a function to check if a binary tree is a binary search tree.
*/

type Tree struct {
	Val int
	Left *Tree
	Right *Tree
}

func main() {
	a := &Tree{
		Val: 5,
		Left: &Tree{
			Val: 3,
			Left: &Tree{
				Val: 2,
			},
			Right: &Tree{
				Val: 4,
			},
		},
		Right: &Tree{
			Val: 7,
			Left: &Tree{
				Val: 6,
			},
			Right: &Tree{
				Val: 8,
			},
		},
	}

	fmt.Printf("Is Binary %v \n", IsBinarySearchTree(a))
}

func IsBinarySearchTree(t *Tree) bool {
	if t == nil {
		return true
	}

	if t.Left == nil && t.Right == nil {
		return true
	}

	//how does this work?
	if t.Left == nil && t.Right != nil {
		return false
	}

	return t.Val >= t.Left.Val && 
		t.Right.Val >= t.Val && 
		IsBinarySearchTree(t.Left) && 
		IsBinarySearchTree(t.Right)
}