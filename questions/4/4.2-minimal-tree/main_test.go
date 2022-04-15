package main

import (
	"testing"
	"fmt"
)

type Case struct {
	name     string
	input    []int
	expected *Tree
}

func TestBuildTree(t *testing.T) {
	cases := []Case{
		{
			name:     "works with empty array",
			input:    []int{},
			expected: nil,
		},
		{
			name:  "works with array of length 1",
			input: []int{5},
			expected: &Tree{
				Val: 5,
			},
		},
		{
			name:     "works with array of length 3",
			input:    []int{3, 5, 7},
			expected: &Tree{
				Val: 5,
				LeftNode: &Tree{
					Val: 3,
				},
				RightNode: &Tree{
					Val: 7,
				},
			},
		},
		{
			name:     "is able to create a binary search tree with odd array length",
			input:    []int{2, 3, 4, 5, 6, 7, 8},
			expected: &Tree{
				Val: 5,
				LeftNode: &Tree{
					Val: 3,
					LeftNode: &Tree{
						Val: 2,
					},
					RightNode: &Tree{
						Val: 4,
					},
				},
				RightNode: &Tree{
					Val: 7,
					LeftNode: &Tree{
						Val: 6,
					},
					RightNode: &Tree{
						Val: 8,
					},
				},
			},
		},
		{
			name:     "is able to create a binary search tree with even array length",
			input:    []int{2, 3, 4, 5, 6, 7},
			expected: &Tree{
				Val: 5,
				LeftNode: &Tree{
					Val: 3,
					LeftNode: &Tree{
						Val: 2,
					},
					RightNode: &Tree{
						Val: 4,
					},
				},
				RightNode: &Tree{
					Val: 7,
					LeftNode: &Tree{
						Val: 6,
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := BuildTree(c.input)
			if !AreTreesIdentical(result, c.expected) {
				t.Errorf("Expecting %v, got %v", c.expected, result)
			}
		})
	}
}

func AreTreesIdentical(a, b *Tree) bool {
	if a == nil && b == nil {
		return true
	}

	return a.Val == b.Val && AreTreesIdentical(a.LeftNode, b.LeftNode) && AreTreesIdentical(a.RightNode, b.RightNode)
}
