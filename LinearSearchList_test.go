package main

import (
	"testing"
)

func testLinear(t *testing.T, haystack []int, needle int, expect int) {
	index := LinearSearchList(haystack, needle)
	if index != expect {
		t.Errorf("expected: %v, got: %v", index, expect)
	}
}

func TestLinearSearchList(t *testing.T) {
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	testLinear(t, haystack, 69, 3)
	testLinear(t, haystack, 1336, -1)
	testLinear(t, haystack, 69420, 10)
	testLinear(t, haystack, 69421, -1)
	testLinear(t, haystack, 1, 0)
	testLinear(t, haystack, 0, -1)

}
