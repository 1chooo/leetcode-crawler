package file

import (
	"strings"

	"github.com/1chooo/leetcode-crawler/config"
)

// FormatDirectorySlug formats a LeetCode kebab-case title slug for use in a problem directory name.
func FormatDirectorySlug(titleSlug, convention string) string {
	titleSlug = strings.TrimSpace(titleSlug)
	if titleSlug == "" {
		return titleSlug
	}

	parts := strings.Split(titleSlug, "-")
	var words []string
	for _, p := range parts {
		if p != "" {
			words = append(words, p)
		}
	}
	if len(words) == 0 {
		return titleSlug
	}

	switch convention {
	case config.DefaultConfig.NamingConvention.KebabCase:
		return strings.Join(words, "-")
	case config.DefaultConfig.NamingConvention.SnakeCase:
		return strings.Join(words, "_")
	case config.DefaultConfig.NamingConvention.LowerCamelCase:
		out := strings.ToLower(words[0])
		for i := 1; i < len(words); i++ {
			p := words[i]
			if len(p) > 0 {
				out += strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
			}
		}
		return out
	case config.DefaultConfig.NamingConvention.UpperCamelCase:
		var out string
		for _, p := range words {
			if len(p) > 0 {
				out += strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
			}
		}
		return out
	default:
		return strings.Join(words, "-")
	}
}
