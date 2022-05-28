package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"tester/internal/containers"
	"tester/internal/ioutils"
	"tester/internal/languages"
	"tester/internal/webserver_structs"
)

type fileToWrite struct {
	Name    string
	Content []string
}

func writeToFilesAndRun(lang languages.Language, filesToWrite []fileToWrite) (*webserver_structs.OutgoingJson, error) {
	path, err := os.MkdirTemp(fmt.Sprintf("assets/user_solutions/%s", lang), "*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(path)

	for _, file := range filesToWrite {
		fileName := fmt.Sprintf("%s/%s", path, file.Name)
		var startOfFile string

		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, ioutils.FilePerm)
		if err != nil {
			return nil, err
		}

		for j, content := range file.Content {
			if j != 0 {
				startOfFile = "\n"
			}
			if _, err := f.WriteString(fmt.Sprintf("%s%s", startOfFile, content)); err != nil {
				return nil, err
			}
		}
	}

	out := &webserver_structs.OutgoingJson{}
	out, err = containers.RunSolution(filepath.Base(path), lang)
	if err != nil {
		return nil, err
	}

	return out, nil
}
