package hard

type Set map[rune]struct{}

func (set *Set) Add(e rune) {
	if _, ok := (*set)[e]; ok {
		return
	}
	(*set)[e] = struct{}{}
}

func minWindow(s string, t string) string {
	var tSet Set
	tSet = make(map[rune]struct{})
	for _, val := range t {
		tSet.Add(val)
	}

	var inds []int
	for i, val := range s {
		if _, ok := tSet[val]; ok {
			inds = append(inds, i)
		}
	}

	count := Counter(t)
	c := len(count)
	min, l, bl, br := 100000, 0, 0, 0
	for _, r := range inds {
		count[s[r]] -= 1
		if count[s[r]] == 0 {
			c--
		}
		if c == 0 {
			for ; count[s[inds[l]]] < 0; l++ {
				count[s[inds[l]]]++
			}

			if r-inds[l] < min {
				min, bl, br = r-inds[l], inds[l], r+1
			}
		}
	}

	return s[bl:br]
}

func Counter(s string) map[byte]int {
	counter := make(map[byte]int)
	for i := range s {
		if _, ok := counter[s[i]]; !ok {
			counter[s[i]] = 0
		}
		counter[s[i]]++
	}
	return counter
}
