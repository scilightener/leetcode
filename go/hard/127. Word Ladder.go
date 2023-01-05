package hard

type Set map[string]struct{}

func (s Set) contains(k string) bool {
	if _, ok := s[k]; ok {
		return true
	}
	return false
}

func (s *Set) add(k string) {
	if (*s).contains(k) {
		return
	}
	(*s)[k] = struct{}{}
}

func ladderLength(start, end string, wordList []string) int {
	var (
		set1  Set = make(map[string]struct{})
		set2  Set = make(map[string]struct{})
		words Set = make(map[string]struct{})
	)

	set1.add(start)
	set2.add(end)
	for _, word := range wordList {
		words.add(word)
	}

	if !words.contains(end) {
		return 0
	}
	return bfs(words, set1, set2, 1)
}
func bfs(words, set1, set2 Set, level int) int {
	if len(set1) == 0 {
		return 0
	}
	if len(set1) > len(set2) {
		return bfs(words, set2, set1, level)
	}

	for word := range set1 {
		delete(words, word)
	}

	for word := range set2 {
		delete(words, word)
	}

	var set Set = make(map[string]struct{})
	for word := range set1 {
		for i := range word {
			b := []rune(word)
			for ch := 'a'; ch <= 'z'; ch++ {
				b[i] = ch
				newWord := string(b)
				if set2.contains(newWord) {
					return level + 1
				}

				if words.contains(newWord) {
					set.add(newWord)
				}
			}
		}
	}

	return bfs(words, set, set2, level+1)
}
