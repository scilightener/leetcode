package easy

func detectCapitalUse(word string) bool {
	if word[0] < 97 {
		return allLettersSameCase(word[1:])
	}
	return allLettersSameCase(word)
}

func allLettersSameCase(word string) bool {
	if len(word) == 0 {
		return true
	}
	isCapital := word[0] < 97
	for _, l := range word {
		if isCapital != (l < 97) {
			return false
		}
	}
	return true
}

/*
complexity:
	time: O(N)
	space: O(1)

the idea is pretty straight forward
first, check the first letter:
	if it is capital, we ensure that the left part of the word is either full lowercase or uppercase
	if not, check that the whole word is lowercase
*/
