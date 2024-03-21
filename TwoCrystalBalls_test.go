package main

import (
	"math/rand"
	"testing"
)

func TestTowCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)

	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	result1 := TwoCrystalBalls(data)
	if result1 != idx {
    t.Errorf("expected: %v, got: %v", idx, result1)
	}

	result2 := TwoCrystalBalls(make([]bool, 1000))
	if result2 != -1 {
    t.Errorf("expected: %v, got: %v", -1, result1)
	}
}
