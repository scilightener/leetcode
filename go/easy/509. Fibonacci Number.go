package easy

func fib(n int) int {
	if n < 2 {
		return n
	}
	last, cur := 0, 1
	for n > 1 {
		last, cur = cur, last+cur
		n--
	}
	return cur
}
