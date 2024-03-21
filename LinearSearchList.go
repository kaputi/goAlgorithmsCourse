package main

func LinearSearchList(haystack []int, needle int) int {
	idx := -1

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			idx = i
			break
		}
	}

	return idx
}
