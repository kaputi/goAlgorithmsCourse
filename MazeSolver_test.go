package main

import (
	"testing"
)

func TestMazeSolver(t *testing.T) {

	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	answer := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	result := MazeSolver(maze, byte('x'), Point{x: 10, y: 0}, Point{x: 1, y: 5})

	for i, v := range answer {
		if v.x != result[i].x || v.y != result[i].y {
			t.Logf("result[%d] %v != answer[%d] %v\n", i, result[i], i, answer[i])
			t.Error("Wrong")
		}
	}
}
