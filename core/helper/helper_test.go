package helper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllProblems(t *testing.T) {
	payload := map[string]interface{}{
		"stat_status_pairs": []map[string]interface{}{
			{
				"difficulty": map[string]int{"level": 1},
				"stat": map[string]interface{}{
					"frontend_question_id":   1,
					"question__title_slug":   "two-sum",
				},
			},
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/problems/all/" {
			http.NotFound(w, r)
			return
		}
		_ = json.NewEncoder(w).Encode(payload)
	}))
	defer srv.Close()

	out, err := GetAllProblems(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(out.StatStatusPairs) != 1 {
		t.Fatalf("pairs: %d", len(out.StatStatusPairs))
	}
	if out.StatStatusPairs[0].Stat.FrontendQuestionID != 1 {
		t.Fatalf("id: %d", out.StatStatusPairs[0].Stat.FrontendQuestionID)
	}
}

func TestGetQuestion(t *testing.T) {
	graphqlBody := map[string]interface{}{
		"data": map[string]interface{}{
			"question": map[string]interface{}{
				"title":                "Two Sum",
				"titleSlug":            "two-sum",
				"questionFrontendId":   "1",
				"content":              "<p>x</p>",
				"similarQuestions":     "[]",
				"hints":                []interface{}{},
				"codeSnippets":         []interface{}{},
				"translatedTitle":      "",
				"translatedContent":    "",
				"stats":                nil,
			},
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" || r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(graphqlBody)
	}))
	defer srv.Close()

	q, err := GetQuestion(srv.URL, "two-sum")
	if err != nil {
		t.Fatal(err)
	}
	if q.Title != "Two Sum" || q.TitleSlug != "two-sum" {
		t.Fatalf("question: %+v", q)
	}
	if q.QuestionFrontendID != "1" {
		t.Fatalf("frontend id: %q", q.QuestionFrontendID)
	}
}

func TestGetQuestion_HTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusInternalServerError)
	}))
	defer srv.Close()

	if _, err := GetQuestion(srv.URL, "x"); err == nil {
		t.Fatal("expected error")
	}
}
