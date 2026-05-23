package core

import (
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

// Testing with a slice that will cause an error in Mean
