package config

type Config struct {
	CN             SiteConfig
	EN             SiteConfig
	Level          LevelMap
	Language       []Language
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
		Easy:   1,
		Medium: 2,
		Hard:   3,
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
