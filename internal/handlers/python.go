package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"path/filepath"
	"tester/internal/containers"
	"tester/internal/ioutils"
	"tester/internal/languages"
	"tester/internal/structs"
)

func Python(c echo.Context) error {
	in := &structs.IncomingJson{}
	out := &structs.OutgoingJson{}

	if err := c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	path, err := os.MkdirTemp("assets/user_solutions/python", "*")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer os.RemoveAll(path)
	solutionFile := fmt.Sprintf("%s/main.py", path)
	if err = os.WriteFile(solutionFile, []byte(in.Solution), ioutils.FilePerm); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = ioutils.AppendToFile(solutionFile, in.Test); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	out, err = containers.RunSolution(filepath.Base(path), languages.Python)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, out)
}
