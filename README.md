# LeetCode Problems Crawler

A Go CLI (Command Line Interface) tool to crawl LeetCode problems and generate code files in various programming languages with different naming schemes.

## 📖 Usage

pattern:

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

