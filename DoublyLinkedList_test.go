package main

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {

	list := NewDoublyLinkedList()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	//
	res1 := list.Get(2)
	ExpectGetInt(t, 9, res1)
	//
	res2, err := list.RemoveAt(1)
	HandleErr(t, err)
	ExpectGetInt(t, 7, res2)
	//
	ExpectGetInt(t, 2, list.len)

	list.Append(11)

	//
	res3, err := list.RemoveAt(1)
	HandleErr(t, err)
	ExpectGetInt(t, 9, res3)
	// in this case there should be an error for the test to pass so no handleErr
	res4, err := list.RemoveValue(9)
	if err == nil {
		t.Error("there should be an error when value is not found")
	}
	ExpectGetInt(t, -1, res4)
	//
	res5, err := list.RemoveAt(0)
	HandleErr(t, err)
	ExpectGetInt(t, 5, res5)
	//
	res6, err := list.RemoveAt(0)
	HandleErr(t, err)
	ExpectGetInt(t, 11, res6)
	//
	ExpectGetInt(t, 0, list.len)
	//

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	//
	res7 := list.Get(2)
	ExpectGetInt(t, 5, res7)
	//
	res8 := list.Get(0)
	ExpectGetInt(t, 9, res8)
	//
	res9, err := list.RemoveValue(9)
	HandleErr(t, err)
	ExpectGetInt(t, 9, res9)
	//
	ExpectGetInt(t, 2, list.len)
	//
	res10 := list.Get(0)
	ExpectGetInt(t, 7, res10)

	// TODO: insert at
}
