package medium

func counter(words []string) map[string]int {
	result := make(map[string]int, len(words))
	for _, word := range words {
		if _, ok := result[word]; !ok {
			result[word] = 1
		} else {
			result[word]++
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestPalindrome(words []string) int {
	cnt := counter(words)
	ln := 0
	for word, freq := range cnt {
		if word[0] == word[1] {
			if ln%4 == 0 || freq%2 == 0 {
				ln += 2 * freq
			} else {
				ln += 2 * (freq - 1)
			}
			continue
		}

		word_rvrsd := string([]byte{word[1], word[0]})
		if freq_rvrsd, ok := cnt[word_rvrsd]; ok {
			d := min(freq, freq_rvrsd)
			ln += 4 * d
			delete(cnt, word_rvrsd)
			delete(cnt, word)
		}
	}

	return ln
}
