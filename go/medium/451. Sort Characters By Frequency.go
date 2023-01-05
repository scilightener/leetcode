package medium

import "strings"

func frequencySort(s string) string {
	counter := make(map[rune]int, 52)
	arr := make([][]rune, len(s)+1)
	for _, l := range s {
		if _, ok := counter[l]; !ok {
			counter[l] = 1
		} else {
			counter[l]++
		}
	}

	for k, v := range counter {
		arr[len(s)-v] = append(arr[len(s)-v], k)
	}

	var sb strings.Builder
	for i, runes := range arr {
		for _, l := range runes {
			sb.WriteString(strings.Repeat(string(l), len(s)-i))
		}
	}

	return sb.String()
}
