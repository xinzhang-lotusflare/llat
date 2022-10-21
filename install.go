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

	// create .bdz folder
	userhome, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Fail to find user home dir")
	}

	bdzFolder := userhome + "/.bdz"
	if _, err := os.Stat(bdzFolder); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating folder: ~/.bdz")
		if err := os.Mkdir(bdzFolder, os.ModePerm); err != nil {
			log.Println(err.Error())
			return fmt.Errorf("Fail to create dir:" + "~/.bdz")
		}
		if allowPrintInfo() {
			fmt.Println("Created folder: ~/.bdz")
		}
	}

	bashExecFile := bdzFolder + "/bash"
	if _, err := os.Stat(bashExecFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(bashExecFile); err != nil {
			log.Println(err.Error())
		  return fmt.Errorf("Fail to create 'bash' file under ~/.bdz")
		}
	}

	if err := os.WriteFile(bashExecFile, []byte(bashPath), 0644); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Fail to write to 'bash' file under ~/.bdz")
	}

	fmt.Println("Installed")
	return nil
}
