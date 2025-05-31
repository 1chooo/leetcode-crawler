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

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while crawling leetcode problems: \n%s\n", err)
		os.Exit(1)
	}
}
