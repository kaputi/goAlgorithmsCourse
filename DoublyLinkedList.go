package main

import (
	"errors"
	"fmt"
)

// Define node structure
type DllNode struct {
	value int
	prev  *DllNode // Pointer to previous Node
	next  *DllNode // Pointer to next Node
}

// Define the DoublyLinkedList structure
type DoublyLinkedList struct {
	len  int
	head *DllNode
	tail *DllNode
}

// Instatiaate a new DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Add a new node at the end of DoublyLinkedList
func (dll *DoublyLinkedList) Append(value int) {
	newNode := &DllNode{value: value, prev: nil, next: nil}

	dll.len++

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
}

// add a new node at the start of DoublyLinkedList
func (dll *DoublyLinkedList) Prepend(value int) {
	newNode := &DllNode{value: value, prev: nil, next: nil}

	dll.len++

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}

// add new node at idx of DoublyLinkedList
func (dll *DoublyLinkedList) InsertAt(value int, idx int) {
	if idx > dll.len {
		panic("Oh no")
	}

	if idx == dll.len {
		dll.Append(value)
		return
	} else if idx == 0 {
		dll.Prepend(value)
		return
	}

	// fist attach new node
	dll.len++
	curr := dll.getAt(idx)
	newNode := &DllNode{value: value, prev: nil, next: nil}
	newNode.next = curr
	newNode.prev = curr.prev

	// then remove old links
	curr.prev = newNode
	if newNode.prev != nil {
		newNode.prev.next = curr
	}
}

// get value of node at idx
func (dll *DoublyLinkedList) Get(idx int) int {
	return dll.getAt(idx).value
}

// remove node with value and return the value, if not found returns -1, error
func (dll *DoublyLinkedList) RemoveValue(value int) (int, error) {
	curr := dll.head

	// walk the list until value is found
	for i := 0; i < dll.len; i++ {
		if curr.value == value {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return -1, errors.New("Value not found.")
	}

	return dll.removeNode(curr), nil
}

// Remove node at idx
func (dll *DoublyLinkedList) RemoveAt(idx int) (int, error) {
	node := dll.getAt(idx)
	if node == nil {
		return -1, fmt.Errorf("No Node At idx %v", idx)
	}

	return dll.removeNode(node), nil
}

// get node at idx from DoublyLinkedList
func (dll *DoublyLinkedList) getAt(idx int) *DllNode {
	curr := dll.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr
}

// remove node from DoublyLinkedList and return removed node value
func (dll *DoublyLinkedList) removeNode(node *DllNode) int {
	dll.len--
	// if only one node in list
	if dll.len == 0 {
		value := dll.head.value
		dll.head = nil
		dll.tail = nil
		return value
	}

	// next points to prev and prev points to next
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	// update head and tail if needed
	if node == dll.head {
		dll.head = node.next
	}
	if node == dll.tail {
		dll.tail = node.prev
	}

	node.prev = nil
	node.next = nil

	return node.value
}
