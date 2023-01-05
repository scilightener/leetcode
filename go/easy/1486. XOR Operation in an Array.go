package easy

func xorOperation(n int, start int) int {
	last := start + 2*(n-1)
	if start%4 <= 1 {
		switch n % 4 {
		case 0:
			return 0
		case 1:
			return last
		case 2:
			return 2
		case 3:
			return 2 ^ last
		}
	} else {
		switch n % 4 {
		case 0:
			return start ^ last ^ 2
		case 1:
			return start
		case 2:
			return start ^ last
		case 3:
			return start ^ 2
		}
	}
	return -1
}
