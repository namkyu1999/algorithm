package main

func romanToInt(s string) int {
	hash := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	var prev rune
	for _, v := range []rune(s) {
		if (v == 'V' || v == 'X') && prev == 'I' {
			result -= 2 * hash['I']
		}
		if (v == 'L' || v == 'C') && prev == 'X' {
			result -= 2 * hash['X']
		}
		if (v == 'D' || v == 'M') && prev == 'C' {
			result -= 2 * hash['C']
		}
		prev = v
		result += hash[v]
	}
	return result
}
