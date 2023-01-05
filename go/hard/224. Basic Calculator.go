package hard

func calculate(s string) int {
	prev := 0
	cur := 0
	sign := 1
	stack := [][2]int{}
	for _, c := range s {
		if c == ' ' {
			continue
		}
		if '0' <= c && c <= '9' {
			cur = cur*10 + int(c-'0')
			continue
		}
		if sign > 0 {
			prev += cur
		} else {
			prev -= cur
		}
		cur = 0
		switch c {
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, [2]int{prev, sign})
			prev = 0
			sign = 1
		case ')':
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last[1] > 0 {
				prev += last[0]
			} else {
				prev = last[0] - prev
			}
		}
	}

	if sign > 0 {
		prev += cur
	} else {
		prev -= cur
	}

	return prev
}
