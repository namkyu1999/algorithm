package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	hash := make(map[string]int)
	for i := 0; i < len(results); i++ {
		hash[competitions[i][HOME_TEAM_WON-results[i]]] += 3
	}
	maxValue, winner := -1, ""
	for k, v := range hash {
		if v >= maxValue {
			winner = k
			maxValue = v
		}
	}
	return winner
}

//func main() {
//	competitions := [][]string{
//		{"HTML", "C#"},
//		{"C#", "Python"},
//		{"Python", "HTML"},
//	}
//	results := []int{0, 1, 1}
//	result := TournamentWinner(competitions, results)
//	fmt.Println(result)
//}
