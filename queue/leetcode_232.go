package main

type MyQueue struct {
	items []int
}

//
//func Constructor() MyQueue {
//	return MyQueue{items: []int{}}
//}
//
//func (this *MyQueue) Push(x int) {
//	this.items = append(this.items, x)
//}
//
//func (this *MyQueue) Pop() int {
//	value := this.items[0]
//	this.items = this.items[1:]
//	return value
//}
//
//func (this *MyQueue) Peek() int {
//	return this.items[0]
//}
//
//func (this *MyQueue) Empty() bool {
//	return len(this.items) == 0
//}
