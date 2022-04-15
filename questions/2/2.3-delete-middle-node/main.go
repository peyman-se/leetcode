package main

/**
Implement an algorithm to delete a node in the middle of a singly linked list, 
given only access to that node.
EXAMPLE
Input: the node c from the linked list a->b->c->d->e
Result: nothing is returned, but the new linked list looks like a->b->d->e
*/

type Node struct {
	Data interface{}
	Next *Node
}

func Delete(node *Node){
	nn := node.Next.Next
	node = node.Next
	node.Next = nn
}