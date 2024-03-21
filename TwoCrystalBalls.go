package main

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {
	jump := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jump

	for i < len(breaks) {
		if breaks[i] {
			break
		}
		i += jump
	}

	i -= jump

	for j := 0; j < jump && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1
}
