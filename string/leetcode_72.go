package main

//func main() {
//	word1, word2 := "horse", "ros"
//	result := minDistance(word1, word2)
//	fmt.Println(result)
//}
//
//func minDistance(word1 string, word2 string) int {
//	table := make([][]int, len(word1)+1)
//	for i := 0; i < len(word1)+1; i++ {
//		table[i] = make([]int, len(word2)+1)
//	}
//	for i := 0; i < len(word1)+1; i++ {
//		table[i][0] = i
//	}
//	for i := 0; i < len(word2)+1; i++ {
//		table[0][i] = i
//	}
//	for i := 1; i < len(word1)+1; i++ {
//		for k := 1; k < len(word2)+1; k++ {
//			if word1[i-1] == word2[k-1] {
//				table[i][k] = table[i-1][k-1]
//			} else {
//				table[i][k] = min(table[i-1][k], table[i][k-1], table[i-1][k-1]) + 1
//			}
//		}
//	}
//	return table[len(word1)][len(word2)]
//}
//func min(a, b, c int) int {
//	if a < b {
//		if a < c {
//			return a
//		} else {
//			return c
//		}
//	} else {
//		if b < c {
//			return b
//		} else {
//			return c
//		}
//	}
//}
