package main

//func main() {
//	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//	answer := trap(height)
//	fmt.Println(answer)
//}

func trap(height []int) int {
	volume, startIndex, isDesc := 0, 0, false
	for i := 0; i < len(height)-1; i++ {
		if height[i] > height[i+1] {
			if !isDesc {
				isDesc = true
				startIndex = i
			}
		} else if height[i] < height[i+1] {
			if isDesc {
				if len(height) > i+2 && height[i+1] < height[i+2] {
					continue
				}
				if height[startIndex] >= height[i+1] {
					continue
				}
				h := 0
				if height[startIndex] >= height[i+1] {
					h = height[i+1]
				} else {
					h = height[startIndex]
				}
				for t := startIndex + 1; t < i+1; t++ {
					volume += h - height[t]
				}
				isDesc = false
			}
		}
	}
	return volume
}
