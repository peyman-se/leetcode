package main

import "testing"

type Case struct {
	name string
	input string
	expected string
}

func TestCompress(t *testing.T) {
	cases := []Case {
		{
			input: "aabcccccaaa",
			expected: "a2b1c5a3",
		},
		{
			input: "abca",
			expected: "abca",
		},
		{
			input: "",
			expected: "",
		},
		{
			input: "a",
			expected: "a",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Compress(c.input)

			if result != c.expected {
				t.Errorf("Expected %s but got %s", c.expected, result)
			}
		})
	}
}