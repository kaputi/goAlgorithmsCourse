package main

import (
	"errors"
	"math"
)

type StackNode struct {
	value int
	prev  *StackNode
}

type Stack struct {
	len  int
	head *StackNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (stk *Stack) Push(value int) {

	node := &StackNode{value: value}

	stk.len++

	if stk.head == nil {
		stk.head = node
		return
	}

	node.prev = stk.head
	stk.head = node
}

func (stk *Stack) Pop() (int, error) {

	if stk.len == 0 || stk.head == nil {
		return -1, errors.New("Stack is empty")
	}
	stk.len = int(math.Max(0, float64(stk.len-1)))

	value := stk.head.value

	if stk.len == 0 {
		// value := stk.head.value
		stk.head = nil
		return value, nil
	}

	// value := stk.head.value
	stk.head = stk.head.prev
	return value, nil
}

func (stk *Stack) Peek() (int, error) {
	if stk.head == nil {
		return -1, errors.New("Stack is empty")
	}
	return stk.head.value, nil
}
