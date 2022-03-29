package languages

type Language int

const (
	Go Language = iota
	Python
)

var (
	Languages         = [...]string{"go", "python"}
	CompiledLanguages = map[Language]struct{}{Go: {}}
)

func (l Language) String() string {
	return Languages[l]
}

func LanguageIsCompiled(lang Language) bool {
	_, ok := CompiledLanguages[lang]
	return ok
}
