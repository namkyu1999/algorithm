package main

//func main() {
//	input := []int{3, 2, 2, 3}
//	result := removeElement(input, 3)
//	fmt.Println(result)
//}

func removeElement(nums []int, val int) int {
	result := 0
	for _, v := range nums {
		if v != val {
			result++
		}
	}
	return result
}
