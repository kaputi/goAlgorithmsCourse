package main

import (
	"testing"
)

func test(t *testing.T, haystack []int, needle int, expect int) {
	index := BinarySearchList(haystack, needle)
	if index != expect {
		t.Errorf("expected: %v, got: %v", index, expect)
	}
}

func TestBinarySearchList(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	test(t, haystack, 69, 3)
	test(t, haystack, 1336, -1)
	test(t, haystack, 69420, 10)
	test(t, haystack, 69421, -1)
	test(t, haystack, 1, 0)
	test(t, haystack, 0, -1)
}
