package easy

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n < 4 {
		return true
	}
	var allowed = [3]int{2, 3, 5}
	for i := 0; i < 3; i++ {
		factor := allowed[i]
		for n%factor == 0 {
			n /= factor
		}
	}

	return n == 1 || n == -1
}
