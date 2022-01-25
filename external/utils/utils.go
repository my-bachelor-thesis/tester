package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"tester/internal/consts"
)

func FolderExists(path string) bool {
	s, err := os.Stat(path)
	return !os.IsNotExist(err) && s.IsDir()
}

func FileExists(path string) bool {
	s, err := os.Stat(path)
	return !os.IsNotExist(err) && !s.IsDir()
}

func CreateFoldersIfNotExist(path string) error {
	if !FolderExists(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BytesToMB(bytes float32) float32 {
	return float32(int(bytes/consts.BytesInMB*100)) / 100
}

func KBytesToMB(bytes float32) float32 {
	return float32(int(bytes/consts.KBytesInMB*100)) / 100
}

func PrettyPrintJson(data *interface{}) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "/t")

	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(buffer.String())
}

func EscapeJson(i string) string {
	b, err := json.Marshal(i)
	PanicIfErr(err)
	s := string(b)
	return s[1 : len(s)-1]
}
