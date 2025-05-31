package config

type SiteConfig struct {
	Domain string
}

type LevelMap struct {
	One   string
	Two   string
	Three string
}

type LanguageMap struct {
	Java       string
	JavaScript string
	Python3    string
	CPP        string
	C          string
	Golang     string
	Rust       string
	TypeScript string
}

type NamingMap struct {
	SnakeCase      string
	CamelCase      string
	LowerCamelCase string
	UpperCamelCase string
	KebabCase      string
}
