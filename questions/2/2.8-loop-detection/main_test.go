package main

import "testing"

type TestCase struct {
	Input  func() List
	Expected *Node
}
func TestLoopStart(t *testing.T) {
	b := &Node{
		Data: 2,
	}

	c := &Node{
		Data: 3,
	}

	cases := []TestCase{
		{
			Input: func() List {
				list := List{
					Head: &Node{
						Data: 1,
					},
				}

				list.Head.Next = b

				b.Next = c

				d := &Node{
					Data: 4,
				}
				c.Next = d

				e := &Node{
					Data: 5,
				}
				d.Next = e

				e.Next = c

				return list
			},
			Expected: c,
		},
	}

	for _, c := range cases {
		output := LoopStart(c.Input())
		if output != c.Expected {
			t.Errorf("got %v, want %v", output, c.Expected)
		}
	}
}