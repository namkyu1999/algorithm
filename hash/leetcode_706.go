package main

type MyHashMap struct {
	hash map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		hash: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hash[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if v, ok := this.hash[key]; ok {
		return v
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.hash, key)
}
