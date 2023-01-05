package easy

func minDeletionSize(strs []string) int {
	ans := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				ans++
				break
			}
		}
	}
	return ans
}

/*
complexity:
	time: O(NK) where N is the length of strs and K is the length of each string in strs
	space: O(1) only need to store ans

algorithm is pretty straight forward:
	for every letter in string we loop over strs and check if the current symbol in the current string is more than or equal to the current symbol in the previous string. If not, mark this column to delete
*/
