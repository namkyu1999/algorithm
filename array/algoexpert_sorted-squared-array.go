package main

import (
	"sort"
)

func SortedSquaredArray(array []int) []int {
	newArr := make([]int, len(array))
	left, right := -1, -1
	i := 0
	for ; i < len(array); i++ {
		if array[i] < 0 {
			left = i
		}
		array[i] = array[i] * array[i]
	}
	if left == -1 {
		if array[0] > array[i-1] {
			k, j := 0, i-1
			for ; k < j; k, j = k+1, j-1 {
				newArr[k], newArr[j] = array[j], array[k]
			}
			if k == j {
				newArr[k] = array[k]
			}
		} else {
			for i, v := range array {
				newArr[i] = v
			}
		}
	} else {
		right = left + 1
		index := 0
		for left >= 0 && right < i {
			if array[left] > array[right] {
				newArr[index] = array[right]
				right++
			} else {
				newArr[index] = array[left]
				left--
			}
			index++
		}
		for left >= 0 {
			newArr[index] = array[left]
			left--
			index++
		}
		for right < i {
			newArr[index] = array[right]
			right++
			index++
		}
	}
	return newArr
}

func SortedSquaredArray2(array []int) []int {
	newArr := make([]int, len(array))
	for i, v := range array {
		newArr[i] = v * v
	}
	sort.Ints(newArr)
	return newArr
}

//func main() {
//	arr := []int{-4, -2, -1}
//	result := SortedSquaredArray(arr)
//	fmt.Println(result)
//}
