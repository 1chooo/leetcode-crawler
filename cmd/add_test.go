package cmd

import (
	"testing"
)

func TestAdd_ValidInputs(t *testing.T) {
	tests := []struct {
		first    string
		second   string
		expected string
	}{
		{"1", "2", "3.000000"},
		{"2.5", "3.5", "6.000000"},
		{"-1", "1", "0.000000"},
		{"0", "0", "0.000000"},
		{"100", "200", "300.000000"},
	}

	for _, tt := range tests {
		result := Add(tt.first, tt.second)
		if result != tt.expected {
			t.Errorf("Add(%q, %q) = %q; want %q", tt.first, tt.second, result, tt.expected)
		}
	}
}

func TestAdd_InvalidFirstInput(t *testing.T) {
	result := Add("abc", "2")
	if result != "" {
		t.Errorf("Add(%q, %q) = %q; want empty string", "abc", "2", result)
	}
}

func TestAdd_InvalidSecondInput(t *testing.T) {
	result := Add("2", "xyz")
	if result != "" {
		t.Errorf("Add(%q, %q) = %q; want empty string", "2", "xyz", result)
	}
}

func TestAdd_BothInputsInvalid(t *testing.T) {
	result := Add("foo", "bar")
	if result != "" {
		t.Errorf("Add(%q, %q) = %q; want empty string", "foo", "bar", result)
	}
}
