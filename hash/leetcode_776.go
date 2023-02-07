package main

//func main() {
//	jewels := "aA"
//	stones := "aAAbbbb"
//	result := numJewelsInStones(jewels, stones)
//	fmt.Println(result)
//}

func numJewelsInStones(jewels string, stones string) int {
	hash := make(map[byte]int)
	for _, stone := range []byte(stones) {
		hash[stone] += 1
	}
	result := 0
	for _, jewel := range []byte(jewels) {
		result += hash[jewel]
	}
	return result
}
