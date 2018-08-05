package histogram

import (
	"math"
)

func approx(x, y float64) bool {
	return math.Abs(x-y) < 0.0001
}

func approx2(x, y float64) bool {
	return math.Abs(x-y) < 0.01
}

func sqrt(x []float64) []float64 {
	r := make([]float64, len(x))
	for i := range x {
		r[i] = math.Sqrt(x[i])
	}
	return r
}

func add(x, y []float64) []float64 {
	if len(x) != len(y) {
		return []float64{}
	}
	r := make([]float64, len(x))
	for i := range x {
		r[i] = x[i] + y[i]
	}
	return r
}

func subtract(x, y []float64) []float64 {
	if len(x) != len(y) {
		return []float64{}
	}
	r := make([]float64, len(x))
	for i := range x {
		r[i] = x[i] - y[i]
	}
	return r
}

func multiply(y int, x []float64) []float64 {
	r := make([]float64, len(x))
	for i := range x {
		r[i] = float64(y) * x[i]
	}
	return r
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

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func log(x float64) float64 {
	if x == 0 {
		return 1
	}
	return math.Log(x)
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
