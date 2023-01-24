package main

//func main() {
//	s := []byte("hello")
//	reverseString(s)
//	fmt.Println(string(s))
//}

func reverseString(s []byte) {
	lastIndex := len(s) - 1
	for i := 0; i <= lastIndex/2; i++ {
		tmp := s[i]
		s[i] = s[lastIndex-i]
		s[lastIndex-i] = tmp
	}
}
