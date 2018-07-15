package histogram

import (
	// "fmt"
	"math/rand"
	"testing"
)

func TestHistogram(t *testing.T) {
	for _, b := range []int{2, 3, 4, 5, 10} {

		for _, d := range []int{2, 3, 4, 5, 10} {

			h := NewHistogram(b, d)
			var sample = [][]float64{}

			for _, n := range []int{1, 5, 10, 50, 100} {

				for j := 0; j < n; j++ {
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
		}
	}
}

func BenchmarkHistogramBinsD1N50(t *testing.B)  { benchmarkHistogramBins(1, 50, t) }
func BenchmarkHistogramBinsD2N50(t *testing.B)  { benchmarkHistogramBins(2, 50, t) }
func BenchmarkHistogramBinsD5N50(t *testing.B)  { benchmarkHistogramBins(5, 50, t) }
func BenchmarkHistogramBinsD10N50(t *testing.B) { benchmarkHistogramBins(10, 50, t) }

func BenchmarkHistogramBinsD1N100(t *testing.B)  { benchmarkHistogramBins(1, 100, t) }
func BenchmarkHistogramBinsD2N100(t *testing.B)  { benchmarkHistogramBins(2, 100, t) }
func BenchmarkHistogramBinsD5N100(t *testing.B)  { benchmarkHistogramBins(5, 100, t) }
func BenchmarkHistogramBinsD10N100(t *testing.B) { benchmarkHistogramBins(10, 100, t) }

func benchmarkHistogramBins(d int, n int, t *testing.B) {
	for b := 1; b < t.N; b++ {
		h := NewHistogram(b, d)
		var sample = [][]float64{}

		for j := 0; j < 100; j++ {
			var values = []float64{}
			for i := 0; i < d; i++ {
				values = append(values, float64(rand.Intn(100)))
			}
			sample = append(sample, values)
			h.Add(values)
		}
	}
}
