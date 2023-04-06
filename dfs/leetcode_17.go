package main

//func main() {
//	input := "23"
//	result := letterCombinations(input)
//	fmt.Println(result)
//}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	ldfs(&result, digits, "")
	return result
}

func ldfs(result *[]string, digits string, current string) {
	hash := map[byte]string{
		byte('2'): "abc",
		byte('3'): "def",
		byte('4'): "ghi",
		byte('5'): "jkl",
		byte('6'): "mno",
		byte('7'): "pqrs",
		byte('8'): "tuv",
		byte('9'): "wxyz",
	}
	if len(digits) > 0 {
		for _, v := range hash[digits[0]] {
			ldfs(result, digits[1:], current+string(v))
		}
	} else {
		if current != "" {
			*result = append(*result, current)
		}
	}
}
