package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/1chooo/leetcode-crawler/config"
	"io"
	"net/http"
)

func GetAllProblems(domain string) (interface{}, error) {
	resp, err := http.Get(domain + "/api/problems/all/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

/*
Example usage

	func main() {
		question, err := GetQuestion("https://leetcode.com", "two-sum")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Title: %s\n", question.Title)
		fmt.Printf("Title Slug: %s\n", question.TitleSlug)
		fmt.Printf("Question ID: %s\n", question.QuestionFrontendID)
		fmt.Printf("Number of code snippets: %d\n", len(question.CodeSnippets))
	}
*/
func GetQuestion(domain, titleSlug string) (*config.Question, error) {
	// Create the GraphQL request payload
	payload := config.DefaultConfig.QuestionDataQL(titleSlug)

	// Marshal the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP POST request
	url := fmt.Sprintf("%s/graphql", domain)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d - %s", resp.StatusCode, string(body))
	}

	var response config.GraphQLResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response.Data.Question, nil
}
