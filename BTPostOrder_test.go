package main

import (
	"testing"
)

func TestPostOrderBTSearch(t *testing.T) {
	answerForTree := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	answerForTree2 := []int{7, 5, 15, 10, 21, 29, 49, 45, 30, 50}

	resultForTree := PostOrderBTSearch(Tree)
	resultForTree2 := PostOrderBTSearch(Tree2)

	for i, ans1 := range answerForTree {

		res1 := resultForTree[i]
		ans2 := answerForTree2[i]
		res2 := resultForTree2[i]

		if ans1 != res1 || ans2 != res2 {
			t.Errorf("idx: %v // answer1: %v, result1: %v", i, ans1, res1)
		}
	}
}
