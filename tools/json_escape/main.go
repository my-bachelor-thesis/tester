package main

import (
	"encoding/json"
	"log"
	"os"
	"tester/internal/ioutils"
)

func main() {
	exitIfErr(os.Chdir("tools/json_escape"))

	data, err := os.ReadFile("orig")
	exitIfErr(err)

	cleaned, err := escapeJson(string(data))
	exitIfErr(err)

	exitIfErr(os.WriteFile("cleaned", []byte(cleaned), ioutils.FilePerm))
}

func escapeJson(i string) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	s := string(b)
	return s[1 : len(s)-1], nil
}

func exitIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
