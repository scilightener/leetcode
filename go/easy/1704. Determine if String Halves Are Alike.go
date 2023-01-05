package easy

func halvesAreAlike(s string) bool {
	vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}
	string1 := s[:len(s)/2]
	string2 := s[len(s)/2:]

	count1 := 0
	count2 := 0

	for i := 0; i < len(string1); i++ {
		if _, ok := vowels[string1[i]]; ok {
			count1++
		}
		if _, ok := vowels[string2[i]]; ok {
			count2++
		}
	}

	return count1 == count2
}
