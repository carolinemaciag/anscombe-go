package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

type Dataset struct {
	x []float64
	y []float64
}

//establish linear regression func
func linearRegression(x, y []float64) (slope, intercept float64, err error) {
	if len(x) != len(y) {
		return 0, 0, fmt.Errorf("x and y must be same length")
	}

	meanX, _ := stats.Mean(x)
	meanY, _ := stats.Mean(y)

	var numerator, denominator float64
	for i := range x {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += (x[i] - meanX) * (x[i] - meanX)
	}

	if denominator == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero in slope calculation")
	}

	slope = numerator / denominator
	intercept = meanY - slope*meanX
	return slope, intercept, nil
}

func main() {
	start := time.Now()

	datasets := map[string]Dataset{
		"I": {
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		"II": {
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		},
		"III": {
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		"IV": {
			x: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
			y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
		},
	}

	for name, data := range datasets {
		slope, intercept, err := linearRegression(data.x, data.y)
		if err != nil {
			fmt.Printf("Dataset %s: error %v\n", name, err)
			continue
		}
		fmt.Printf("Dataset %s: Intercept = %.4f, Slope = %.4f\n", name, intercept, slope)
	}

	duration := time.Since(start)
	fmt.Printf("\nGo Execution Time: %v\n", duration)
}
