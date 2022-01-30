package main

import (
	"os"
	"tester/internal/consts"
	"tester/internal/util"
)

func main() {
	err := os.Chdir("tools/json_escape")
	util.PanicIfErr(err)

	data, err := os.ReadFile("orig")
	util.PanicIfErr(err)

	cleaned, err := util.EscapeJson(string(data))
	util.PanicIfErr(err)

	err = os.WriteFile("cleaned", []byte(cleaned), consts.FilePerm)
	util.PanicIfErr(err)
}
