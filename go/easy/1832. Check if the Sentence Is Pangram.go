package easy

func checkIfPangram(sentence string) bool {
	var alph [26]bool
	for _, val := range sentence {
		alph[val-'a'] = true
	}
	for _, val := range alph {
		if !val {
			return false
		}
	}
	return true
}
