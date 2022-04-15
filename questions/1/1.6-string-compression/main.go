package main

import (
	"fmt"
	"strconv"
)

/**
Implement a method to perform basic string compression using the counts of repeated characters.
For example, the string aabcccccaaa would become a2b1c5a3.
If the "compressed" string would not become smaller than the original string,
your method should return the original string.
You can assume the string has only uppercase and lowercase letters (a - z).
*/

func main() {
	s := "aabcccccaaa"
	fmt.Println(Compress(s))
}

func Compress(s string) string {
	if len(s) < 3 {
		return s
	}

	count := 0
	result := []Compressed{}
	current := s[0]
	for i:=0; i < len(s); i++ {
		if current != s[i] {
			result = append(result, Compressed{
				Character: current,
				Repeats: count,
			})
			
			count = 1
			current = s[i]

			continue
		}

		current = s[i]
		count +=1
	}

	result = append(result, Compressed{
		Character: current,
		Repeats: count,
	}) 

	r := ""
	if len(result) == len(s) {
		return s
	}

	for _, c := range result {
		r = r + string(c.Character) + strconv.Itoa(c.Repeats)
	}

	return string(r)
}

type Compressed struct {
	Character byte
	Repeats int
}