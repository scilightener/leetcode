package medium

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	intersection := getIntersectionArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)

	areaA := getArea(ax1, ay1, ax2, ay2)
	areaB := getArea(bx1, by1, bx2, by2)
	if areaA == intersection {
		return areaB
	}
	if areaB == intersection {
		return areaA
	}
	return areaA + areaB - intersection
}

func getIntersectionArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	left := max(ax1, bx1)
	right := min(ax2, bx2)
	top := min(ay2, by2)
	bottom := max(ay1, by1)
	if left <= right && bottom <= top {
		return getArea(left, bottom, right, top)
	}
	return 0
}

func getArea(x1, y1, x2, y2 int) int {
	return (x2 - x1) * (y2 - y1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
