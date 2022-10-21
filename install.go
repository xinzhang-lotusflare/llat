package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func testBash(bashPath string) bool {
	testing := exec.Command(bashPath, "-c", "echo", "hello")
	output, err := testing.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		log.Println(string(output))
		return false
	}
	log.Println(string(output))
	return true
}

func install(cCtx *cli.Context) error {
	log.Println("Installing")
	bashPath := cCtx.String("bash-executable")
	log.Println("bash execuatble path: " + bashPath)

	if bashPath == "" {
		errMsg := "The bash execuatble path is required"
		fmt.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	if !testBash(bashPath) {
		errMsg := "Illegal bash execuatble"
		fmt.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	// create .llat folder
	userhome, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Fail to find user home dir")
	}

	llatFolder := userhome + "/.llat"
	if _, err := os.Stat(llatFolder); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating folder: ~/.llat")
		if err := os.Mkdir(llatFolder, os.ModePerm); err != nil {
			log.Println(err.Error())
			return fmt.Errorf("Fail to create dir:" + "~/.llat")
		}
		if allowPrintInfo() {
			fmt.Println("Created folder: ~/.llat")
		}
	}

	bashExecFile := llatFolder + "/bash"
	if _, err := os.Stat(bashExecFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(bashExecFile); err != nil {
			log.Println(err.Error())
			return fmt.Errorf("Fail to create 'bash' file under ~/.llat")
		}
	}

	if err := os.WriteFile(bashExecFile, []byte(bashPath), 0644); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Fail to write to 'bash' file under ~/.llat")
	}

	fmt.Println("Installed")
	return nil
}
