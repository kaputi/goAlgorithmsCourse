package main

func BFS(head *BinaryNode, needle int) bool {

	q := []*BinaryNode{head}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == nil {
			continue
		}

		if curr.value == needle {
			return true
		}

		q = append(q, curr.left)
		q = append(q, curr.right)

	}

	return false
}
