package cmd

import (
	"testing"
)

func TestSubtract_ValidInputs(t *testing.T) {
	tests := []struct {
		first    string
		second   string
		expected string
	}{
		{"1", "2", "-1.000000"},
		{"2.5", "3.5", "-1.000000"},
		{"-1", "1", "-2.000000"},
		{"0", "0", "0.000000"},
		{"200", "100", "100.000000"},
	}

	for _, tt := range tests {
		result := Subtract(tt.first, tt.second)
		if result != tt.expected {
			t.Errorf("Subtract(%q, %q) = %q; want %q", tt.first, tt.second, result, tt.expected)
		}
	}
}

func TestSubtract_InvalidFirstInput(t *testing.T) {
	result := Subtract("abc", "2")
	if result != "" {
		t.Errorf("Subtract(%q, %q) = %q; want empty string", "abc", "2", result)
	}
}

func TestSubtract_InvalidSecondInput(t *testing.T) {
	result := Subtract("2", "xyz")
	if result != "" {
		t.Errorf("Subtract(%q, %q) = %q; want empty string", "2", "xyz", result)
	}
}

func TestSubtract_BothInputsInvalid(t *testing.T) {
	result := Subtract("foo", "bar")
	if result != "" {
		t.Errorf("Subtract(%q, %q) = %q; want empty string", "foo", "bar", result)
	}
}
