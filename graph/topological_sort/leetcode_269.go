package main

//func main() {
//	words := []string{"z", "z"}
//	result := alienOrder(words)
//	fmt.Println(result)
//}

func alienOrder(words []string) string {
	inOrders := map[byte]map[byte]bool{}
	result := ""
	var queue []byte
	inOrdersCount := map[byte]int{}
	for _, word := range words {
		for _, w := range []byte(word) {
			inOrdersCount[w] = 0
		}
	}
	if len(words) > 1 {
		for i := 1; i < len(words); i++ {
			prev, current := words[i-1], words[i]
			var minCount int
			if len(prev) < len(current) {
				minCount = len(prev)
			} else {
				minCount = len(current)
			}
			isBreak := false
			for k := 0; k < minCount; k++ {
				if prev[k] != current[k] {
					if inOrders[prev[k]] == nil {
						inOrders[prev[k]] = map[byte]bool{}
					}
					inOrders[prev[k]][current[k]] = true
					isBreak = true
					break
				}
			}
			if !isBreak && len(prev) > len(current) {
				return ""
			}
		}
	}
	for _, v := range inOrders {
		for post := range v {
			inOrdersCount[post]++
		}
	}
	for k, v := range inOrdersCount {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result += string(current)
		for k := range inOrders[current] {
			inOrdersCount[k]--
			if inOrdersCount[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	if len(inOrdersCount) != len(result) {
		return ""
	}
	return result
}
