package core

import (
	"math"
	"testing"
)

func TestMean_HappyPath(t *testing.T) {
	got, err := Mean([]float64{1, 2, 3, 4})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 2.5 {
		t.Fatalf("want 2.5, got %v", got)
	}
}

// Corrected expected value

func TestVariance_EmptyInput(t *testing.T) {
	result, err := Variance([]float64{})
	if err != ErrEmpty {
		t.Errorf("expected error %v, got %v", ErrEmpty, err)
	}
	if result != 0 {
		t.Errorf("expected variance 0, got %v", result)
	}
}

func TestVariance_SingleValue(t *testing.T) {
	result, err := Variance([]float64{5.0})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 0 {
		t.Errorf("expected variance 0, got %v", result)
	}
}

func TestVariance_MultipleValues(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
	}{
		{"Normal case", []float64{1, 2, 3, 4, 5}, 2.5},
		{"Negative values", []float64{-1, -2, -3, -4, -5}, 2.5},
		{"Mixed values", []float64{-1, 0, 1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Variance(tt.values)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestVariance_ErrorFromMean(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
	}{
		{"Nil slice", nil},
		{"Empty slice", []float64{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Variance(tt.values)
			if err == nil {
				t.Fatalf("expected an error, got none")
			}
		})
	}
}

func TestVariance_BoundaryValues(t *testing.T) {
	tests := []struct {
		name     string
		values   []float64
		expected float64
		err      error
	}{
		{
			name:     "ZeroValues",
			values:   []float64{0, 0, 0},
			expected: 0,
			err:      nil,
		},
		{
			name:     "NegativeValues",
			values:   []float64{-1, -2, -3},
			expected: 1,
			err:      nil,
		},
		{
			name:     "MixedValues",
			values:   []float64{-1, 0, 1},
			expected: 1,
			err:      nil,
		},

		{
			name:     "MaxAndMinValues",
			values:   []float64{math.MaxFloat64, -math.MaxFloat64},
			expected: math.Inf(1), // Variance should be infinite due to the large difference
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Variance(tt.values)
			if err != nil && err != tt.err {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
