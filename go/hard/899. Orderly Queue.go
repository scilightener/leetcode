package hard

import (
	"sort"
	"strings"
)

func orderlyQueue(s string, k int) string {
	if k == 1 {
		return minRotate(s)
	} else {
		return sortedString(s)
	}
}

func sortedString(s string) string {
	res := []byte(s)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return string(res)
}

func minRotate(s string) string {
	res := s
	for i := range s {
		new_s := s[i:] + s[:i]
		if strings.Compare(new_s, res) < 0 {
			res = new_s
		}
	}

	return res
}
