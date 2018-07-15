package histogram

import (
	// "fmt"
	"math/rand"
	"testing"
)

func TestHistogram(t *testing.T) {
	b := 4
	d := 3

	h := NewHistogram(b, d)
	var sample = [][]float64{}

	for j := 0; j < 10; j++ {
		var values = []float64{}
		for i := 0; i < d; i++ {
			values = append(values, float64(rand.Intn(100)))
		}
		sample = append(sample, values)
		h.Add(values)
	}
	// fmt.Println(sample)
	count := h.Count()
	if !approx(count, float64(len(sample))) {
		t.Errorf("Count mismatch %v != %v", count, len(sample))
	}

	var sum = make([]float64, d)
	for _, values := range sample {
		for i := 0; i < d; i++ {
			sum[i] = sum[i] + values[i]
		}
	}

	for k, _ := range sum {
		sum[k] = sum[k] / float64(len(sample))
	}
	// fmt.Println(sum)

	mean := h.Mean()
	for k, _ := range sum {
		if !approx(mean[k], sum[k]) {
			t.Errorf("Mean mismatch %v != %v", mean, sum)
		}
	}

	var sumsquare = make([]float64, d)
	for _, values := range sample {
		for i := 0; i < d; i++ {
			sumsquare[i] = sumsquare[i] + (values[i]-sum[i])*(values[i]-sum[i])
		}
	}

	for k, _ := range sumsquare {
		sumsquare[k] = sumsquare[k] / float64(len(sample))
	}
	// fmt.Println(sumsquare)

	variance := h.Variance()
	for k, _ := range sumsquare {
		if !approx(variance[k], sumsquare[k]) {
			t.Errorf("Variance mismatch %v != %v", variance, sumsquare)
		}
	}
}
