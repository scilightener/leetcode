package hard

import "sort"

type time struct {
	plant int
	grow  int
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	times := make([]time, len(plantTime))
	for i := range plantTime {
		times[i] = time{plantTime[i], growTime[i]}
	}

	sort.Slice(times, func(i, j int) bool { return times[i].grow > times[j].grow })
	curPlant, res := 0, 0
	for _, t := range times {
		curPlant += t.plant
		if res < curPlant+t.grow {
			res = curPlant + t.grow
		}
	}

	return res
}
