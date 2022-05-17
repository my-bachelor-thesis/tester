package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tester/internal/languages"
	"tester/internal/structs"
)

func Python(c echo.Context) error {
	in := &structs.IncomingJson{}

	if err := c.Bind(in); err != nil {
		return err
	}

	out, err := writeToFilesAndRun(languages.Python, []fileToWrite{
		{Name: "main.py", Content: []string{in.Solution, in.Test}},
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
	//path, err := os.MkdirTemp("assets/user_solutions/python", "*")
	//if err != nil {
	//	return err
	//}
	//defer os.RemoveAll(path)
	//solutionFile := fmt.Sprintf("%s/main.py", path)
	//if err = os.WriteFile(solutionFile, []byte(in.Solution), ioutils.FilePerm); err != nil {
	//	return err
	//}
	//if err = ioutils.AppendToFile(solutionFile, in.Test); err != nil {
	//	return err
	//}
	//
	//out, err = containers.RunSolution(filepath.Base(path), languages.Python)
	//if err != nil {
	//	return err
	//}
	//
	//return c.JSON(http.StatusOK, out)
}
