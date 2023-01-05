package medium

func largestOverlap(img1 [][]int, img2 [][]int) int {
	ans := 0
	n := len(img1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, calculateOverlap(img1, img2, i, j))
			ans = max(ans, calculateOverlap(img1, img2, i, -j))
			ans = max(ans, calculateOverlap(img1, img2, -i, j))
			ans = max(ans, calculateOverlap(img1, img2, -i, -j))
		}
	}

	return ans
}

func calculateOverlap(img1, img2 [][]int, shift_i, shift_j int) int {
	n := len(img1)
	overlap := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+shift_i < n && j+shift_j < n &&
				i+shift_i > -1 && j+shift_j > -1 {
				overlap += (img1[i+shift_i][j+shift_j] + img2[i][j]) >> 1
			}
		}
	}

	return overlap
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
