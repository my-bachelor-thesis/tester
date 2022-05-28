package containers

import (
	"encoding/json"
	"fmt"
	"github.com/enriquebris/goconcurrentqueue"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"tester/internal/ioutils"
	"tester/internal/languages"
	"tester/internal/webserver_structs"
)

var fifo = goconcurrentqueue.NewFIFO()

func RunSolution(folderName string, lang languages.Language) (*webserver_structs.OutgoingJson, error) {
	containerName := <-GetFreeContainer()
	defer FreeContainer(containerName)

	res := webserver_structs.OutgoingJson{}

	gotFromScript, err := exec.Command("scripts/run_solution.sh", folderName, containerName, lang.String()).Output()
	if exitErr, ok := err.(*exec.ExitError); ok && exitErr.Stderr != nil {
		return nil, err
	}
	if err = json.Unmarshal(gotFromScript, &res); err != nil {
		return nil, err
	}

	res.MaxRamUsage = ioutils.KBytesToMB(res.MaxRamUsage)

	if languages.LanguageIsCompiled(lang) {
		res.BinarySize = ioutils.BytesToMB(res.BinarySize)
	}

	return &res, nil
}

func getFreeRam() (int, error) {
	out, err := exec.Command("free", "-m").Output()
	if err != nil {
		return 0, err
	}
	size := strings.Fields(strings.Split(string(out), "\n")[1])[1]
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return 0, err
	}
	return sizeInt, nil
}

func calculateResources() (numberOfContainers int, ramPerContainer int, err error) {
	numberOfContainers = runtime.NumCPU() * 2
	var ram int
	if ram, err = getFreeRam(); err != nil {
		return
	}
	ramPerContainer = ram / numberOfContainers
	return
}

func fillQue(numberOfContainers int) error {
	for i := 1; i <= numberOfContainers; i++ {
		if err := fifo.Enqueue(fmt.Sprintf("package-tester-%d", i)); err != nil {
			return err
		}
	}
	return nil
}

func StartAll() error {
	numberOfContainers, ramPerContainer, err := calculateResources()
	if err != nil {
		return err
	}

	if err := runScript("scripts/start_containers.sh",
		strconv.Itoa(numberOfContainers), strconv.Itoa(ramPerContainer)); err != nil {
		return err
	}

	return fillQue(numberOfContainers)
}

func runScript(path string, args ...string) error {
	if err := exec.Command("chmod", "+x", path).Run(); err != nil {
		return err
	}
	return exec.Command(path, args...).Run()
}

func GetFreeContainer() chan string {
	done := make(chan string)
	go func() {
		name, _ := fifo.DequeueOrWaitForNextElement()
		done <- name.(string)
	}()
	return done
}

func FreeContainer(name string) {
	_ = fifo.Enqueue(name)
}
