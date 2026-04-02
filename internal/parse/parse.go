package parse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1chooo/leetcode-crawler/config"
)

// ProblemIDs parses the --problem flag: single IDs, comma-separated, or start-end range.
func ProblemIDs(problemFlag string) ([]int, error) {
	var ids []int

	if strings.Contains(problemFlag, "-") {
		parts := strings.Split(problemFlag, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format: %s (expected format: start-end)", problemFlag)
		}

		start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start number in range: %s", parts[0])
		}

		end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end number in range: %s", parts[1])
		}

		if start > end {
			return nil, fmt.Errorf("start number (%d) cannot be greater than end number (%d)", start, end)
		}

		for i := start; i <= end; i++ {
			ids = append(ids, i)
		}
	} else {
		problemStrs := strings.Split(problemFlag, ",")
		for _, problemStr := range problemStrs {
			problemStr = strings.TrimSpace(problemStr)
			if problemStr == "" {
				continue
			}

			id, err := strconv.Atoi(problemStr)
			if err != nil {
				return nil, fmt.Errorf("invalid problem number: %s", problemStr)
			}
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		return nil, fmt.Errorf("no valid problem IDs found")
	}

	return ids, nil
}

// Languages parses comma-separated language slugs; defaults to python3 when empty.
func Languages(langFlag string) []string {
	langs := strings.Split(langFlag, ",")
	var result []string

	for _, lang := range langs {
		lang = strings.TrimSpace(lang)
		if lang != "" {
			result = append(result, config.NormalizeLangSlug(lang))
		}
	}

	if len(result) == 0 {
		return []string{"python3"}
	}

	return result
}
