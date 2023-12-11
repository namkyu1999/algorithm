package main

func TwoNumberSum(array []int, target int) []int {
	hash := map[int]byte{}
	for _, v := range array {
		hash[v] = 1
	}
	for _, v := range array {
		if _, ok := hash[target-v]; ok {
			if v == target-v {
				continue
			}
			return []int{v, target - v}
		}
	}
	return []int{}
}

//func main() {
//	array := []int{1, 2, 3, 4, 5}
//	target := 4
//	result := TwoNumberSum(array, target)
//	fmt.Println(result)
//}
