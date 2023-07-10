package main

//func main() {
//	s := "abcdefg"
//	shift := [][]int{
//		{1, 1},
//		{1, 1},
//		{0, 2},
//		{1, 3},
//	}
//	answer := stringShift(s, shift)
//	fmt.Println(answer)
//}

func stringShift(s string, shift [][]int) string {
	shifted := 0
	for _, v := range shift {
		if v[0] == 0 {
			shifted -= v[1]
		} else {
			shifted += v[1]
		}
	}
	if shifted < 0 {
		shifted = -shifted
		shifted = shifted % len(s)
		s = s[shifted:] + s[:shifted]
	} else {
		shifted = shifted % len(s)
		n := len(s) - shifted
		s = s[n:] + s[:n]
	}
	return s
}
