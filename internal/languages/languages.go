package languages

type Language int

const (
	Go Language = iota
	Python
	Javascript
	Cpp
)

var (
	Languages         = [...]string{"go", "python", "javascript", "cpp"}
	CompiledLanguages = map[Language]struct{}{Go: {}, Cpp: {}}
)

func (l Language) String() string {
	return Languages[l]
}

func LanguageIsCompiled(lang Language) bool {
	_, ok := CompiledLanguages[lang]
	return ok
}
