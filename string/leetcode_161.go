package main

//func main() {
//	s, t := "a", "ac"
//	result := isOneEditDistance(s, t)
//	fmt.Println(result)
//}

func isOneEditDistance(s string, t string) bool {
	length := len(s) - len(t)
	sRune, tRune := []rune(s), []rune(t)
	diffCount := 0
	if length == 0 {
		// replace
		for i := 0; i < len(sRune); i++ {
			if sRune[i] != tRune[i] {
				if diffCount > 0 {
					return false
				}
				diffCount++
			}
		}
		if diffCount == 0 {
			return false
		}
		return true
	} else if length == 1 {
		// delete
		if len(t) == 0 {
			return true
		}
		k := 0
		for i := 0; i < len(sRune) && k < len(tRune); i++ {
			if sRune[i] != tRune[k] {
				if diffCount > 0 {
					return false
				}
				diffCount++
			} else {
				k++
			}
		}
		return true
	} else if length == -1 {
		// insert
		if len(s) == 0 {
			return true
		}
		k := 0
		for i := 0; i < len(tRune) && k < len(sRune); i++ {
			if tRune[i] != sRune[k] {
				if diffCount > 0 {
					return false
				}
				diffCount++
			} else {
				k++
			}
		}
		return true
	}
	return false
}
