package main

//func main() {
//	candidates := []int{2, 3, 6, 7}
//	target := 7
//	result := combinationSum(candidates, target)
//	fmt.Println(result)
//}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int
	csDfs(current, candidates, &result, 0, target, 0)
	return result
}

func csDfs(current, candidates []int, result *[][]int, sum, target, index int) {
	if sum == target {
		*result = append(*result, append([]int{}, current...))
		return
	} else if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		csDfs(append(current, candidates[i]), candidates, result, sum+candidates[i], target, i)
	}
}
