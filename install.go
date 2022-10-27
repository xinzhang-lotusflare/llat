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

	llatWorkspace := userhome + "/.llat"
	if _, err := os.Stat(llatWorkspace); errors.Is(err, os.ErrNotExist) {
		log.Println("Creating folder: ~/.llat")
		if err := os.Mkdir(llatWorkspace, os.ModePerm); err != nil {
			log.Println(err.Error())
			return fmt.Errorf("Fail to create dir:" + "~/.llat")
		}
		if allowPrintInfo() {
			fmt.Println("Created folder: ~/.llat")
		}
	}

	bashPathInBytes := func() ([]byte, error) {
		return []byte(bashPath), nil
	}
	err = writeFileInWorkspace(llatWorkspace, "bash", bashPathInBytes)
	if err != nil {
		return err
	}

	readConfig := func() ([]byte, error) {
		configPath := cCtx.String("llat-config")
		log.Println("Copy config from: " + configPath)
		bytes, err := os.ReadFile(configPath)
		if err != nil {
			err = fmt.Errorf("Fail to copy from file: " + configPath)
			return nil, err
		}
		return bytes, nil
	}
	err = writeFileInWorkspace(llatWorkspace, "config", readConfig)
	if err != nil {
		return err
	}

	fmt.Println("Installed")
	return nil
}
