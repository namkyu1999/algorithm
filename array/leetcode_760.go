package main

import "fmt"

func main() {
	nums1 := []int{12, 28, 46, 32, 50}
	nums2 := []int{50, 12, 32, 46, 28}
	result := anagramMappings(nums1, nums2)
	fmt.Println(result)
}

func anagramMappings(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	for i, v := range nums2 {
		hash[v] = i
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = hash[nums1[i]]
	}
	return nums1
}
