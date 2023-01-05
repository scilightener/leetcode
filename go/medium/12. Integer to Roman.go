package medium

import (
	"sort"
	"strings"
)

func intToRoman(num int) string {
	conv := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	var sb strings.Builder
	nums := make([]int, len(conv))
	i := 0
	for k := range conv {
		nums[i] = k
		i++
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for _, n := range nums {
		for ; num >= n; num -= n {
			sb.WriteString(conv[n])
		}
		if num == 0 {
			break
		}
	}
	return sb.String()
}
