package main

import (
	"testing"
)

func TestPreOrderBTSearch(t *testing.T) {
	answerForTree := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	answerForTree2 := []int{20, 10, 5, 7, 15, 50, 30, 29, 21, 45, 49}

	resultForTree := PreOrderBTSearch(Tree)
	resultForTree2 := PreOrderBTSearch(Tree2)

	for i, ans1 := range answerForTree {

		res1 := resultForTree[i]
		ans2 := answerForTree2[i]
		res2 := resultForTree2[i]

		if ans1 != res1 || ans2 != res2 {
			t.Errorf("idx: %v // answer1: %v, result1: %v", i, ans1, res1)
		}
	}
}
