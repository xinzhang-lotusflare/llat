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

func writeFileInWorkspace(workspace string, fileName string, getContent func() ([]byte, error)) error {
	configFile := workspace + "/" + fileName
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(configFile); err != nil {
			log.Println(err.Error())
			return fmt.Errorf("Fail to create '" + fileName + "' file under:" + workspace)
		}
	}

	content, err := getContent()
	if err != nil {
		return nil
	}

	if err := os.WriteFile(configFile, content, 0644); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Fail to write to '" + fileName + "' file under:" + workspace)
	}

	fmt.Println("Installed")
	return nil
}

func writeWgConfigFile() (string, error) {
	interfaceName := "llat"
	userhome, err := os.UserHomeDir()
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("Fail to find user home dir")
	}

	llatWorkspace := userhome + "/.llat/"
	if _, err := os.Stat(llatWorkspace); errors.Is(err, os.ErrNotExist) {
		errMsg := "You need to run 'install' first. The workspace doesn't exist: ~/.llat"
		fmt.Println(errMsg)
		return "", fmt.Errorf(errMsg)
	}

	getWgConfig := func() ([]byte, error) {
		llatConfig, err := getConfig(llatWorkspace)
		if err != nil {
			errMsg := "fail to get llat config"
			fmt.Println(errMsg)
			return nil, fmt.Errorf(errMsg)
		}
		wgConfig := []byte(wgConfig(
			llatConfig.ipAddress,
			llatConfig.privateKey,
			llatConfig.publicKey,
			llatConfig.presharedKey,
		))
		return wgConfig, nil
	}

	err = writeFileInWorkspace(llatWorkspace, interfaceName+".conf", getWgConfig)
	if err != nil {
		log.Println(err.Error())
		return "", fmt.Errorf("Fail to write to WireGuard config file")
	}

	wgConfigFileName := llatWorkspace + interfaceName + ".conf"
	return wgConfigFileName, nil
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
	file := userhome + "/.llat/bash"
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
