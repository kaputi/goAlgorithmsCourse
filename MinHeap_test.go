package main

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	//
	mhState := []int{1, 4, 3, 7, 5, 69, 8, 420}
	for i, v := range mhState {
		heapVal := heap.data[i]
		ExpectGet(t, v, heapVal)
	}
	//
	ExpectGet(t, 8, heap.len)
	//
	res1, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 1, res1)
	res2, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 3, res2)
	//
	res3, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 4, res3)
	//
	res4, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 5, res4)
	//
	mhState = []int{7, 69, 8, 420}
	newState := heap.Get()
	for i, v := range mhState {
		ExpectGet(t, v, heap.data[i])
		ExpectGet(t, v, newState[i])
	}
	//
	ExpectGet(t, 4, heap.len)
	//
	res5, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 7, res5)
	//
	res6, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 8, res6)
	//
	res7, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 69, res7)
	//
	res8, err := heap.Delete()
	HandleErr(t, err)
	ExpectGet(t, 420, res8)
	//
	ExpectGet(t, 0, heap.len)
}
