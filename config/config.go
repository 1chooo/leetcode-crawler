package config

type Config struct {
	CN             SiteConfig
	EN             SiteConfig
	Level          LevelMap
	Language       LanguageMap
	Naming         NamingMap
	QuestionDataQL func(titleSlug string) GraphQLRequest
}

var DefaultConfig = Config{
	CN: SiteConfig{
		Domain: "https://leetcode-cn.com",
	},
	EN: SiteConfig{
		Domain: "https://leetcode.com",
	},
	Level: LevelMap{
		One:   "easy",
		Two:   "medium",
		Three: "hard",
	},
	Language: LanguageMap{
		Java:       "Java",
		JavaScript: "JavaScript",
		Python3:    "Python3",
		CPP:        "C++",
		C:          "C",
		Golang:     "Go",
		Rust:       "Rust",
		TypeScript: "TypeScript",
	},
	Naming: NamingMap{
		SnakeCase:      "snake_case",
		CamelCase:      "CamelCase",
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
