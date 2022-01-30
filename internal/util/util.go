package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"tester/internal/consts"
)

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

func EscapeJson(i string) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	s := string(b)
	return s[1 : len(s)-1], nil
}

func LanguageIsCompiled(lang consts.Language) bool {
	_, ok := consts.CompiledLanguages[lang]
	return ok
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func AppendToFile(filename, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, consts.FilePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("\n%s", text))
	return err
}
