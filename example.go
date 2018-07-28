package main

import (
	"fmt"

	"histogram"
)

var testData = [][]float64{
	[]float64{8, 1, 3},
	[]float64{13, 23, 12},
	[]float64{17, 45, 44},
	[]float64{16, 23, 11},
	[]float64{15, 87, 71},
}

func main() {
	h := histogram.NewHistogram(4, 3)

	for _, val := range testData {
		h.Add(val)
	}

	fmt.Println("MEAN", h.Mean())
	fmt.Println("VARIANCE", h.Variance())
	fmt.Println("QUANTILE", h.Quantile(0.5))
	fmt.Println("CDF", h.CDF([]float64{100, 100, 100}))
	fmt.Println("STRING", h.String())
	fmt.Println("COUNT", h.Count())
}
