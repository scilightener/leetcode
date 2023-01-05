package medium

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	counter1, counter2 := map[byte]int{}, map[byte]int{}
	for i := range word1 {
		if _, ok := counter1[word1[i]]; !ok {
			counter1[word1[i]] = 1
		} else {
			counter1[word1[i]]++
		}
		if _, ok := counter2[word2[i]]; !ok {
			counter2[word2[i]] = 1
		} else {
			counter2[word2[i]]++
		}
	}
	rev1, rev2 := map[int]int{}, map[int]int{}
	for k, v := range counter1 {
		if _, ok := counter2[k]; !ok {
			return false
		}
		if _, ok := rev1[v]; !ok {
			rev1[v] = 1
		} else {
			rev1[v]++
		}
	}
	for k, v := range counter2 {
		if _, ok := counter1[k]; !ok {
			return false
		}
		if _, ok := rev2[v]; !ok {
			rev2[v] = 1
		} else {
			rev2[v]++
		}
	}
	for k := range rev1 {
		if _, ok := rev2[k]; !ok {
			return false
		}
		if rev1[k] != rev2[k] {
			return false
		}
	}

	return true
}
