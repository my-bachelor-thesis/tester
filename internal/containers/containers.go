package containers

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"tester/external/utils"
	"tester/internal/consts"
	"tester/internal/structs"
)

func RunSolution(userId, taskId int, container, lang string) (*structs.OutgoingJson, error) {
	res := structs.OutgoingJson{}

	if lang == consts.Go.String() {
		gotFromScript, _ := exec.Command("scripts/run_solution.sh", strconv.Itoa(userId), strconv.Itoa(taskId), container, lang).Output()
		//fmt.Println(string(out))
		//fmt.Println(err)
		json.Unmarshal(gotFromScript, &res)
		//fmt.Println(err)
		res.BinarySize = utils.BytesToMB(res.BinarySize)
		res.MaxRamUsage = utils.KBytesToMB(res.MaxRamUsage)
		//utils.PrettyPrintJson(res)
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
	return "tester_1"
}
