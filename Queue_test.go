package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(5)
	q.Enqueue(7)
	q.Enqueue(9)

	//
	res1, err := q.Deque()
	HandleErr(t, err)
	ExpectGetInt(t, 5, res1)
	//
	ExpectGetInt(t, 2, q.len)

	q.Enqueue(11)

	res2, err := q.Deque()
	HandleErr(t, err)
	ExpectGetInt(t, 7, res2)
	//
	res3, err := q.Deque()
	HandleErr(t, err)
	ExpectGetInt(t, 9, res3)
	//
	res4, err := q.Peek()
	HandleErr(t, err)
	ExpectGetInt(t, 11, res4)
	//
	res5, err := q.Deque()
	HandleErr(t, err)
	ExpectGetInt(t, 11, res5)
	//
	res6, err := q.Deque()
	if err == nil {
		t.Error("there should be an error when value is not found")
	}
	ExpectGetInt(t, -1, res6)
	//
	ExpectGetInt(t, 0, q.len)

	q.Enqueue(69)

	//
	res7, err := q.Peek()
	HandleErr(t, err)
	ExpectGetInt(t, 69, res7)
	//
	ExpectGetInt(t, 1, q.len)
}
