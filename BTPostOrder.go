package main

func walkPostOrder(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	walkPostOrder(curr.left, path)
	walkPostOrder(curr.right, path)

	(*path) = append((*path), curr.value)
}

func PostOrderBTSearch(head *BinaryNode) []int {
	path := []int{}

	walkPostOrder(head, &path)

	return path
}
