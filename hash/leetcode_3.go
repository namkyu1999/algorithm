package main

//	func main() {
//		s := "pwwkew"
//		result := lengthOfLongestSubstring(s)
//		fmt.Println(result)
//	}

func lengthOfLongestSubstring(s string) int {
	dict := [128]bool{}
	candidate, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		v := s[i]
		if dict[v] {
			for ; dict[v]; j++ {
				candidate--
				dict[s[j]] = false
			}
		}

		dict[v] = true
		candidate++
		if candidate > max {
			max = candidate
		}
	}
	return max
}

func lengthOfLongestSubstring1(s string) int {
	result, candidate := 0, 0
	isDuplicated := make(map[byte]int)
	for i := 32; i < 127; i++ {
		isDuplicated[byte(i)] = -1
	}
	for i, v := range []byte(s) {
		if isDuplicated[v] > -1 {
			if candidate > result {
				result = candidate
			}
			candidate = i - isDuplicated[v] - 1
			for _, d := range []byte(s)[:isDuplicated[v]] {
				if isDuplicated[v] > isDuplicated[d] {
					isDuplicated[d] = -1
				}
			}
		}
		candidate += 1
		isDuplicated[v] = i
	}
	if result > candidate {
		return result
	}
	return candidate
}
