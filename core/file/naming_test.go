package file

import (
	"testing"

	"github.com/1chooo/leetcode-crawler/config"
)

func TestFormatDirectorySlug(t *testing.T) {
	tests := []struct {
		slug   string
		conv   string
		expect string
	}{
		{"two-sum", config.DefaultConfig.NamingConvention.KebabCase, "two-sum"},
		{"two-sum", config.DefaultConfig.NamingConvention.SnakeCase, "two_sum"},
		{"two-sum", config.DefaultConfig.NamingConvention.LowerCamelCase, "twoSum"},
		{"two-sum", config.DefaultConfig.NamingConvention.UpperCamelCase, "TwoSum"},
		{"", config.DefaultConfig.NamingConvention.SnakeCase, ""},
		{"single", config.DefaultConfig.NamingConvention.LowerCamelCase, "single"},
	}
	for _, tt := range tests {
		if got := FormatDirectorySlug(tt.slug, tt.conv); got != tt.expect {
			t.Errorf("FormatDirectorySlug(%q, %q) = %q, want %q", tt.slug, tt.conv, got, tt.expect)
		}
	}
}

func TestFormatDirectorySlug_unknownConventionFallsBackToKebab(t *testing.T) {
	got := FormatDirectorySlug("a-b", "weird")
	if got != "a-b" {
		t.Fatalf("got %q", got)
	}
}
