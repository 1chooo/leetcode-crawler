package config

import "strings"

var DefaultConfig = Config{
	Domain: Domain{
		EN: "https://leetcode.com",
		CN: "https://leetcode-cn.com",
	},
	Level: Level{
		Easy:    1,
		Medium:  2,
		Hard:    3,
		Unknown: 0,
	},
	Language: []Language{
		{Lang: "Java", LangSlug: "java", LangExt: ".java"},
		{Lang: "JavaScript", LangSlug: "javascript", LangExt: ".js"},
		{Lang: "Python3", LangSlug: "python3", LangExt: ".py"},
		{Lang: "C++", LangSlug: "cpp", LangExt: ".cpp"},
		{Lang: "C", LangSlug: "c", LangExt: ".c"},
		{Lang: "Golang", LangSlug: "golang", LangExt: ".go"},
		{Lang: "Rust", LangSlug: "rust", LangExt: ".rs"},
		{Lang: "TypeScript", LangSlug: "typescript", LangExt: ".ts"},
	},
	NamingConvention: NamingConvention{
		SnakeCase:      "snake_case",
		LowerCamelCase: "lowerCamelCase",
		UpperCamelCase: "UpperCamelCase",
		KebabCase:      "kebab-case",
	},
	QuestionDataQL: func(titleSlug string) GraphQLRequest {
		return GraphQLRequest{
			OperationName: "questionData",
			Query: `
			query questionData($titleSlug: String!) {
				question(titleSlug: $titleSlug) {
					translatedTitle
					translatedContent
					content
					similarQuestions
					stats
					hints
					title
					titleSlug
					questionFrontendId
					codeSnippets {
						lang
						langSlug
						code
						__typename
					}
				}
			}`,
			Variables: map[string]interface{}{
				"titleSlug": titleSlug,
			},
		}
	},
}

func GetDefaultConfig() Config {
	return DefaultConfig
}

func GetDomain(isCN bool) string {
	if isCN {
		return DefaultConfig.Domain.CN
	}
	return DefaultConfig.Domain.EN
}

func GetLevel(level string) int {
	switch level {
	case "easy":
		return DefaultConfig.Level.Easy
	case "medium":
		return DefaultConfig.Level.Medium
	case "hard":
		return DefaultConfig.Level.Hard
	default:
		return DefaultConfig.Level.Unknown
	}
}

func GetLanguageExt(langSlug string) (string, bool) {
	for _, lang := range DefaultConfig.Language {
		if lang.LangSlug == langSlug {
			return lang.LangExt, true
		}
	}
	return "", false
}

// NormalizeLangSlug maps common CLI aliases to LeetCode API slugs.
func NormalizeLangSlug(langSlug string) string {
	s := strings.ToLower(strings.TrimSpace(langSlug))
	if s == "go" {
		return "golang"
	}
	return s
}

func GetLanguageBySlug(langSlug string) (Language, bool) {
	for _, lang := range DefaultConfig.Language {
		if lang.LangSlug == langSlug {
			return lang, true
		}
	}
	return Language{}, false
}

func GetLanguageByName(langName string) (Language, bool) {
	for _, lang := range DefaultConfig.Language {
		if lang.Lang == langName {
			return lang, true
		}
	}
	return Language{}, false
}

func GetNamingConvention(convention string) string {
	switch convention {
	case "snake_case":
		return DefaultConfig.NamingConvention.SnakeCase
	case "lowerCamelCase":
		return DefaultConfig.NamingConvention.LowerCamelCase
	case "upperCamelCase":
		return DefaultConfig.NamingConvention.UpperCamelCase
	case "kebab-case":
		return DefaultConfig.NamingConvention.KebabCase
	default:
		return ""
	}
}

// NormalizeNamingConvention maps CLI / README variants to the canonical naming value used in DefaultConfig.
// Unknown values default to kebab-case.
func NormalizeNamingConvention(convention string) string {
	c := strings.TrimSpace(strings.ToLower(strings.ReplaceAll(convention, "_", "")))
	switch c {
	case "snakecase":
		return DefaultConfig.NamingConvention.SnakeCase
	case "lowercamelcase", "camelcase":
		return DefaultConfig.NamingConvention.LowerCamelCase
	case "uppercamelcase", "pascalcase":
		return DefaultConfig.NamingConvention.UpperCamelCase
	case "kebabcase", "kebab-case":
		return DefaultConfig.NamingConvention.KebabCase
	default:
		if v := GetNamingConvention(convention); v != "" {
			return v
		}
		return DefaultConfig.NamingConvention.KebabCase
	}
}
