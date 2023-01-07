package medium

func canCompleteCircuit(gas []int, cost []int) int {
	n, tank, debt, debt_idx := len(gas), 0, 0, -1
	for i := 0; i < n; i++ {
		tank += gas[i] - cost[i]
		if tank < debt {
			debt = tank
			debt_idx = i
		}
	}
	if tank < 0 {
		return -1
	}
	return (debt_idx + 1) % n
}

/*
complexity:
	time: O(N) only need to iterate through the input once
	space: O(1)

well it's more like math problem to see why we always have an answer when sum(gas) >= sum(cost)
let's imagine that our car can run with a negative fuel value
let's start at the starting point and see what the maximum gasoline debt we can get in this case
now that we have reached the end of the route, if we have a negative value of gasoline left, then it is impossible to pass this route
otherwise, we say the following:
since from the moment when we fixed the maximum gasoline debt, our fuel supply only increased (at the end it became >= 0), it means that on this section of the road it would be possible to stock up on gasoline for the rest of the ride
this means that if we started our ride at the next gas station after the maximum gasoline debt station, we would be able to drive the entire route
*/
