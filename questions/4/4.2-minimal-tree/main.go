package main

import "fmt"

// import "math"

/**
Give a sorted (increasing order) array with unique integer elements,
write an algorithm to create a binary search tree with minimal height.
*/

type Tree struct {
	Val int
	LeftNode *Tree
	RightNode *Tree
}

func main() {
	// a := []int{1, 2, 4, 5, 6, 8, 9, 10, 12, 13, 15, 17, 18, 19}
	a := []int{3, 5, 7}
	t := BuildTree(a)
	fmt.Printf("%v", t)
}

func BuildTree(a []int) *Tree {
	// optimumLength := int(math.Log2(float64(len(a))))

	if len(a) == 0 {
		return nil
	}

	if len(a) == 1 {
		return &Tree{
			Val: a[0],
		}
	}

	middleIndex := func(lengthOfArray int) int {
		return lengthOfArray / 2
	}

	i := middleIndex(len(a))
	fmt.Println("middle index", i)
	top := &Tree{
		Val: a[i],
	}

	top.LeftNode = BuildTree(a[:i])
	top.RightNode = BuildTree(a[i+1:])

	return top
}