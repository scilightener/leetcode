package easy

func reverseVowels(s string) string {
	vowels := map[byte]struct{}{
		'a': {},
		'A': {},
		'e': {},
		'E': {},
		'i': {},
		'I': {},
		'o': {},
		'O': {},
		'u': {},
		'U': {}}
	buf := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for ; l < r; l++ {
			if _, ok := vowels[buf[l]]; ok {
				break
			}
		}

		for ; l < r; r-- {
			if _, ok := vowels[buf[r]]; ok {
				break
			}
		}

		if l >= r {
			break
		}
		buf[l], buf[r] = buf[r], buf[l]
	}

	return string(buf)
}
