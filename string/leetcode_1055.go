package main

//func main() {
//	source, target := "aaaaa", "aaaaaaaaaaaaa"
//	result := shortestWay(source, target)
//	fmt.Println(result)
//}

func shortestWay(source string, target string) int {
	sourceIndex, prev, targetIndex := 0, 0, 0
	rune1, rune2 := []rune(source), []rune(target)
	result := 0
	for {
		if sourceIndex == len(rune1) {
			if prev == targetIndex {
				return -1
			}
			sourceIndex = 0
			prev = targetIndex
			result++
		}
		if targetIndex == len(rune2) {
			if sourceIndex != 0 {
				result++
			}
			break
		}
		if rune1[sourceIndex] == rune2[targetIndex] {
			targetIndex++
		}
		sourceIndex++
	}
	return result
}
