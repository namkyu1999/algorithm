package main

type MyCircularQueue struct {
	head    *node
	rear    *node
	size    int
	maxsize int
}

type node struct {
	val  int
	next *node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size:    0,
		maxsize: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	tmp := &node{val: value, next: this.head}

	if this.head == nil && this.rear == nil {
		this.head = tmp
		this.rear = tmp
		this.head.next = this.rear
		this.rear.next = this.head
	} else {
		prev := this.rear
		prev.next = tmp
		this.rear = tmp
	}
	// fmt.Println(this)
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	if this.head == this.rear {
		this.head = nil
		this.rear = nil
	} else {
		this.head = this.head.next
		this.rear.next = this.head
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.head == nil {
		return -1
	}
	return this.head.val
}

func (this *MyCircularQueue) Rear() int {
	if this.rear == nil {
		return -1
	}
	return this.rear.val

}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.size == this.maxsize {
		return true
	}
	return false
}
