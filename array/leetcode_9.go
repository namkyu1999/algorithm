package main

//func main() {
//	result := isPalindrome(90)
//	fmt.Println(result)
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	arr := make([]int, 0)
	n := -1
	for ; x > 0; n++ {
		arr = append(arr, x%10)
		x = x / 10

	}
	for i := 0; i <= n/2; i++ {
		if arr[i] != arr[n-i] {
			return false
		}
	}
	return true
}
