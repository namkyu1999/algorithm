// This file is initialized with a code version of this
// question's sample test case. Feel free to add, edit,
// or remove test cases in this file as you see fit!

package main

import (
	"fmt"
)

var hash = map[int]int{}

func main() {
	input := [][]int{{1}, {4}, {7}}
	actual := TransposeMatrix(input)
	fmt.Println(actual)
}

func TransposeMatrix(matrix [][]int) [][]int {
	result := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		result[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for k := 0; k < len(matrix[0]); k++ {
			result[k][i] = matrix[i][k]
		}
	}
	return result
}
