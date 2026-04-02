package config

import "testing"

func TestNormalizeLangSlug(t *testing.T) {
	if got := NormalizeLangSlug("go"); got != "golang" {
		t.Errorf("NormalizeLangSlug(go) = %q, want golang", got)
	}
	if got := NormalizeLangSlug("  PYTHON3 "); got != "python3" {
		t.Errorf("NormalizeLangSlug = %q, want python3", got)
	}
}

func TestNormalizeNamingConvention(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"kebab-case", DefaultConfig.NamingConvention.KebabCase},
		{"camelCase", DefaultConfig.NamingConvention.LowerCamelCase},
		{"pascalCase", DefaultConfig.NamingConvention.UpperCamelCase},
		{"snake_case", DefaultConfig.NamingConvention.SnakeCase},
		{"lowerCamelCase", DefaultConfig.NamingConvention.LowerCamelCase},
		{"unknown", DefaultConfig.NamingConvention.KebabCase},
		{"", DefaultConfig.NamingConvention.KebabCase},
	}
	for _, tt := range tests {
		if got := NormalizeNamingConvention(tt.in); got != tt.want {
			t.Errorf("NormalizeNamingConvention(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
