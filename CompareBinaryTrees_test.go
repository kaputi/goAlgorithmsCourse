package main

import (
	"testing"
)

func TestCompareBinaryTrees(t *testing.T) {
	res1 := CompareBinaryTrees(Tree, Tree2)
	ExpectGet(t, false, res1)

	res2 := CompareBinaryTrees(Tree, Tree)
	ExpectGet(t, true, res2)

	res3 := CompareBinaryTrees(Tree2, Tree2)
	ExpectGet(t, true, res3)

}
