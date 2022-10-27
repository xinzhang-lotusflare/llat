package main

import (
	"fmt"
	"log"
	"os/exec"
)

const wgQuick = "wg-quick"
const up = "up"
const down = "down"

func switchInterface(toggle bool) error {
	if notRoot() {
		err := "Need root permission"
		fmt.Println(err)
		return fmt.Errorf(err)
	}

	configFile, err := writeWgConfigFile()
	if err != nil {
		return err
	}
	defer clearConfigFile(configFile)

	bash, err := getBash()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var motion string
	if toggle {
		motion = up
	} else {
		motion = down
	}

	if allowPrintInfo() {
		fmt.Println("motion: " + motion)
	}

	var running *exec.Cmd
	if isReleaseMode {
		running = exec.Command(bash, wgQuick, motion, configFile)
	} else {
		running = exec.Command("bash", "-c", "echo", motion)
	}

	output, err := running.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		log.Println(string(output))
		return err
	}
	log.Println(string(output))
	return nil
}
