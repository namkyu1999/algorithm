package main

//func main() {
//	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
//	result := removeDuplicates(nums)
//	fmt.Println(result)
//}

func removeDuplicates(nums []int) int {
	currentIndex := 0
	result := 1
	for _, v := range nums[1:] {
		if nums[currentIndex] != v {
			result++
			currentIndex++
			nums[currentIndex] = v
		}
	}
	return result
}
