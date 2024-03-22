package main

import (
	"testing"
)

func expectGetError(t *testing.T, expect int, get int) {
	if expect != get {
		t.Errorf("expected: %v, got: %v", expect, get)
	}
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error %v", err)
	}
}

func TestDoublyLinkedList(t *testing.T) {

	list := NewDoublyLinkedList()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	//
	res1 := list.Get(2)
	expectGetError(t, 9, res1)
	//
	res2, err := list.RemoveAt(1)
	handleErr(t, err)
	expectGetError(t, 7, res2)
	//
	expectGetError(t, 2, list.len)

	list.Append(11)

	//
	res3, err := list.RemoveAt(1)
	handleErr(t, err)
	expectGetError(t, 9, res3)
	// in this case there should be an error for the test to pass so no handleErr
	res4, err := list.RemoveValue(9)
	if err == nil {
		t.Error("there should be an error when value is not found")
	}
	expectGetError(t, -1, res4)
	//
	res5, err := list.RemoveAt(0)
	handleErr(t, err)
	expectGetError(t, 5, res5)
	//
	res6, err := list.RemoveAt(0)
	handleErr(t, err)
	expectGetError(t, 11, res6)
	//
	expectGetError(t, 0, list.len)
	//

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	//
	res7 := list.Get(2)
	expectGetError(t, 5, res7)
	//
	res8 := list.Get(0)
	expectGetError(t, 9, res8)
	//
	res9, err := list.RemoveValue(9)
	handleErr(t, err)
	expectGetError(t, 9, res9)
	//
	expectGetError(t, 2, list.len)
	//
	res10 := list.Get(0)
	expectGetError(t, 7, res10)

	// TODO: insert at
}