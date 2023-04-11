package main

import (
	"math"
)

func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
	lastWord1Loc, lastWord2Loc := -1, -1

	shortestDistance := float64(len(wordsDict))
	for i, word := range wordsDict {
		if word == word1 {
			if lastWord2Loc != -1 {
				curDistance := math.Abs(float64(i - lastWord2Loc))
				shortestDistance = math.Min(shortestDistance,
					curDistance,
				)
			}

		} else if word == word2 {
			if lastWord1Loc != -1 {
				curDistance := math.Abs(float64(i - lastWord1Loc))
				shortestDistance = math.Min(shortestDistance,
					curDistance,
				)
			}
		}

		if word == word1 {
			lastWord1Loc = i
		}

		if word == word2 {
			lastWord2Loc = i
		}
	}

	return int(shortestDistance)
}
