package containers

import (
	"encoding/json"
	"os/exec"
	"tester/internal/consts"
	"tester/internal/structs"
	"tester/internal/util"
)

func RunSolution(folderName string, lang consts.Language) (*structs.OutgoingJson, error) {
	containerName := GetFreeContainer()

	res := structs.OutgoingJson{}

	gotFromScript, err := exec.Command("scripts/run_solution.sh", folderName, containerName, lang.String()).Output()
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.Stderr != nil {
		return nil, err
	}
	if err = json.Unmarshal(gotFromScript, &res); err != nil {
		return nil, err
	}

	res.MaxRamUsage = util.KBytesToMB(res.MaxRamUsage)

	if util.LanguageIsCompiled(lang) {
		res.BinarySize = util.BytesToMB(res.BinarySize)
	}

	return &res, nil
}

func StartAll() error {
	if err := exec.Command("chmod", "+x", "scripts/start_containers.sh").Run(); err != nil {
		return err
	}
	return exec.Command("scripts/start_containers.sh").Run()
}

func GetFreeContainer() string {
	return "package-tester-1"
}
