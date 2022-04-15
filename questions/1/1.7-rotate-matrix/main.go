package main

import "fmt"

/**
Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
write a method to rotate the image by 90 degrees. Can you do this in place?
*/

 func main() {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix := rotateMatrix(m)
	fmt.Println(matrix)
 }

 func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	for i :=0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
 }