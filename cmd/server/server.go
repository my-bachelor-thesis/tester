package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"runtime"
	"tester/internal/consts"
	"tester/internal/containers"
	"tester/internal/handlers"
	"tester/internal/util"
)

const addr = ":4000"

func init() {
	if runtime.GOOS != "linux" {
		log.Fatal("can only run on Linux")
	}
}

func createAllFolders() {
	for _, language := range consts.Languages {
		err := os.MkdirAll(fmt.Sprintf("assets/user_solutions/%s", language), consts.FolderPerm)
		util.PanicIfErr(err)
	}
}

func main() {
	createAllFolders()

	err := containers.StartAll()
	util.PanicIfErr(err)

	e := echo.New()
	e.POST("/go", handlers.Golang)
	e.POST("/python", handlers.Python)

	// disable all CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Logger.Fatal(e.Start(addr))
}
