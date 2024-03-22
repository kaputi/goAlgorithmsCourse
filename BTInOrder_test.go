package main

import (
	"testing"
)

func TestInOrderBTSearch(t *testing.T) {
	// wrong
	// answerForTree := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 10}
	// answerForTree2 := []int{5, 7, 10, 15, 20, 21, 29, 30, 45, 49, 5}

	answerForTree := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	answerForTree2 := []int{5, 7, 10, 15, 20, 21, 29, 30, 45, 49, 50}

	resultForTree := InOrderBTSearch(Tree)
	resultForTree2 := InOrderBTSearch(Tree2)

	for i, ans1 := range answerForTree {
		ans2 := answerForTree2[i]
		res1 := resultForTree[i]
		res2 := resultForTree2[i]

		if ans1 != res1 || ans2 != res2 {
			t.Errorf("idx: %v // answer1: %v, result1: %v // answer2: %v, result2: %v", i, ans1, res1, ans2, res2)
		}
	}
}
