package crawl

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/1chooo/leetcode-crawler/config"
	"github.com/1chooo/leetcode-crawler/core/file"
	"github.com/1chooo/leetcode-crawler/core/helper"
)

// AllProblemsResponse represents the structure of the API response
type AllProblemsResponse struct {
	StatStatusPairs []config.Pair `json:"stat_status_pairs"`
}

/*
Example usage

	func main() {
		ids := []int{1, 2, 3, 4, 5}
		langSlugs := []string{"java", "python3", "cpp", "c", "golang", "rust", "typescript"}
		ProblemCrawler(ids, langSlugs)
	}
*/
func ProblemCrawler(ids []int, langSlugs []string) error {
	domain := "https://leetcode.com"

	// Get all problems
	problemsData, err := helper.GetAllProblems(domain)
	if err != nil {
		return fmt.Errorf("failed to get all problems: %w", err)
	}

	// Convert interface{} to proper struct
	problemsJSON, err := json.Marshal(problemsData)
	if err != nil {
		return fmt.Errorf("failed to marshal problems data: %w", err)
	}

	var problems AllProblemsResponse
	if err := json.Unmarshal(problemsJSON, &problems); err != nil {
		return fmt.Errorf("failed to unmarshal problems data: %w", err)
	}

	// Create a map for quick ID lookup
	idMap := make(map[int]bool)
	for _, id := range ids {
		idMap[id] = true
	}

	// Level map for difficulty conversion
	levelMap := map[int]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}

	// Use a WaitGroup to handle concurrent processing
	var wg sync.WaitGroup

	// Process each problem
	for _, pair := range problems.StatStatusPairs {
		frontendID := pair.Stat.FrontendQuestionID
		titleSlug := pair.Stat.QuestionTitleSlug
		difficultyLevel := pair.Difficulty.Level

		// Check if this problem ID is in our target list
		if !idMap[frontendID] {
			continue
		}

		wg.Add(1)
		go func(fid int, slug string, level int) {
			defer wg.Done()

			if err := processProblem(domain, fid, slug, level, langSlugs, levelMap); err != nil {
				fmt.Printf("Error processing problem %d: %v\n", fid, err)
			}
		}(frontendID, titleSlug, difficultyLevel)
	}

	wg.Wait()
	return nil
}

func processProblem(domain string, frontendID int, titleSlug string, difficultyLevel int, langSlugs []string, levelMap map[int]string) error {
	// Create directory name with zero-padded ID
	dirName := filepath.Join(".", fmt.Sprintf("%04d-%s", frontendID, titleSlug))

	// Create directory
	if err := file.WriteDirectory(dirName); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirName, err)
	}

	// Get question details
	question, err := helper.GetQuestion(domain, titleSlug)
	if err != nil {
		return fmt.Errorf("failed to get question %s: %w", titleSlug, err)
	}

	// Prepare question config
	questionConfig := file.QuestionConfig{
		QuestionFrontendID: question.QuestionFrontendID,
		Title:              question.Title,
		Domain:             domain,
		TitleSlug:          question.TitleSlug,
		Content:            question.Content,
		Hints:              convertHintsToStringSlice(question.Hints),
	}

	// Write question README
	if err := file.WriteQuestion(dirName, questionConfig); err != nil {
		return fmt.Errorf("failed to write question for %s: %w", titleSlug, err)
	}

	// Write solution files for each requested language
	for _, lang := range langSlugs {
		// Find the code snippet for this language
		var snippet *config.CodeSnippet
		for _, cs := range question.CodeSnippets {
			if cs.LangSlug == lang {
				snippet = &cs
				break
			}
		}

		if snippet == nil {
			continue // Skip if language not found
		}

		// Write solution file
		if err := file.WriteSolution(dirName, lang, snippet.Code, config.Config{}); err != nil {
			fmt.Printf("Warning: failed to write solution for %s in %s: %v\n", titleSlug, lang, err)
		}
	}

	// Prepare question data for information.json
	questionData := map[string]interface{}{
		"similarQuestions": question.SimilarQuestions,
	}

	// Write information.json
	difficulty := levelMap[difficultyLevel]
	if difficulty == "" {
		difficulty = "Unknown"
	}

	if err := file.WriteInformation(dirName, questionData, difficulty); err != nil {
		return fmt.Errorf("failed to write information for %s: %w", titleSlug, err)
	}

	fmt.Printf("Successfully processed: %s\n", dirName)
	return nil
}

// convertHintsToStringSlice converts the hints interface{} to []string
func convertHintsToStringSlice(hints interface{}) []string {
	if hints == nil {
		return []string{}
	}

	switch h := hints.(type) {
	case []interface{}:
		result := make([]string, len(h))
		for i, hint := range h {
			if str, ok := hint.(string); ok {
				result[i] = str
			} else {
				result[i] = fmt.Sprintf("%v", hint)
			}
		}
		return result
	case []string:
		return h
	case string:
		// If it's a JSON string, try to unmarshal it
		var stringSlice []string
		if err := json.Unmarshal([]byte(h), &stringSlice); err == nil {
			return stringSlice
		}
		// If it's just a regular string, return it as a single-element slice
		return []string{h}
	default:
		return []string{}
	}
}
