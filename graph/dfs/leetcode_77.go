package main

//func main() {
//	n, k := 4, 2
//	result := combine(n, k)
//	fmt.Println(result)
//}

func combine(n int, k int) [][]int {
	var result [][]int
	var current []int
	cdfs(&result, current, 1, n, k)
	return result
}

func cdfs(result *[][]int, current []int, start, n, k int) {
	if k == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for i := start; i <= n-k+1; i++ {
		current = append(current, i)
		cdfs(result, current, i+1, n, k-1)
		current = current[:len(current)-1]
	}
}
