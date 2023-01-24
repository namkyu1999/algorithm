package main

//func main() {
//	nums := []int{1, 2, 3, 4}
//	answer := productExceptSelf(nums)
//	fmt.Println(answer)
//}

func productExceptSelf(nums []int) []int {
	n, zeroIndex, allProduct := len(nums), -1, 1
	answer := make([]int, n)
	for i, v := range nums {
		if v == 0 {
			if zeroIndex >= 0 {
				return answer
			}
			zeroIndex = i
		} else {
			allProduct *= v
		}
	}
	if zeroIndex >= 0 {
		answer[zeroIndex] = allProduct
	} else {
		for i, v := range nums {
			answer[i] = allProduct / v
		}
	}
	return answer
}
