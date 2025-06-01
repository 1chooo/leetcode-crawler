package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/1chooo/leetcode-crawler/core/crawl"
	"github.com/spf13/cobra"
)

var crawlCmd = &cobra.Command{
	Use:     "crawl",
	Aliases: []string{"c"},
	Short:   "Crawl leetcode problems",
	Long:    "Crawl leetcode problems based on the specified problem number or tag.",
	Example: "leetcode-crawler crawl --problem 5 --lang rust or leetcode-crawler crawl --problem 1-5 --lang rust,python3",
	Run: func(cmd *cobra.Command, args []string) {
		// Get flag values
		problemFlag, _ := cmd.Flags().GetString("problem")
		langFlag, _ := cmd.Flags().GetString("lang")
		pathFlag, _ := cmd.Flags().GetString("path")

		// Parse problem IDs
		problemIDs, err := parseProblemIDs(problemFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing problem IDs: %v\n", err)
			os.Exit(1)
		}

		// Parse languages
		languages := parseLanguages(langFlag)

		// Display what we're about to crawl
		fmt.Printf("Crawling problems: %v\n", problemIDs)
		fmt.Printf("Languages: %v\n", languages)
		fmt.Printf("Output path: %s\n", pathFlag)

		// Change to the specified directory if needed
		if pathFlag != "./" && pathFlag != "" {
			if err := os.MkdirAll(pathFlag, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", pathFlag, err)
				os.Exit(1)
			}
			if err := os.Chdir(pathFlag); err != nil {
				fmt.Fprintf(os.Stderr, "Error changing to directory %s: %v\n", pathFlag, err)
				os.Exit(1)
			}
		}

		// Call the crawler
		if err := crawl.ProblemCrawler(problemIDs, languages); err != nil {
			fmt.Fprintf(os.Stderr, "Error crawling problems: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Crawling completed successfully!")
	},
}

// parseProblemIDs parses the problem flag and returns a slice of problem IDs
func parseProblemIDs(problemFlag string) ([]int, error) {
	var ids []int

	// Handle range format (e.g., "1-5")
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
		// Handle single number or comma-separated numbers
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

// parseLanguages parses the language flag and returns a slice of language slugs
func parseLanguages(langFlag string) []string {
	// Split by comma and clean up
	langs := strings.Split(langFlag, ",")
	var result []string

	for _, lang := range langs {
		lang = strings.TrimSpace(lang)
		if lang != "" {
			// Convert to lowercase for consistency with LeetCode API
			result = append(result, strings.ToLower(lang))
		}
	}

	// If no languages specified, default to python3
	if len(result) == 0 {
		result = []string{"python3"}
	}

	return result
}

func init() {
	crawlCmd.PersistentFlags().StringP("problem", "p", "", "crawl leetcode problems by problem number or tag, e.g., 1, 1-10, array, string")
	crawlCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		problem, _ := cmd.Flags().GetString("problem")
		if problem == "" {
			fmt.Fprintln(os.Stderr, "Error: You must specify a problem number or tag using the --problem flag.")
			os.Exit(1)
		}
		fmt.Printf("Crawling problems for: %s\n", problem)
	}

	crawlCmd.Flags().StringP("lang", "l", "python3", "Programming language to use (default: python3)")
	crawlCmd.Flags().StringP("path", "d", "./problems/", "Directory to save the crawled problems")
	crawlCmd.Flags().StringP("naming", "n", "kebab-case", "Naming convention for the problem files (default: kebab-case)")
	rootCmd.AddCommand(crawlCmd)
}
