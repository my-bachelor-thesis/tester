package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tester/internal/languages"
	"tester/internal/structs"
)

func Golang(c echo.Context) error {
	in := &structs.IncomingJson{}

	if err := c.Bind(in); err != nil {
		return err
	}

	out, err := writeToFilesAndRun(languages.Go, []fileToWrite{
		{Name: "main.go", Content: []string{in.Solution}},
		{Name: "main_test.go", Content: []string{in.Test}},
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
	//in := &structs.IncomingJson{}
	//out := &structs.OutgoingJson{}
	//
	//if err := c.Bind(in); err != nil {
	//	return err
	//}
	//
	//path, err := os.MkdirTemp("assets/user_solutions/go", "*")
	//if err != nil {
	//	return err
	//}
	//defer os.RemoveAll(path)
	//
	//if err = os.WriteFile(fmt.Sprintf("%s/main_test.go", path), []byte(in.Test), ioutils.FilePerm); err != nil {
	//	return err
	//}
	//if err = os.WriteFile(fmt.Sprintf("%s/main.go", path), []byte(in.Solution), ioutils.FilePerm); err != nil {
	//	return err
	//}
	//
	//out, err = containers.RunSolution(filepath.Base(path), languages.Go)
	//if err != nil {
	//	return err
	//}
	//
	//return c.JSON(http.StatusOK, out)
}
