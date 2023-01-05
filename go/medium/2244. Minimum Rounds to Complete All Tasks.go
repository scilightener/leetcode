package medium

func minimumRounds(tasks []int) int {
	tasksCount := map[int]int{}
	for _, task := range tasks {
		if _, ok := tasksCount[task]; !ok {
			tasksCount[task] = 0
		}
		tasksCount[task]++
	}
	ans := 0
	for _, cnt := range tasksCount {
		if cnt == 1 {
			return -1
		}
		ans += minimumRoundsForTask(cnt)
	}
	return ans
}

func minimumRoundsForTask(task int) int {
	rem := task % 3
	return task/3 + rem&1 + rem&2/2
}

/*
complexity:
	time: O(N) for creating count map + iterate through it
	space: O(N) for this count map

for each unique task we count how many of its is in the tasks
then just iterate through this count map and add up the number of rounds required for this type of task
by that crazy formula xd
well, actually this is just a vErY cLeVeR wAy to write smth like:

	if rem > 0 {
		return task/3 + 1
	}
	return task/3

bc if we have smth like (...) + 1 where (...) - smth divisible by 3
we may say it's (...) + 3 + 1 <=> (...) + 2 + 2 -> and now can complete all the tasks with one more operation
	!!!but!!! only if (...) > 0. if not, then it's impossible to complete and we return -1
and if we have smth like (...) + 2, then we obviously can complete it
*/
