package medium

import "math/big"

func groupAnagrams(strs []string) [][]string {
	primes := [26]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	var result [][]string
	anams := make(map[string][]string)
	for _, s := range strs {
		h := big.NewInt(1)
		for _, l := range s {
			h.Mul(h, big.NewInt(primes[l-'a']))
		}
		hString := h.String()
		if _, ok := anams[hString]; !ok {
			anams[hString] = []string{}
		}
		anams[hString] = append(anams[hString], s)
	}

	for _, val := range anams {
		result = append(result, val)
	}
	return result
}
