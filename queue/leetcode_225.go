package main

type MyStack struct {
	items []int
	head  int
}

//func Constructor() MyStack {
//	return MyStack{
//		items: []int{},
//		head:  -1,
//	}
//}
//
//func (this *MyStack) Push(x int) {
//	this.items = append(this.items, x)
//	this.head++
//}
//
//func (this *MyStack) Pop() int {
//	value := this.items[this.head]
//	this.items = this.items[:this.head]
//	this.head--
//	return value
//}
//
//func (this *MyStack) Top() int {
//	return this.items[this.head]
//}
//
//func (this *MyStack) Empty() bool {
//	return this.head < 0
//}
