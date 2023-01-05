package medium

func maxLength(arr []string) int {
	bs := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		k := 0
		for _, val := range arr[i] {
			if k&(1<<(val-'a')) != 0 {
				k = -1
				break
			}
			k += 1 << (val - 'a')
		}

		bs[i] = k
	}

	max_len := 0
	for k := 1; k < (1 << len(bs)); k++ {
		cur_len := 0
		word := 0
		for c := 0; (k >> c) != 0; c++ {
			if (k>>c)%2 == 1 {
				if word&bs[c] != 0 || bs[c] == -1 {
					cur_len = 0
					break
				}
				word |= bs[c]
				cur_len += len(arr[c])
			}
		}
		if max_len < cur_len {
			max_len = cur_len
		}
	}

	return max_len
}
