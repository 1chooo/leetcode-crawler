package config

type SiteConfig struct {
	Domain string
}

type LevelMap struct {
	Easy   int
	Medium int
	Hard   int
}

type LanguageMap struct {
	Java       string
	JavaScript string
	Python3    string
	CPP        string
	C          string
	Golang     string
	Rust       string
	TypeScript string
}

type NamingMap struct {
	SnakeCase      string
	CamelCase      string
	LowerCamelCase string
	UpperCamelCase string
	KebabCase      string
}

// CodeSnippet represents a code snippet for a specific language
type CodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}

// Question represents the question data structure
type Question struct {
	TranslatedTitle    string        `json:"translatedTitle"`
	TranslatedContent  string        `json:"translatedContent"`
	Content            string        `json:"content"`
	SimilarQuestions   interface{}   `json:"similarQuestions"`
	Stats              interface{}   `json:"stats"`
	Hints              interface{}   `json:"hints"`
	Title              string        `json:"title"`
	TitleSlug          string        `json:"titleSlug"`
	QuestionFrontendID string        `json:"questionFrontendId"`
	CodeSnippets       []CodeSnippet `json:"codeSnippets"`
}

type GraphQLResponse struct {
	Data struct {
		Question Question `json:"question"`
	} `json:"data"`
}

// GraphQLRequest represents the GraphQL request payload
type GraphQLRequest struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

type Pair struct {
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
	Stat struct {
		FrontendQuestionID int    `json:"frontend_question_id"`
		QuestionTitleSlug  string `json:"question__title_slug"`
	} `json:"stat"`
}
