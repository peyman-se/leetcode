package main

import "testing"

type TestCase struct{
	name string
	input [][]int
	expected [][]int
}

func TestRotateMatrix(t *testing.T) {
	cases := []TestCase{
		{
			name: "it can rotate a N*N matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: [][]int{
				{1, 5, 9, 13},
				{2, 6, 10, 14},
				{3, 7, 11, 15},
				{4, 8, 12, 16},
			},
		},
		{
			name: "it can rotate a N*N matrix",
			input: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func (t *testing.T){
			result := rotateMatrix(c.input)

			if !isSame(result, c.expected) {
				t.Errorf("Expected %v but got %v", c.expected, result)
			}
		})
	}
}

func isSame(a, b [][]int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
        return false
    }

	if len(a[0]) != len(b[0]) {
		return false
	}

	m := len(a[0])
    for i := range a {
		for j:=0; j < m; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
    }

    return true
}