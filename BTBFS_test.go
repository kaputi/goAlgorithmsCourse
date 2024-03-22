package main

import (
	"testing"
)

func TestBFS(t *testing.T) {

	res1 := BFS(Tree, 45)
	ExpectGet(t, true, res1)
	//
	res2 := BFS(Tree, 7)
	ExpectGet(t, true, res2)
	//
	res3 := BFS(Tree, 69)
	ExpectGet(t, false, res3)

}
