package medium

func maxProfit(prices []int) int {
	buy1, buy2, sell1, sell2 := -1000, 0, 0, 0
	for _, p := range prices {
		buy1, buy2 = max(sell2-p, buy2), buy1
		sell1, sell2 = max(buy1+p, sell2), sell1
	}
	return sell1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
complexity:
	time: O(N) simple
	space: O(1) simple

the idea is dp. let me explain
buy1 is maxProfit if we ended up yesterday with 1 stock on our hands
buy2 is the same but for the day before yesterday
sell1 is maxProfit if we ended up yesterdat with 0 stocks on our hands
sell2 likewise

now, how we update our vars? well,
buy1: we can buy after the sell two days ago (sell2-p) or do nothing (buy2)
sell1: we can sell the only stock that is on our hands (buy1+p) or do nothing (sell2)
don't forget to update buy2 & sell2 as the new day has come!!

at the end, return the max of those, which is obviously sell1
*/
