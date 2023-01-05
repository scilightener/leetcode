package easy

func guess(num int) int

func guessNumber(n int) int {
	var mid, resp int
	low, high := 1, n
	for low <= high {
		mid = low + (high-low)/2
		resp = guess(mid)
		if resp == 0 {
			return mid
		}
		if resp > 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return mid
}
