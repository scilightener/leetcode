package medium

import "strings"

func reverseWords(s string) string {
	splitted := strings.Fields(s)
	for i, j := 0, len(splitted)-1; i < j; i, j = i+1, j-1 {
		splitted[i], splitted[j] = splitted[j], splitted[i]
	}

	return strings.Join(splitted, " ")
}
