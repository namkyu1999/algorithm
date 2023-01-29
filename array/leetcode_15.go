package main

import (
	"sort"
)

//func main() {
//	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
//	answer := threeSum(nums)
//	fmt.Println(answer)
//}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sumValues := v + nums[left] + nums[right]
			if sumValues > 0 {
				right--
			} else if sumValues < 0 {
				left++
			} else {
				result = append(result, []int{v, nums[left], nums[right]})
				for ; left < right && nums[left] == nums[left+1]; left++ {
				}
				for ; left < right && nums[right] == nums[right-1]; right-- {
				}
				left++
				right--
			}
		}
	}
	return result
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	keys := make(map[[3]int]bool)
	for i, v := range nums {
		hash := make(map[int]int)
		for k := i + 1; k < len(nums); k++ {
			if l, ok := hash[nums[k]]; ok {
				temp := [3]int{v, nums[l], nums[k]}
				if _, ok := keys[temp]; !ok {
					result = append(result, temp[:])
					keys[temp] = true
				}
			}
			hash[-v-nums[k]] = k
		}
	}
	return result
}
