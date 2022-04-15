package main

import "testing"

type Case struct {
	Name string
	Input *Tree
	Expected bool
}

func TestIsBinarySearchTree(t *testing.T) {
	cases := []Case{
		{
			Name: "It returns true for a binary search tree",
			Input: &Tree{
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
			},
			Expected: true,
		},
		{
			Name: "It returns false if not a binary search tree",
			Input: &Tree{
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
						Val: 5,
					},
				},
			},
			Expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := IsBinarySearchTree(c.Input)

			if result != c.Expected {
				t.Errorf("expected %v got %v", c.Expected, result)
			}
		})
	}
}