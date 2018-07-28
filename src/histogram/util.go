package histogram

import (
	"math"
)

func approx(x, y float64) bool {
	return math.Abs(x-y) < 0.001
}

func square(x float64) float64 {
	return x * x
}

func sort(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}
