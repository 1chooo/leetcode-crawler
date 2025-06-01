package main

import (
	// "github.com/1chooo/leetcode-crawler/cmd"
	"github.com/1chooo/leetcode-crawler/core/crawl"
)

func main() {
	crawl.ProblemCrawler([]int{1, 2, 3, 4, 5}, []string{"java", "python3", "cpp", "c", "golang", "rust", "typescript"})
	// cmd.Execute()
}
