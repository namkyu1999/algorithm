package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 13
	answer := confusingNumber(input)
	fmt.Println(answer)
}

func confusingNumber(n int) bool {
	r := []rune(strconv.Itoa(n))
	var arr []rune
	hash := map[rune]rune{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	for _, v := range r {
		if invalidNumberContains(v) {
			return false
		}
		arr = append(arr, hash[v])
	}
	reversed := reverseRune(arr)
	if string(r) == string(reversed) {
		return false
	}
	return true
}

func invalidNumberContains(v rune) bool {
	invalid := []rune{'2', '3', '4', '5', '7'}
	for _, invalidNum := range invalid {
		if v == invalidNum {
			return true
		}
	}
	return false
}

func reverseRune(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
