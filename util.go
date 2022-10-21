package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const isReleaseMode = mode == "release"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func allowPrintInfo() bool {
	return !(isReleaseMode && !verbose)
}

func writeConfigFile() (string, error) {
	interfaceName := "bulldozer"
	userhome, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("Fail to find user home dir")
	}
	configFileName := userhome + "/.bdz/" + interfaceName + ".conf"
	config := []byte(wgConfig())

	if _, err := os.Stat(configFileName); errors.Is(err, os.ErrNotExist) {
	  if _, err = os.Create(configFileName); err != nil {
			log.Println(err.Error())
		  return "", fmt.Errorf("Fail to create config file")
		}
	}

	if err := os.WriteFile(configFileName, config, 0644); err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("Fail to write to config file")
	} else {
		return configFileName, nil
	}
}

func clearConfigFile(filename string) error {
	empty := ""
	if err := os.WriteFile(filename, []byte(empty), 0644); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Fail to clear the config file")
	} else {
		return nil
	}
}

func getBash() (string, error) {
	userhome, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("Fail to find user home dir")
	}
	file := userhome + "/.bdz/bash"
	bytes, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("Fail to find bash executable from file: " + file)
		return "", err
	}

	bashExec := strings.TrimSpace(string(bytes))
	return bashExec, nil
}

func notRoot() bool {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	user := strings.TrimSpace(string(stdout))
	if allowPrintInfo() {
		fmt.Println("Running as " + user)
	}

	if user == "root" {
		return false
	}
	return true
}
