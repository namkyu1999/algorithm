package main

func maximumWealth(accounts [][]int) int {
	max := -1
	for _, v := range accounts {
		tmp := 0
		for _, k := range v {
			tmp += k
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
