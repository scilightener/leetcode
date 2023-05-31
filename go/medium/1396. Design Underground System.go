package medium

type StringTuple2 struct {
	Item1 string
	Item2 string
}

type UserState struct {
	Station string
	Time    int
}

type StationState struct {
	Sum   int
	Count int
}

type UndergroundSystem struct {
	lastUsersState map[int]UserState
	stationsStates map[StringTuple2]StationState
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		lastUsersState: make(map[int]UserState),
		stationsStates: make(map[StringTuple2]StationState),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.lastUsersState[id] = UserState{Station: stationName, Time: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if state, ok := this.lastUsersState[id]; ok {
		key := StringTuple2{Item1: state.Station, Item2: stationName}
		if stationsInfo, ok := this.stationsStates[key]; ok {
			this.stationsStates[key] = StationState{Sum: stationsInfo.Sum + t - state.Time, Count: stationsInfo.Count + 1}
		} else {
			this.stationsStates[key] = StationState{Sum: t - state.Time, Count: 1}
		}
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	state := this.stationsStates[StringTuple2{Item1: startStation, Item2: endStation}]
	return float64(state.Sum) / float64(state.Count)
}
