package consts

type Language int

func (l Language) String() string {
	return Languages[l]
}

var Languages = [...]string{"go", "python"}

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
