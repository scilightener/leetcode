package medium

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

func minMutation(start, end string, bank []string) int {
	var set Set = make(map[string]struct{}, len(bank))
	for _, word := range bank {
		set.add(word)
	}

	if !set.contains(end) {
		return -1
	}
	choices := [4]string{"A", "C", "G", "T"}
	ans := 0
	cur, next := []string{start}, []string{}
	var visited Set = make(map[string]struct{})
	for len(cur) > 0 {
		for _, word := range cur {
			if word == end {
				return ans
			}
			for i := range word {
				s := word[:i]
				e := word[i+1:]
				for _, c := range choices {
					new_word := s + c + e
					if set.contains(new_word) && !visited.contains(new_word) {
						next = append(next, new_word)
					}
					visited.add(new_word)
				}
			}
		}

		cur, next = next, cur[:0]
		ans++
	}

	return -1
}
