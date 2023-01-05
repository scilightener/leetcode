package medium

func numTilings(n int) int {
	dp0, dp1, dp2, dp3 := 1, 0, 0, 1
	for ; n > 1; n-- {
		d0, d1, d2, d3 := dp0, dp1, dp2, dp3
		dp0 = (d0 + d1 + d2 + d3) % 1000000007
		dp1 = (d2 + d3) % 1000000007
		dp2 = (d1 + d3) % 1000000007
		dp3 = d0 % 1000000007
	}
	return dp0
}

/*
complexity:
	time: O(N) just one for loop
	space: O(1) just 4 vars

the idea is dp on subsets. well, hard to explain in few lines...
let's say there are several states of a column (2 x 1 grid):
	0 - all two cells of column are free
	1 - top cell is taken by domino or tromino, bottom is free
	2 - bottom cell is taken, top is free
	3 - both taken
now consider we know the number of ways to tile board with n-1 columns (2 x (n-1))
how do we tile the n-th column (may be not exactly accurate, some may stick out, but necessarily fulfill all the current column cells)?

well, our current state is the 0-th state, we may place dominos and trominos 4 ways: 2 horizontal doms, 1 vertical dom, 1 trom with the top sticking out, 1 trom with the bottom sticking out. it will produce states for the next column as follows: 3, 0, 1, 2

if we happened to start from the 1-st state, we may place 1 horizontal dom or 1 trom. it will give us states 2 and 3 accordingly

the same for state 2: it will give 1 and 3 states

if the state is 3, then we can't do anything. produce state 1 for the next

well, if you understood all the text above, then the code shouldn't seem so hard to you :)
*/
