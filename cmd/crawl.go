package cmd

import (
	"fmt"
	"os"

	"github.com/1chooo/leetcode-crawler/config"
	"github.com/1chooo/leetcode-crawler/core/crawl"
	"github.com/1chooo/leetcode-crawler/internal/parse"
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
		namingFlag, _ := cmd.Flags().GetString("naming")
		naming := config.NormalizeNamingConvention(namingFlag)

		problemIDs, err := parse.ProblemIDs(problemFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing problem IDs: %v\n", err)
			os.Exit(1)
		}

		languages := parse.Languages(langFlag)

		// Display what we're about to crawl
		fmt.Printf("Crawling problems: %v\n", problemIDs)
		fmt.Printf("Languages: %v\n", languages)
		fmt.Printf("Output path: %s\n", pathFlag)
		fmt.Printf("Directory naming: %s\n", naming)

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
		if err := crawl.ProblemCrawler(problemIDs, languages, naming); err != nil {
			fmt.Fprintf(os.Stderr, "Error crawling problems: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Crawling completed successfully!")
	},
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
	crawlCmd.Flags().StringP("naming", "n", "kebab-case", "Directory name style: kebab-case, snake_case, camelCase, pascalCase")
	rootCmd.AddCommand(crawlCmd)
}
