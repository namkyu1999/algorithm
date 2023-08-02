package main

import (
	"sort"
)

//func main() {
//	tickets := [][]string{
//		{"MUC", "LHR"},
//		{"JFK", "MUC"},
//		{"SFO", "SJC"},
//		{"LHR", "SFO"},
//	}
//	result := findItinerary(tickets)
//	fmt.Println(result)
//}

func findItinerary(tickets [][]string) []string {
	// make a graph, map to keep tickets from source => []destinations
	adjancency := make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		adjancency[tickets[i][0]] = append(adjancency[tickets[i][0]], tickets[i][1])
	}

	// constraint is to give preferrence to lexicographically sorted solution, lets sort all destinations in asc order
	// we do it in a separate loop and not the above loop because we should call sort once for a list of destinations
	for i := 0; i < len(tickets); i++ {
		sort.Strings(adjancency[tickets[i][0]])
	}

	itinerary := []string{}

	// dfs function made inside this function only because, it is much cleaner this way
	// you need to pass less variables around with the recursive DFS calls
	var dfs func(source string) bool
	dfs = func(source string) bool {
		itinerary = append(itinerary, source)

		// we only want to check solutions only for totalChances, if solution
		// is not found in that, means the solution is somewhere else and we backtrack!
		totalChances := len(adjancency[source])
		chances := 0

		// find all the other possible destinations
		for len(adjancency[source]) > 0 && chances < totalChances {
			destination := adjancency[source][0]
			adjancency[source] = adjancency[source][1:]
			if !dfs(destination) { // backtrack, destination is not a possible destination
				chances += 1
				itinerary = itinerary[:len(itinerary)-1]
				adjancency[source] = append(adjancency[source], destination)
			} else { // hurrah! we found our solution
				return true
			}
		}

		// so when the iteneary is made up of all tickets we have found our solution else not
		// if not we trigger backtrack!
		if len(itinerary)-1 != len(tickets) {
			return false
		}

		// solution found
		return true
	}

	dfs("JFK")

	return itinerary
}

func findItinerary2(tickets [][]string) []string {
	adj := map[string][]string{}
	for _, ticket := range tickets {
		adj[ticket[0]] = append(adj[ticket[0]], ticket[1])
	}
	for _, list := range adj {
		sort.Strings(list)
	}
	ret := []string{}
	addNext("JFK", adj, 0, len(tickets), &ret)
	// for next != "" {
	//     curr := next
	//     ret = append(ret, curr)
	//     next = ""
	//     list := adj[curr]
	//     if len(list) == 0 {
	//         break
	//     }
	//     next = list[0]
	//     list = list[1:]
	//     adj[curr]=list
	//     count++
	// }
	return ret
}

func addNext(curr string, adj map[string][]string, count, target int, ret *[]string) bool {
	list := adj[curr]
	if len(list) == 0 {
		if count == target {
			*ret = append(*ret, curr)
			return true
		}
		return false
	}
	*ret = append(*ret, curr)
	for i := range list {
		tmp := list[i]
		list = append(list[:i], list[i+1:]...)
		adj[curr] = list
		if addNext(tmp, adj, count+1, target, ret) {
			return true
		}
		list = append(list[:i+1], list[i:]...)
		list[i] = tmp
		adj[curr] = list
	}
	*ret = (*ret)[:len(*ret)-1]
	return false
}
