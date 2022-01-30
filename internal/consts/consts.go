package consts

type Language int

const (
	Go Language = iota
	Python
)

const (
	FolderPerm = 0755
	FilePerm   = 0644
	BytesInMB  = 1e6
	KBytesInMB = 1e3
)

var (
	Languages         = [...]string{"go", "python"}
	CompiledLanguages = map[Language]struct{}{Go: {}}
)

func (l Language) String() string {
	return Languages[l]
}
