# LeetCode Problems Crawler

A Go CLI (Command Line Interface) tool to crawl LeetCode problems and generate code files in various programming languages with different naming schemes.

## Install

From the repository root (requires [Go](https://go.dev/dl/) 1.21+):

```bash
go install github.com/1chooo/leetcode-crawler/cmd/leetcode-crawler@latest
```

Or build / install with an embedded version string:

```bash
make install
# optional: VERSION=$(git describe --tags --always) make install
```

The binary is named `leetcode-crawler`. Use `leetcode-crawler -v` to print version and author.

**Naming:** `--naming` controls how the **problem directory** is named (the segment after the zero-padded problem id, derived from the title slug). Solution source files are always `Solution.<ext>` to match LeetCode templates. Accepted styles include `kebab-case`, `snake_case`, `camelCase` (lower camel), and `pascalCase` (upper camel). The flag `go` is accepted as an alias for LeetCode’s `golang` slug.

## 📖 Usage

```text
leetcode-crawler crawl --problem <problem_id> --lang <language> --path <output_path> --naming <naming_scheme>
```

- `--problem`: one id (`1`), comma-separated (`1,2,3`), or range (`1-5`).
- `--lang`: comma-separated LeetCode slugs (e.g. `python3`, `golang`, `java`). Default: `python3`.

### Examples (installed binary)

Example 1: problem 1, Go (`go` → `golang`), snake_case directory names

```bash
leetcode-crawler crawl --problem 1 --lang go --path ./problems/ --naming snake_case
```

Example 2: problem 2, Python3, camelCase directory names

```bash
leetcode-crawler crawl --problem 2 --lang python3 --path ./problems/ --naming camelCase
```

Example 3: problem 3, Java, pascalCase directory names

```bash
leetcode-crawler crawl --problem 3 --lang java --path ./problems/ --naming pascalCase
```

Example 4: problem 4, C++, snake_case

```bash
leetcode-crawler crawl --problem 4 --lang cpp --path ./problems/ --naming snake_case
```

Example 5: problem 5, Rust, kebab-case

```bash
leetcode-crawler crawl --problem 5 --lang rust --path ./problems/ --naming kebab-case
```

### Runnable scripts

See [examples/](examples/): `crawl-single.sh`, `crawl-range.sh`, and `crawl-multilang.sh` run `go run ./cmd/leetcode-crawler` and write under `examples/out/` (ignored by git). Run from repo root, e.g. `bash examples/crawl-single.sh`.

## Tests

```bash
make test
# or: go test ./...
```

## 📱 Contact

> **Chun-Ho (Hugo) Lin**
> 
> <aside>
>   📩 E-mail: <a href="mailto:hugo970217@gmail.com">hugo970217@gmail.com</a>
> <br>
>   🧳 Linkedin: <a href="https://www.linkedin.com/in/1chooo/">in/1chooo</a>
> <br>
>   👨🏻‍💻 GitHub: <a href="https://github.com/1chooo">@1chooo</a>
>    
> </aside>


## 🪪 License

Released under [MIT](./LICENSE) by [@1chooo](https://www.1chooo.com)

This software can be modified and reused without restriction. The original license must be included with any copies of this software. If a significant portion of the source code is used, please provide a link back to this repository.

Made with 💙 by [@1chooo](https://www.1chooo.com)

## 🙏🏻 Special Thanks

- [Cobra](https://cobra.dev/)
- [Build CLI Apps with Go and Cobra](https://www.jetbrains.com/guide/go/tutorials/cli-apps-go-cobra/)
- [vv13/leetcode-problems-crawler](https://github.com/vv13/leetcode-problems-crawler)

