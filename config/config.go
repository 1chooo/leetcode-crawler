package config

type Config struct {
	CN       SiteConfig
	EN       SiteConfig
	Level    LevelMap
	Language LanguageMap
	naming   NamingMap
}

var DefaultConfig = Config{
	CN: SiteConfig{
		Domain: "https://leetcode-cn.com",
	},
	EN: SiteConfig{
		Domain: "https://leetcode.com",
	},
	Level: LevelMap{
		One:   "easy",
		Two:   "medium",
		Three: "hard",
	},
	Language: LanguageMap{
		Java:       "Java",
		JavaScript: "JavaScript",
		Python3:    "Python3",
		CPP:        "C++",
		C:          "C",
		Golang:     "Go",
		Rust:       "Rust",
		TypeScript: "TypeScript",
	},
	naming: NamingMap{
		SnakeCase:      "snake_case",
		CamelCase:      "CamelCase",
		LowerCamelCase: "lowerCamelCase",
		UpperCamelCase: "UpperCamelCase",
		KebabCase:      "kebab-case",
	},
}
