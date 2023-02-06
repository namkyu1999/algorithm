package main

import "errors"

type Stack[T any] struct {
	head  int
	items []T
}

func newStack[T any]() *Stack[T] {
	return &Stack[T]{
		head:  -1,
		items: []T{},
	}
}

func (stack *Stack[T]) push(item T) {
	stack.head += 1
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) pop() (T, error) {
	var result T
	if stack.head > -1 {
		result = stack.items[stack.head]
		stack.items = stack.items[:stack.head]
		stack.head--
		return result, nil
	} else {
		return result, errors.New("no element here")
	}
}

func (stack *Stack[T]) peek() (T, error) {
	var result T
	if stack.head > -1 {
		result = stack.items[stack.head]
		return result, nil
	} else {
		return result, errors.New("no element here")
	}
}

func (stack *Stack[T]) size() int {
	return stack.head + 1
}

func (stack *Stack[T]) isEmpty() bool {
	return stack.head == -1
}
