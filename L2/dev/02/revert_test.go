package revert

import (
	"testing"

)

func TestRLERevert(t *testing.T) {
	tests := []struct {
		name string
		input    string
		expected string
	}{
		{"first", "a4bc2d5e", "aaaabccddddde"},
		{"second", "abcd", "abcd"},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RLERevert(tt.input)
			if got != tt.expected {
				t.Errorf("Revert(%q) = %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}
