package histogram

import (
	"fmt"
)

type Histogram interface {
	Add(vector []float64)

	Mean() []float64

	Variance() []float64

	CDF(x []float64) float64

	Quantile(q float64) []float64

	String() (str string)

	Count() float64
}

type histogram struct {
	bins      []bin
	maxbins   int
	total     uint64
	dimension int
}

func NewHistogram(n int, d int) Histogram {
	return &histogram{
		bins:      make([]bin, 0),
		maxbins:   n,
		total:     0,
		dimension: d,
	}
}

func (h *histogram) Add(values []float64) {
	m := NewVector(values)
	v := NewVector(make([]float64, len(values)))

	if h.dimension != m.Dimension() {
		return
	}
	h.total++
	for i := range h.bins {
		if h.bins[i].vec.Equals(v) {
			h.bins[i].count++
			return
		}
	}
	h.bins = append(h.bins, bin{count: 1, vec: m, variance: v, min: m, max: m})
	h.trim()
}

func (h *histogram) Mean() []float64 {
	if h.total == 0 {
		return []float64{}
	}

	sum := make([]float64, h.dimension)

	for i := range h.bins {
		for j := range sum {
			sum[j] += h.bins[i].vec.Value(j) * h.bins[i].count
		}
	}

	for k, s := range sum {
		s = s / float64(h.total)
		sum[k] = s
	}
	return sum
}

// http://www.science.canterbury.ac.nz/nzns/issues/vol7-1979/duncan_b.pdf
func (h *histogram) Variance() []float64 {
	if h.total == 0 {
		return []float64{}
	}

	sum := make([]float64, h.dimension)
	mean := h.Mean()

	for i := range h.bins {
		for j := range sum {
			sum[j] += (h.bins[i].count * (h.bins[i].variance.Value(j) + h.bins[i].vec.Value(j)*h.bins[i].vec.Value(j)))
		}
	}

	for k, _ := range sum {
		sum[k] = sum[k] / float64(h.total)
		sum[k] = sum[k] - mean[k]*mean[k]
	}
	return sum
}

func (h *histogram) Quantile(q float64) []float64 {
	count := q * float64(h.total)
	for i := range h.bins {
		count -= float64(h.bins[i].count)

		if count <= 0 {
			return h.bins[i].vec.Values()
		}
	}

	return []float64{}
}

func (h *histogram) CDF(x []float64) float64 {
	xVec := NewVector(x)
	if xVec.Dimension() != h.dimension {
		return -1
	}
	sum := 0.0
	for i := range h.bins {
		count := h.bins[i].count
		for j := 0; j < h.dimension; j++ {
			var (
				factor float64
				x      = xVec.Value(j)
				min    = h.bins[i].min.Value(j)
				max    = h.bins[i].max.Value(j)
			)
			if x < min {
				factor = 1
			} else if x >= max {
				factor = 0
			} else {
				factor = (x - min) / float64(max-min)
			}
			count *= factor
		}
		sum += count
	}

	return sum / float64(h.total)
}

func (h *histogram) String() (str string) {
	str += fmt.Sprintln("Total:", h.total)

	for i := range h.bins {
		var bar string
		// for j := 0; j < int(float64(h.bins[i].count)/float64(h.total)*200); j++ {
		for j := 0; j < int(h.bins[i].count); j++ {
			bar += "."
		}
		// str += fmt.Sprintln(h.bins[i].vec.String(), "\t", bar)
		str += fmt.Sprintln(h.bins[i].vec.String(), h.bins[i].min.String(), h.bins[i].max.String(), "\t", h.bins[i].count)
	}

	return
}

func (h *histogram) Count() float64 {
	return float64(h.total)
}

// ==============================================================================
// trim merges adjacent bins to decrease the bin count to the maximum value
func (h *histogram) trim() {
	for len(h.bins) > h.maxbins {
		// Find closest bins in terms of value
		minDelta := 1e99
		min_i := 0
		min_j := 0
		for i := range h.bins {
			for j := range h.bins { // TODO: iterate from i + 1 to end
				if i == j {
					continue
				}

				if delta := h.bins[i].vec.Distance(h.bins[j].vec); delta < minDelta {
					minDelta = delta
					min_i = i
					min_j = j
				}

			}
		}

		// We need to merge bins min_i-1 and min_j
		mergedbin := h.bins[min_i].Merge(h.bins[min_j])

		// Remove min_i and min_j bins
		min, max := sort(min_i, min_j)

		head := h.bins[0:min]
		mid := h.bins[min+1 : max]
		tail := h.bins[max+1:]

		h.bins = append(head, mid...)
		h.bins = append(h.bins, tail...)

		h.bins = append(h.bins, mergedbin)
	}
}
