package main

func IsValidSubsequence(array []int, sequence []int) bool {
	i, k := 0, 0
	for i < len(array) && k < len(sequence) {
		if array[i] == sequence[k] {
			k++
		}
		i++
	}
	return k == len(sequence)
}

//func main() {
//	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
//	sequence := []int{1, 6, -1, 10}
//	result := IsValidSubsequence(array, sequence)
//	fmt.Println(result)
//}
