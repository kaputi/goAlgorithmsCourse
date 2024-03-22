package main

import (
	"testing"
)

func expectGetErrorStack(t *testing.T, expect int, get int) {
	if expect != get {
		t.Errorf("expected: %v, got: %v", expect, get)
	}
}

func handleErrStack(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error %v", err)
	}
}

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(5)
	stack.Push(7)
	stack.Push(9)

	//
	res1, err := stack.Pop()
	handleErrStack(t, err)
	expectGetErrorStack(t, 9, res1)
	//
	expectGetErrorStack(t, 2, stack.len)

	stack.Push(11)
	//
	res2, err := stack.Pop()
	handleErrStack(t, err)
	expectGetErrorStack(t, 11, res2)
	//
	res3, err := stack.Pop()
	handleErrStack(t, err)
	expectGetErrorStack(t, 7, res3)
	//
	res4, err := stack.Peek()
	handleErrStack(t, err)
	expectGetErrorStack(t, 5, res4)
	//
	res5, err := stack.Pop()
	handleErrStack(t, err)
	expectGetErrorStack(t, 5, res5)
	//
	res6, err := stack.Pop()
	if err == nil {
		t.Error("there should be an error when value is not found")
	}
	expectGetErrorStack(t, -1, res6)

	stack.Push(69)

	//
	res7, err := stack.Peek()
	handleErrStack(t, err)
	expectGetErrorStack(t, 69, res7)
	//
	expectGetErrorStack(t, 1, stack.len)
}
