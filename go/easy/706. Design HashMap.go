package easy

type MyHashMapNode struct {
	key   int
	value int
	next  *MyHashMapNode
}

type MyHashMap struct {
	things []*MyHashMapNode
}

func Constructor() MyHashMap {
	return MyHashMap{
		things: make([]*MyHashMapNode, 100_000),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.hash(key)
	el := this.things[idx]
	if el == nil || el.key == key {
		this.things[idx] = &MyHashMapNode{key: key, value: value}
		return
	}

	for el.next != nil {
		if el.next.key == key {
			el.next.value = value
			return
		}
		el = el.next
	}

	el.next = &MyHashMapNode{key: key, value: value}
}

func (this *MyHashMap) Remove(key int) {
	idx := this.hash(key)
	el := this.things[idx]
	if el == nil {
		return
	}

	if el.key == key || el.next == nil {
		if el.key == key {
			this.things[idx] = nil
		}

		return
	}

	for el.next != nil {
		if el.next.key == key {
			el.next = el.next.next
			return
		}

		el = el.next
	}
}

func (this *MyHashMap) Get(key int) int {
	el := this.things[this.hash(key)]
	if el == nil {
		return -1
	}

	for el != nil {
		if el.key == key {
			return el.value
		}

		el = el.next
	}

	return -1
}

func (this *MyHashMap) hash(x int) int {
	return x % 100_000
}
