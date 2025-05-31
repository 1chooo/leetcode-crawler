package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var crawlCmd = &cobra.Command{
	Use:     "crawl",
	Aliases: []string{"c"},
	Short:   "Crawl leetcode problems",
	Long:    "Crawl leetcode problems based on the specified problem number or tag.",
	Example: "leetcode-crawler crawl --problem 5 --lang rust --path ./problems/ --naming kebab-case",
	Run: func(cmd *cobra.Command, args []string) {
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
	
	crawlCmd.Flags().StringP("lang", "l", "Python3", "Programming language to use (default: Python3)")
	crawlCmd.Flags().StringP("path", "d", "./problems/", "Directory to save the crawled problems")
	crawlCmd.Flags().StringP("naming", "n", "kebab-case", "Naming convention for the problem files (default: kebab-case)")
	rootCmd.AddCommand(crawlCmd)
}
