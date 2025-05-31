package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "leetcode-crawler",
	Short: "leetcode-crawler is a cli tool for crawling leetcode problems.",
	Long:  "leetcode-crawler is a command line tool that allows you to crawl and manage leetcode problems efficiently.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	version = "0.0.1"
	author  = "Chun-Ho (Hugo) Lin"
)

func Execute() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version information and quit")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Printf("Version: %s\nAuthor: %s\n", version, author)
			os.Exit(0)
		}
	}

	rootCmd.PersistentFlags().StringP("problem", "p", "", "crawl leetcode problems by problem number or tag, e.g., 1, 1-10, array, string")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		problem, _ := cmd.Flags().GetString("problem")
		if problem == "" {
			fmt.Fprintln(os.Stderr, "Error: You must specify a problem number or tag using the --problem flag.")
			os.Exit(1)
		}
		// fmt.Printf("Crawling problems for: %s\n", problem)
	}

	rootCmd.PersistentFlags().StringP("international", "i", "", "Specify the international version of leetcode, options: en, cn")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		international, _ := cmd.Flags().GetString("international")
		if international == "" {
			international = "en" // Default to English if not specified
		}
		if international != "en" && international != "cn" {
			fmt.Fprintln(os.Stderr, "Error: Invalid international option. Use 'en' for English or 'cn' for Chinese.")
			os.Exit(1)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while crawling leetcode problems: \n%s\n", err)
		os.Exit(1)
	}
}
