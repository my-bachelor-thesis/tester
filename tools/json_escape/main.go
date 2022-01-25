package main

import (
	"os"
	"tester/external/utils"
	"tester/internal/consts"
)



func main() {
	err := os.Chdir("tools/json_escape")
	utils.PanicIfErr(err)

	data, err := os.ReadFile("orig")
	utils.PanicIfErr(err)

	err = os.WriteFile("cleaned", []byte(utils.EscapeJson(string(data))), consts.FilePerm)
	utils.PanicIfErr(err)
}
