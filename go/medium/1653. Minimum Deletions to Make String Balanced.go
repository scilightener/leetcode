package medium

func minimumDeletions(s string) int {
	ans, countB := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			countB++
			continue
		}
		ans = min(ans+1, countB)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
