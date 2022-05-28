package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tester/internal/languages"
	"tester/internal/webserver_structs"
)

func Javascript(c echo.Context) error {
	in := &webserver_structs.IncomingJson{}

	if err := c.Bind(in); err != nil {
		return err
	}

	out, err := writeToFilesAndRun(languages.Javascript, []fileToWrite{
		{Name: "main.js", Content: []string{in.Solution, in.Test}},
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}
