package main

import "fmt"

func main() {
	arr := [][]int{
		[]int{1, 5},
		[]int{3, 4},
	}
	result := maxDistance(arr)
	fmt.Println(result)
}

func maxDistance(arrays [][]int) int {
	top, topCandidate, topIndex, bottom, bottomCandidate, bottomIndex := arrays[0][0], -100000, 0, arrays[0][len(arrays[0])-1], 100000, 0
	for i, v := range arrays {
		if top < v[len(v)-1] {
			topCandidate = top
			top = v[len(v)-1]
			topIndex = i
		} else if topCandidate < v[len(v)-1] {
			topCandidate = v[len(v)-1]
		}
		if bottom > v[0] {
			bottomCandidate = bottom
			bottom = v[0]
			bottomIndex = i
		} else if bottomCandidate > v[0] {
			bottomCandidate = v[0]
		}
	}
	if topIndex == bottomIndex {
		one := topCandidate - bottom
		two := top - bottomCandidate
		if one > two {
			return one
		} else {
			return two
		}
	}
	return top - bottom
}
