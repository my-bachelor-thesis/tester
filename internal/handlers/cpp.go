package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tester/internal/languages"
	"tester/internal/structs"
)

func Cpp(c echo.Context) error {
	in := &structs.IncomingJson{}

	if err := c.Bind(in); err != nil {
		return err
	}

	testHeader := `#define CATCH_CONFIG_MAIN
#include "/usr/include/catch2/catch.hpp"`

	out, err := writeToFilesAndRun(languages.Cpp, []fileToWrite{
		{Name: "main.cpp", Content: []string{
			in.Solution,
			testHeader,
			in.Test,
		}},
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}
