package histogram

import (
	"math"
)

func approx(x, y float64) bool {
	return math.Abs(x-y) < 0.0001
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

func linspace(start, stop float64, num int) []float64 {
	step := 0.
	if num == 1 {
		return []float64{start}
	}
	step = (stop - start) / float64(num-1)

	r := make([]float64, num, num)
	for i := 0; i < num; i++ {
		r[i] = start + float64(i)*step
	}
	return r
}
