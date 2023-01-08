package hard

func maxPoints(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		vectors := map[[2]int]int{}
		for j := i + 1; j < len(points); j++ {
			dx, dy := points[i][0]-points[j][0], points[i][1]-points[j][1]
			g := gcd(dx, dy)
			dx /= g
			dy /= g
			vector := [2]int{dx, dy}
			if _, ok := vectors[vector]; !ok {
				vectors[vector] = 0
			}
			vectors[vector]++
			ans = max(ans, vectors[vector])
		}
	}
	return ans + 1
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
complexity:
	time: O(N*N*logN) for every point we calculate gcd for every other point in input
	space: O(N) for storing the vectors map

for each point we count the vectors to all the others
but to work with these vectors, you need to bring them to such a form that their coordinates are coprime numbers
to do this, we consider the GCD for these coordinates and divide them by it
in the map, we increment the value of the given vector, thereby showing that we have found another point lying on the line given by this vector applied to the current point
now the answer is just the maximum value of the given dictionary for all points + 1 bc we didn't consider the initial point in our calculations

upd: oh, just found out, gcd is not necessary for this problem. direction vector can be given by the slope of the line e.g. ratio y/x
time complexity could be only O(N^2)
*/
