package main

func BinarySearchList(haystack []int, needle int) int {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		mid := lo + (hi-lo)/2
		if haystack[mid] == needle {
			return mid
		}
		if haystack[mid] > needle {
			hi = mid
		}
		if haystack[mid] < needle {
			lo = mid + 1
		}
	}

	return -1
}
