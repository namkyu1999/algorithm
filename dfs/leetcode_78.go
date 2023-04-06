package main

//func main() {
//	nums := []int{1, 2, 3}
//	result := subsets(nums)
//	fmt.Println(result)
//}

func subsets(nums []int) [][]int {
	var result [][]int
	var current []int
	subDfs(&result, current, nums, 0)
	return result
}

func subDfs(result *[][]int, current, nums []int, index int) {
	*result = append(*result, append([]int{}, current...))

	for i := index; i < len(nums); i++ {
		subDfs(result, append(current, nums[i]), nums, i+1)
	}
}
