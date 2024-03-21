package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	BubbleSort(&arr)

	for i, v := range arr {
		if v != arr[i] {
			t.Error("not sorted")
		}
	}
}
