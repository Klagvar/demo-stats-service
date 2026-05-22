package core

import (
	"errors"
	"math"
)

var ErrEmpty = errors.New("core: values are empty")

func Mean(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values)), nil
}

func Min(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}
	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func Max(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, ErrEmpty
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}

func Round(value float64, digits int) float64 {
	shift := math.Pow(10, float64(digits))
	return math.Round(value*shift) / shift
}
