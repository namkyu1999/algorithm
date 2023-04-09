package main

//func main() {
//	nums := []int{-1, 0, 3, 5, 9, 12}
//	target := 9
//	result := search(nums, target)
//	fmt.Println(result)
//}

func search(nums []int, target int) int {
	prefix := 0
	for len(nums) > 0 {
		n := len(nums)
		if nums[n/2] == target {
			return prefix + n/2
		} else if nums[n/2] > target {
			nums = nums[:n/2]
		} else {
			prefix += n/2 + 1
			nums = nums[n/2+1:]
		}
	}
	return -1
}
