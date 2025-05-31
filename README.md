# LeetCode Problems Crawler

This project is inspired by [vv13/leetcode-problems-crawler](https://github.com/vv13/leetcode-problems-crawler).

## Usage

```go
leetcode-crawler crawl --problem <problem_id> --lang <language> --path <output_path> --naming <naming_scheme>
```

Example 1:

to crawl problem 1 in Go language with snake_case naming scheme

```bash
leetcode-crawler crawl --problem 1 --lang go --path ./problems/ --naming snake_case
```

Example 2:

To crawl problem 2 in Python3 language with camelCase naming scheme

```bash
leetcode-crawler crawl --problem 2 --lang python3 --path ./problems/ --naming camelCase
```

Example 3:

To crawl problem 3 in Java language with PascalCase naming scheme

```bash
leetcode-crawler crawl --problem 3 --lang java --path ./problems/ --naming pascalCase
```

Example 4: 

To crawl problem 4 in C++ language with snake_case naming scheme

```bash
leetcode-crawler crawl --problem 4 --lang cpp --path ./problems/ --naming snake_case
```

Example 5:

To crawl problem 5 in Rust language with kebab-case naming scheme

```bash
leetcode-crawler crawl --problem 5 --lang rust --path ./problems/ --naming kebab-case
```


- [Cobra](https://cobra.dev/)
- [Build CLI Apps with Go and Cobra](https://www.jetbrains.com/guide/go/tutorials/cli-apps-go-cobra/)


