package main

//func main() {
//	input := []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}
//	result := dailyTemperatures(input)
//	fmt.Println(result)
//}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack, stackHead := make([]int, 0), -1
	for i := 0; i < len(temperatures)-1; i++ {
		if stackHead != -1 {
			for stackHead > -1 {
				if temperatures[stack[stackHead]] < temperatures[i+1] {
					result[stack[stackHead]] = i - stack[stackHead] + 1
					stack = stack[:len(stack)-1]
					stackHead--
				} else {
					break
				}
			}
		}
		if temperatures[i] < temperatures[i+1] {
			result[i] = 1
		} else {
			stack = append(stack, i)
			stackHead++
		}
	}
	return result
}
