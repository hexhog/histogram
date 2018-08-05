package histogram

import (
	"fmt"
	"testing"
)

func compute(b int, data [][]float64) ([]float64, []float64, float64, float64) {
	d := len(data[0])
	h := NewHistogram(b, d)

	for _, val := range data {
		h.Add(val)
	}

	mean := h.Mean()

	return mean, h.Variance(), h.Count(), h.CDF(mean)
}

func TestSampleData(t *testing.T) {
	var (
		_mean     []float64
		_variance []float64
		_count    float64
	)

	for d, data := range [][][]float64{dataDimension1, dataDimension2, dataDimension3, dataDimension4, dataDimension5} {
		fmt.Println("DIMENSION", d+1)
		_mean, _variance, _count, _ = compute(1, data)

		for _, b := range []int{128, 256, 512} {
			fmt.Println("BINS", b)

			mean, variance, count, cdf := compute(b, data)
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

			fmt.Println("CDF", cdf)
			// Lower the bin count, lower the accuracy.
			// Accuracy of 0.05 for a min bin value of at least 128.
			// For higher accuracy may need to increase bin count at the cost of increased time for merging bins
			if !approx2(cdf, (1 / pow(d+1))) {
				t.Errorf("CDF of size %d dimension %d incorrect %v", b, d+1, cdf)
			}
		}
	}
}
