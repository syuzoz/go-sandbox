package main

import (
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{3, "fizz"},
		{5, "buzz"},
		{15, "fizzbuzz"},
	}

	for _, tt := range tests {
		t.Run("fizzbuzz", func(t *testing.T) {
			actual := fizzbuzz(tt.input)
			if actual != tt.expected {
				t.Errorf("input: %d, got: %s, want: %s", tt.input, actual, tt.expected)
			}
		})
	}
} 
