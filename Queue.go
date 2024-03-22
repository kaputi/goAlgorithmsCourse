package main

import "errors"

type QNode struct {
	value int
	next  *QNode
}

type Queue struct {
	len  int
	head *QNode
	tail *QNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value int) {
	newNode := &QNode{value: value}

	q.len++

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode

}

func (q *Queue) Deque() (int, error) {
	if q.head == nil {
		return -1, errors.New("Queue is empty")
	}
	q.len--

	value := q.head.value

	q.head = q.head.next

	if q.len == 0 {
		q.tail = nil
	}

	return value, nil
}

func (q *Queue) Peek() (int, error) {
	if q.len == 0 || q.head == nil {
		return -1, errors.New("Queue is empty")
	}

	return q.head.value, nil
}
