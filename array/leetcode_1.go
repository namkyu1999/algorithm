package main

//func main() {
//	nums := []int{2, 7, 11, 15}
//	target := 9
//	answer := twoSum(nums, target)
//	fmt.Println(answer)
//}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		if index, isPresent := hash[num]; isPresent {
			return []int{index, i}
		}
		hash[target-num] = i
	}
	return nil
}
