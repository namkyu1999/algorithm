package main

//func main() {
//	s := []byte("hello")
//	reverseString(s)
//	fmt.Println(string(s))
//}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
