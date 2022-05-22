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
}
