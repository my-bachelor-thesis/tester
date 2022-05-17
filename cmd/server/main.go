package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/exec"
	"runtime"
	"tester/internal/config"
	"tester/internal/containers"
	"tester/internal/handlers"
	"tester/internal/ioutils"
	"tester/internal/languages"
)

func init() {
	if runtime.GOOS != "linux" {
		log.Fatal("can only run on Linux")
	}

	xxdCmd := exec.Command("xxd", "-v")
	if err := xxdCmd.Run(); err != nil {
		log.Fatal("didn't find xxd")
	}

	ydCmd := exec.Command("yq", "-V")
	if err := ydCmd.Run(); err != nil {
		log.Fatal("didn't find yq")
	}

	if runtime.NumCPU() < 2 {
		log.Fatal("can only run on more than 2 cores")
	}
}

func createAllFolders(e *echo.Echo) {
	for _, language := range languages.Languages {
		logAndExitIfErr(e, os.MkdirAll(fmt.Sprintf("assets/user_solutions/%s", language), ioutils.FolderPerm))
	}
}

func main() {
	e := echo.New()

	if config.GetInstance().IsProduction {
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())
	}

	createAllFolders(e)

	logAndExitIfErr(e, config.LoadConfig())

	logAndExitIfErr(e, containers.StartAll())

	e.POST("/go", handlers.Golang)
	e.POST("/python", handlers.Python)
	e.POST("/cpp", handlers.Cpp)
	e.POST("/javascript", handlers.Javascript)

	// disable all CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.GetInstance().Port)))
}

func logAndExitIfErr(e *echo.Echo, err error) {
	if err != nil {
		e.Logger.Fatal(err)
	}
}
