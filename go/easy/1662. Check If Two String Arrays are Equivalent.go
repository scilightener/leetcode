package easy

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	l, r := 0, 0
	i, j := 0, 0
	for {
		if word1[l][i] != word2[r][j] {
			return false
		}
		if len(word1[l]) == i+1 {
			l++
			i = 0
		} else {
			i++
		}

		if len(word2[r]) == j+1 {
			r++
			j = 0
		} else {
			j++
		}

		if l == len(word1) {
			return r == len(word2)
		}
		if r == len(word2) {
			return false
		}
	}
}
