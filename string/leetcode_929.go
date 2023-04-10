package main

import (
	"strings"
)

//func main() {
//	emails := []string{
//		"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com",
//	}
//	result := numUniqueEmails(emails)
//	fmt.Println(result)
//}

func numUniqueEmails(emails []string) int {
	hash := make(map[string]byte)
	for _, email := range emails {
		address := strings.Split(email, "@")
		company := address[1]
		ignorePlus := strings.Split(address[0], "+")[0]
		hash[strings.ReplaceAll(ignorePlus, ".", "")+"@"+company] = 1
	}
	return len(hash)
}
