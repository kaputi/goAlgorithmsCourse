package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(5)
	stack.Push(7)
	stack.Push(9)

	//
	res1, err := stack.Pop()
	HandleErr(t, err)
	ExpectGet(t, 9, res1)
	//
	ExpectGet(t, 2, stack.len)

	stack.Push(11)
	//
	res2, err := stack.Pop()
	HandleErr(t, err)
	ExpectGet(t, 11, res2)
	//
	res3, err := stack.Pop()
	HandleErr(t, err)
	ExpectGet(t, 7, res3)
	//
	res4, err := stack.Peek()
	HandleErr(t, err)
	ExpectGet(t, 5, res4)
	//
	res5, err := stack.Pop()
	HandleErr(t, err)
	ExpectGet(t, 5, res5)
	//
	res6, err := stack.Pop()
	if err == nil {
		t.Error("there should be an error when value is not found")
	}
	ExpectGet(t, -1, res6)

	stack.Push(69)

	//
	res7, err := stack.Peek()
	HandleErr(t, err)
	ExpectGet(t, 69, res7)
	//
	ExpectGet(t, 1, stack.len)
}
