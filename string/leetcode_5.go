package main

//func main() {
//	s := "aacabdkacaa"
//	answer := longestPalindrome(s)
//	fmt.Println(answer)
//}

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		// Odd case
		palindromeOdd := expandAroundCenter(s, i, i)
		// Even case
		palindromeEven := expandAroundCenter(s, i, i+1)

		if len(palindromeOdd) > len(longest) {
			longest = palindromeOdd
		}
		if len(palindromeEven) > len(longest) {
			longest = palindromeEven
		}
	}
	return longest
}

func expandAroundCenter(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
