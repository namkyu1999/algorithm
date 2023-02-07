package main

//func main() {
//	input := "bcabc"
//	result := removeDuplicateLetters(input)
//	fmt.Println(result)
//}

func removeDuplicateLetters(s string) string {
	lastIndex := [26]int{}
	seen := [26]bool{}

	for i, c := range []byte(s) {
		lastIndex[c2i(c)] = i
	}

	var res []byte
	for i, current := range []byte(s) {
		currentInteger := c2i(current)
		if seen[currentInteger] {
			continue
		}
		for len(res) > 0 {
			peeked := res[len(res)-1]
			peekedInteger := c2i(peeked)
			if current <= peeked && lastIndex[peekedInteger] >= i && !seen[currentInteger] {
				seen[peekedInteger] = false
				res = res[:len(res)-1]
			} else {
				break
			}
		}

		res = append(res, current)
		seen[currentInteger] = true
	}
	return string(res)
}

func c2i(c byte) int {
	return int(c) - int('a')
}
