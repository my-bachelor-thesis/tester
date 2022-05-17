package ioutils

const (
	FolderPerm = 0755
	FilePerm   = 0644
	BytesInMB  = 1e6
	KBytesInMB = 1e3
)

func BytesToMB(bytes float32) float32 {
	return float32(int(bytes/BytesInMB*100)) / 100
}

func KBytesToMB(bytes float32) float32 {
	return float32(int(bytes/KBytesInMB*100)) / 100
}
