package main

//func main() {
//	//arr := []int{10, 13, 12, 14, 15}
//	//result := oddEvenJumps(arr)
//	//fmt.Println(result)
//}

func oddEvenJumps(arr []int) int {
	result := 0
Entire:
	for i, v := range arr {
		oddFlag := true
		currentIndex := i
		for currentIndex < len(arr) {
			if currentIndex == len(arr)-1 {
				result += 1
				break
			}
			index := -1
			for k := i + 1; k < len(arr); k++ {
				if oddFlag {
					if v <= arr[k] {
						if index == -1 {
							index = k
						} else if arr[index] <= arr[k] {
							index = k
						}
					}
				} else {
					if v >= arr[k] {
						if index == -1 {
							index = k
						} else if arr[index] >= arr[k] {
							index = k
						}
					}
				}
			}
			if index != i {
				currentIndex = index
				oddFlag = !oddFlag
			} else {
				continue Entire
			}
		}
	}
	return result
}
