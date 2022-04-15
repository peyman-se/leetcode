package main

import "testing"

type Case struct {
	Name string
	Input *Tree
	Expected []*List
}


func TestBreakTreeToLists(t *testing.T) {
	cases := []Case{
		{
			Name: "it works with a tree of depth 1",
			Input: &Tree{
				Val: 1,
			},
			Expected: []*List{
				&List{
					Val: 1,
				},
			},
		},
		{
			Name: "it works with a tree of depth 2",
			Input: &Tree{
				Val: 1,
				Left: &Tree{
					Val: 2,
				},
				Right: &Tree{
					Val: 3,
				},
			},
			Expected: []*List{
				&List{
					Val: 1,
				},
				&List{
					Val:2,
					Next: &List{
						Val:3,
					},
				},
			},
		},
		{
			Name: "it works with a tree of depth 3",
			Input: &Tree{
				Val: 1,
				Left: &Tree{
					Val: 2,
					Left: &Tree{
						Val: 3,
					},
					Right: &Tree{
						Val: 4,
					},
				},
				Right: &Tree{
					Val: 5,
					Left: &Tree{
						Val: 6,
					},
					Right: &Tree{
						Val: 7,
					},
				},
			},
			Expected: []*List{
				&List{
					Val: 1,
				},
				&List{
					Val:2,
					Next: &List{
						Val:5,
					},
				},
				&List{
					Val:3,
					Next: &List{
						Val:4,
						Next: &List{
							Val:6,
							Next: &List{
								Val:7,
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := BreakTreeToLists(c.Input)

			if len(result) != len(c.Expected) {
				t.Errorf("Expecting a lengh of %v, got %v", len(c.Expected), len(result))
			}
		})
	}
}

func AreListsIdentical(l1, l2 *List) bool {
	for l1 != nil && l2 != nil {
		if (l1 == nil || l2 == nil) && ! (l1 == nil && l2 == nil) {
			return false
		}

		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}