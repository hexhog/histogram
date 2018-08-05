package histogram

import (
	"fmt"
	"testing"
)

func compute(b int, data [][]float64) ([]float64, []float64, float64, float64, float64, float64, float64, float64) {
	d := len(data[0])
	h := NewHistogram(b, d)

	for _, val := range data {
		h.Add(val)
	}

	mean := h.Mean()
	variance := h.Variance()
	sd := sqrt(variance)

	return mean, variance, h.Count(), h.CDF(mean), h.CDF(subtract(mean, multiply(2, sd))), h.CDF(subtract(mean, sd)), h.CDF(add(mean, sd)), h.CDF(add(mean, multiply(2, sd)))
}

func TestSampleData(t *testing.T) {
	for d, data := range [][][]float64{dataDimension1} {
		fmt.Println("DIMENSION", d+1)
		_mean, _variance, _count, _, _, _, _, _ := compute(1, data)

		for _, b := range []int{32, 64, 128} {
			fmt.Println("BINS", b)

			mean, variance, count, cdf0, cdf1, cdf2, cdf3, cdf4 := compute(b, data)
			sd := sqrt(variance)

			fmt.Println("COUNT", count)
			if !approx(count, _count) {
				t.Errorf("Count across different bins of size %d incorrect %v %v", b, count, _count)
			}

			fmt.Println("MEAN", mean)
			for i := range mean {
				if !approx(mean[i], _mean[i]) {
					t.Errorf("Mean across different bins of size %d incorrect %v %v", b, mean, _mean)
				}
			}
			fmt.Println("VARIANCE", variance)
			for i := range variance {
				if !approx(variance[i], _variance[i]) {
					t.Errorf("Variance across different bins of size %d incorrect %v %v", b, variance, _variance)
				}
			}

			// Lower the bin count, lower the accuracy.
			// Accuracy of 0.01 for a min bin value of at least 32.
			// For higher accuracy may need to increase bin count at the cost of increased time for merging bins
			fmt.Println("CDF MEAN", cdf0)
			if !approx2(cdf0, calculate(data, mean, count)) {
				t.Errorf("CDF of size %d dimension %d incorrect %v", b, d+1, cdf0)
			}

			fmt.Println("CDF MEAN - 2SD", cdf1)
			if !approx2(cdf1, calculate(data, subtract(mean, multiply(2, sd)), count)) {
				t.Errorf("CDF of size %d dimension %d incorrect %v", b, d+1, cdf1)
			}

			fmt.Println("CDF MEAN - SD", cdf2)
			if !approx2(cdf2, calculate(data, subtract(mean, sd), count)) {
				t.Errorf("CDF of size %d dimension %d incorrect %v", b, d+1, cdf2)
			}

			fmt.Println("CDF MEAN + SD", cdf3)
			if !approx2(cdf3, calculate(data, add(mean, sd), count)) {
				t.Errorf("CDF of size %d dimension %d incorrect %v", b, d+1, cdf3)
			}

			fmt.Println("CDF MEAN + 2SD", cdf4)
			if !approx2(cdf4, calculate(data, add(mean, multiply(2, sd)), count)) {
				t.Errorf("CDF of size %d dimension %d incorrect %v", b, d+1, cdf4)
			}
		}
	}
}

func less(x, y []float64) bool {
	for i := range x {
		if x[i] > y[i] {
			return false
		}
	}
	return true
}

func calculate(data [][]float64, compare []float64, count float64) float64 {
	sum := 0.0
	for i := range data {
		if less(data[i], compare) {
			sum += 1
		}
	}
	return sum / count
}
