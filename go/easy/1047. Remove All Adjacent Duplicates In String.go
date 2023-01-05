package easy

func removeDuplicates(s string) string {
	stack := make([]rune, 0, len(s))

	for _, letter := range s {
		if len(stack) > 0 && stack[len(stack)-1] == letter {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, letter)
	}

	return string(stack)
}
