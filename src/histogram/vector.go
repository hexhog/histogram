package histogram

import (
	"fmt"
	"math"
)

type bin struct {
	vec      vector
	variance vector
	count    float64
	min      vector
	max      vector
}

// http://www.science.canterbury.ac.nz/nzns/issues/vol7-1979/duncan_b.pdf
func (b *bin) Merge(o bin) bin {
	dimension := b.vec.Dimension()

	count := b.count + o.count

	mean := make([]float64, dimension)
	variance := make([]float64, dimension)
	min := make([]float64, dimension)
	max := make([]float64, dimension)

	for i := 0; i < dimension; i++ {
		mean[i] = (b.count*b.vec.Value(i) + o.count*o.vec.Value(i)) / float64(count)

		variance[i] =
			((b.count*(b.variance.Value(i)+b.vec.Value(i)*b.vec.Value(i)) +
				o.count*(o.variance.Value(i)+o.vec.Value(i)*o.vec.Value(i))) / float64(count)) - mean[i]*mean[i]

		if b.min.Value(i) <= o.min.Value(i) {
			min[i] = b.min.Value(i)
		} else {
			min[i] = o.min.Value(i)
		}

		if b.max.Value(i) >= o.max.Value(i) {
			max[i] = b.max.Value(i)
		} else {
			max[i] = o.max.Value(i)
		}

	}

	return bin{
		vec:      NewVector(mean),
		variance: NewVector(variance),
		count:    count,
		min:      NewVector(min),
		max:      NewVector(max),
	}
}

type vector struct {
	values []float64
}

func NewVector(v []float64) vector {
	return vector{values: v}
}

func (v *vector) Dimension() int {
	return len(v.values)
}

func (v *vector) Value(k int) float64 {
	return v.values[k]
}

func (v *vector) Values() []float64 {
	return v.values
}

func (v *vector) Distance(o vector) float64 {
	var sum float64 = 0

	for i := range v.values {
		sum = sum + square(v.Value(i)-o.Value(i))
	}

	return math.Sqrt(sum)
}

func (v *vector) LessThanOrEqualTo(o vector) bool {
	vValues := v.Values()
	oValues := o.Values()

	for i := range vValues {
		if vValues[i] > oValues[i] {
			return false
		}
	}

	return true
}

func (v *vector) String() string {
	return fmt.Sprintf("%v", v.values)
}

func (v *vector) Equals(o vector) bool {
	if v.Dimension() != o.Dimension() {
		return false
	}
	for i := range v.values {
		if v.Value(i) != o.Value(i) {
			return false
		}
	}
	return true
}
