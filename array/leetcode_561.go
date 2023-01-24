package main

import (
	"sort"
)

//func main() {
//	nums := []int{1, 4, 3, 2}
//	answer := arrayPairSum(nums)
//	fmt.Println(answer)
//}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	answer := 0
	for i := 0; i < len(nums); i += 2 {
		answer += nums[i]
	}
	return answer
}
