package main

func walkPreOrder(curr *BinaryNode, path *[]int) {
	if curr == nil {
		return
	}

	(*path) = append((*path), curr.value)

	walkPreOrder(curr.left, path)
	walkPreOrder(curr.right, path)

}

func PreOrderBTSearch(head *BinaryNode) []int {
	path := []int{}

	walkPreOrder(head, &path)

	return path
}
