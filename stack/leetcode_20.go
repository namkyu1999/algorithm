package main

//func main() {
//	input := "]"
//	result := isValid(input)
//	fmt.Println(result)
//}

func isValid(s string) bool {
	table := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := newStack[rune]()
	runeSlice := []rune(s)
	for _, r := range runeSlice {
		if opp, ok := table[r]; ok {
			if peeked, err := stack.peek(); err == nil {
				if peeked == opp {
					stack.pop()
				} else {
					stack.push(r)
				}
			} else {
				stack.push(r)
			}
		} else {
			stack.push(r)
		}
	}
	if stack.isEmpty() {
		return true
	} else {
		return false
	}
}
