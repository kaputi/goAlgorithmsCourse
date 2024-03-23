package main

import (
	"errors"
)

type MinHeap struct {
	data []int
	len  int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (mh *MinHeap) Insert(value int) {
	mh.data = append(mh.data, value)
	mh.heapifyUp(mh.len)
	mh.len++
}

func (mh *MinHeap) Delete() (int, error) {
  // TODO: it only actualy deletes when it gets down to len 0, but is setting length so when you use Get it returns correctly

	if mh.len == 0 {
		return -1, errors.New("Heap is empty")
	}

	out := mh.data[0]
	mh.len--

	if mh.len == 0 {
		mh.data = []int{}
		return out, nil
	}

	mh.data[0] = mh.data[mh.len]
	mh.heapifyDown(0)

	return out, nil
}

func(mh *MinHeap) Get() []int{
  return mh.data[:mh.len]
}

func (mh *MinHeap) heapifyDown(idx int) {
	if idx >= mh.len {
		return
	}

	leftIdx := mh.getLeftChildIdx(idx)
	rightIdx := mh.getRightChildIdx(idx)

	if leftIdx >= mh.len {
		return
	}

	leftValue := mh.data[leftIdx]
	rightValue := mh.data[rightIdx]
	value := mh.data[idx]

	if leftValue > rightValue && value > rightValue {
		mh.data[idx] = rightValue
		mh.data[rightIdx] = value
		mh.heapifyDown(rightIdx)
	} else if rightValue > leftValue && value > leftValue {
		mh.data[idx] = leftValue
		mh.data[leftIdx] = value
		mh.heapifyDown(leftIdx)
	}
}

func (mh *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := mh.getParentIdx(idx)
	parentValue := mh.data[parentIdx]
	value := mh.data[idx]

	if parentValue > value {
		mh.data[idx] = parentValue
		mh.data[parentIdx] = value

		mh.heapifyUp(parentIdx)
	}

}

func (mh *MinHeap) getParentIdx(idx int) int {
	return (idx - 1) / 2
}

func (mh *MinHeap) getLeftChildIdx(idx int) int {
	return idx*2 + 1
}

func (mh *MinHeap) getRightChildIdx(idx int) int {
	return idx*2 + 2
}
