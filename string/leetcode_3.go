package main

//func main() {
//	s := "abcabcbb"
//	result := lengthOfLongestSubstring(s)
//	fmt.Println(result)
//}

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
