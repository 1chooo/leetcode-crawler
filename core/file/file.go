package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/1chooo/leetcode-crawler/config"
)

// QuestionConfig represents the question configuration
type QuestionConfig struct {
	QuestionFrontendID string   `json:"questionFrontendId"`
	Title              string   `json:"title"`
	Domain             string   `json:"domain"`
	TitleSlug          string   `json:"titleSlug"`
	Content            string   `json:"content"`
	Hints              []string `json:"hints"`
}

// Information represents the information to be written to JSON
type Information struct {
	Difficulty       string            `json:"difficulty"`
	SimilarQuestions []SimilarQuestion `json:"similarQuestions"`
}

// SimilarQuestion represents a similar question
type SimilarQuestion struct {
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
}

// WriteDirectory creates a directory if it doesn't exist
func WriteDirectory(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return os.MkdirAll(dirName, 0755)
	}
	return nil
}

// WriteQuestion writes the question README.md file
func WriteQuestion(dirName string, questionConfig QuestionConfig) error {
	filePath := filepath.Join(dirName, "README.md")
	fmt.Printf("created: %s\n", filePath)

	var fileContent strings.Builder
	fileContent.WriteString(fmt.Sprintf("## [%s. %s](%s/problems/%s/)\n",
		questionConfig.QuestionFrontendID,
		questionConfig.Title,
		questionConfig.Domain,
		questionConfig.TitleSlug))
	fileContent.WriteString(questionConfig.Content)

	if len(questionConfig.Hints) > 0 {
		fileContent.WriteString("\n\n## Hints\n")
		for i, hint := range questionConfig.Hints {
			fileContent.WriteString(fmt.Sprintf("%d. %s\n", i+1, hint))
		}
	}

	return os.WriteFile(filePath, []byte(fileContent.String()), 0644)
}

// WriteSolution writes the solution file with appropriate extension
func WriteSolution(dirName string, langSlug string, code string, config config.Config) error {
	// Map language slugs to file extensions
	langExtMap := map[string]string{
		"java":       ".java",
		"javascript": ".js",
		"python3":    ".py",
		"cpp":        ".cpp",
		"c":          ".c",
		"golang":     ".go",
		"rust":       ".rs",
		"typescript": ".ts",
	}

	ext, exists := langExtMap[strings.ToLower(langSlug)]
	if !exists {
		ext = ".txt" // fallback extension
	}

	filePath := filepath.Join(dirName, "Solution"+ext)

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists, don't overwrite
		return nil
	}

	return os.WriteFile(filePath, []byte(code), 0644)
}

// WriteInformation writes the information.json file
func WriteInformation(dirName string, question map[string]interface{}, difficulty string) error {
	filePath := filepath.Join(dirName, "information.json")

	// Parse similar questions
	var similarQuestions []SimilarQuestion
	if sq, ok := question["similarQuestions"].(string); ok && sq != "" {
		if err := json.Unmarshal([]byte(sq), &similarQuestions); err != nil {
			// If parsing fails, initialize empty slice
			similarQuestions = []SimilarQuestion{}
		}
	}

	data := Information{
		Difficulty:       difficulty,
		SimilarQuestions: similarQuestions,
	}

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists, read existing data and update
		existingData, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		var jsonData Information
		if err := json.Unmarshal(existingData, &jsonData); err != nil {
			// If parsing fails, use new data
			jsonData = data
		} else {
			// Update existing data (you can modify this logic as needed)
			jsonData.Difficulty = data.Difficulty
			jsonData.SimilarQuestions = data.SimilarQuestions
		}
		data = jsonData
	}

	// Write JSON with proper indentation
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonBytes, 0644)
}

// Helper function to get language extension from config
func GetLanguageExtension(langSlug string, config config.Config) string {
	// You can extend this based on your specific needs
	switch strings.ToLower(langSlug) {
	case "java":
		return ".java"
	case "javascript":
		return ".js"
	case "python3":
		return ".py"
	case "cpp":
		return ".cpp"
	case "c":
		return ".c"
	case "golang", "go":
		return ".go"
	case "rust":
		return ".rs"
	case "typescript":
		return ".ts"
	default:
		return ".txt"
	}
}
