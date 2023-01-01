package _go

import "strings"

func wordPattern(pattern string, s string) bool {
	surjection, injection := make(map[string]rune, len(s)), make(map[rune]string, len(pattern))
	splited := strings.Fields(s)
	if len(pattern) != len(splited) {
		return false
	}
	for i, l := range pattern {
		if _, ok := injection[l]; !ok {
			if _, ok := surjection[splited[i]]; ok {
				return false
			}
			injection[l] = splited[i]
			surjection[splited[i]] = l
			continue
		}
		if injection[l] != splited[i] {
			return false
		}
	}
	return true
}

/*
complexity:
	time: O(N) iterate through the pattern
	space: O(N) need bijection map

idea is simple:
	if one letter corresponds to more than one word -> false (line 20)
	if one word corresponds to more than one letter -> false (line 13)

not sure if i used surjection & injection terms properly though~ but who cares right?
*/
