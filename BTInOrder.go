package main

func walkInOrder(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	walkInOrder(curr.left, path)
	(*path) = append((*path), curr.value)
	walkInOrder(curr.right, path)
}

func InOrderBTSearch(head *BinaryNode) []int {
	path := []int{}

	walkInOrder(head, &path)

	return path
}
