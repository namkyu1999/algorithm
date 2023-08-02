package main

//func main() {
//	input := []int{1, 2, 3}
//	result := permute(input)
//	fmt.Println(result)
//}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	pdfs(nums, []int{}, &result)
	return result
}

func pdfs(nums []int, current []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, current)
	} else {
		for i := 0; i < len(nums); i++ {
			newNums := make([]int, len(nums))
			newCurrent := make([]int, len(current))
			copy(newNums, nums)
			copy(newCurrent, current)
			pdfs(append(newNums[:i], newNums[i+1:]...), append(newCurrent, nums[i]), result)
		}
	}
}
