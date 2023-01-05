package medium

import "math/rand"

type RandomizedSet struct {
	set   map[int]int
	elems []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int, 0), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := (*this).set[val]; ok {
		return false
	}
	(*this).elems = append((*this).elems, val)
	(*this).set[val] = len(this.elems) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := (*this).set[val]; !ok {
		return false
	}
	idx := this.set[val]
	delete((*this).set, val)
	if idx != len(this.elems)-1 {
		(*this).set[this.elems[len(this.elems)-1]] = idx
	}
	(*this).elems[idx], (*this).elems[len(this.elems)-1] = (*this).elems[len(this.elems)-1], (*this).elems[idx]
	(*this).elems = (*this).elems[:len(this.elems)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.elems[rand.Intn(len(this.elems))]
}
