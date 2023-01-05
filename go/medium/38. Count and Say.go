package medium

import "strconv"

func countAndSay(n int) string {
	ans := "1"
	for i := 1; i < n; i++ {
		ans = rle(ans)
	}
	return ans
}

func rle(s string) string {
	var cnt int = 1
	if len(s) == 1 {
		return "1" + s
	}
	i := 1
	for {
		if i > len(s) {
			break
		}
		if i == len(s) || s[i] != s[i-1] {
			cntStr := strconv.Itoa(cnt)
			s = s[:i-cnt] + cntStr + s[i-1:]
			i += 2 + len(cntStr) - cnt
			cnt = 1
			continue
		}
		cnt++
		i++
	}
	return s
}
