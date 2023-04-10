package main

//func main() {
//	result := fib(10)
//	fmt.Println(result)
//}

var hash = make(map[int]int)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	if v, ok := hash[n]; ok {
		return v
	} else {
		hash[n] = fib(n-1) + fib(n-2)
		return hash[n]
	}
}
