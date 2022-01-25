package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"tester/external/utils"
	"tester/internal/consts"
	"tester/internal/containers"
	"tester/internal/structs"
)

const addr = ":4000"

func golang(w http.ResponseWriter, r *http.Request) {
	in := &structs.IncomingJson{}
	out := &structs.OutgoingJson{}

	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	goFolder := "assets/user_solutions/go"
	testFile := fmt.Sprintf("%s/%d_test.go", goFolder, in.TaskId)
	userFolder := fmt.Sprintf("%s/user_%d", goFolder, in.UserId)
	solutionFile := fmt.Sprintf("%s/%d.go", userFolder, in.TaskId)

	if in.Type == "test" {
		if err = os.WriteFile(testFile, []byte(in.Code), consts.FilePerm); err != nil {
			out.ExitCode = 1
			out.Out = err.Error()
		}
	} else if in.Type == "code" {
		utils.CreateFoldersIfNotExist(userFolder)
		os.WriteFile(solutionFile, []byte(in.Code), consts.FilePerm)
		out, _ = containers.RunSolution(in.UserId, in.TaskId, containers.GetFreeContainer(), consts.Go.String())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	_res, _ := json.Marshal(out)
	w.Write(_res)
}
func main() {
	for _, language := range consts.Languages {
		utils.CreateFoldersIfNotExist("assets/user_solutions/" + language)
	}

	containers.StartAll()

	mux := http.NewServeMux()
	mux.HandleFunc("/go", golang)

	fmt.Println("listening on:", addr)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
