package main

//	func main() {
//		height := []int{4, 2, 0, 3, 2, 5}
//		answer := trap(height)
//		fmt.Println(answer)
//	}
func trap(height []int) int {
	volume, left, right := 0, 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax > rightMax {
			right--
			if rightMax > height[right] {
				volume += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
		} else {
			left++
			if leftMax > height[left] {
				volume += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
		}
	}
	return volume
}

func trap2(height []int) int {
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
