package main

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{9, 3, 7, 4, 69, 420, 42}

	QuickSort(&list)

	// wrong
	// answer := []int{9, 3, 7, 4, 69, 420, 42}
	answer := []int{3, 4, 7, 9, 42, 69, 420}

	for i, v := range list {
		if v != answer[i] {
			t.Error("not sorted")
		}
	}
}
