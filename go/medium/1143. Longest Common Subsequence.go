package medium

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	lcs := make([]int, len(text2)+1)
	prevLcs := make([]int, len(text2)+1)

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				lcs[j] = 1 + prevLcs[j+1]
			} else {
				lcs[j] = max(lcs[j+1], prevLcs[j])
			}
		}

		copy(prevLcs, lcs)
	}

	return lcs[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
