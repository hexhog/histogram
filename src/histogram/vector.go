package histogram

import (
	"fmt"
	"math"
)

type bin struct {
	vec   vector
	count float64
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

func (v *vector) Mean(o vector) vector {
	m := make([]float64, v.Dimension())

	for i := range v.values {
		m[i] = (v.Value(i) + o.Value(i)) / 2
	}

	return vector{values: m}
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
