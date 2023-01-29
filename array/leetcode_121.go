package main

//func main() {
//	prices := []int{1, 2}
//	answer := maxProfit(prices)
//	fmt.Println(answer)
//}

func maxProfit(prices []int) int {
	minPrice, answer := prices[0], 0
	for _, v := range prices {
		if minPrice > v {
			minPrice = v
		} else {
			if answer < v-minPrice {
				answer = v - minPrice
			}
		}
	}
	return answer
}
