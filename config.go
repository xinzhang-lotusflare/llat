package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type LlatConfig struct {
	ipAddress    string
	privateKey   string
	publicKey    string
	presharedKey string
}

func getConfig(llatWorkspace string) (*LlatConfig, error) {
	log.Println("Read llat config from folder:" + llatWorkspace)
	llatConfigFile := llatWorkspace + "/config"

	readFile, err := os.Open(llatConfigFile)
	if err != nil {
		log.Println(err)
		err = fmt.Errorf("Fail to open file: " + llatConfigFile)
		return nil, err
	}

	configMap := make(map[string]string)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fmt.Println()
		// the config line is in format: "KEY = VAL"
		splitted := strings.Fields(fileScanner.Text())
		if len(splitted) != 3 || splitted[1] != "=" {
			errMsg := "Unrecognized config"
			fmt.Println(errMsg)
			return nil, fmt.Errorf(errMsg)
		}
		k := strings.TrimSpace(splitted[0])
		v := strings.TrimSpace(splitted[2])
		configMap[k] = v
	}
	readFile.Close()

	if !isReleaseMode && verbose {
		log.Println("---- Get llat config ----")
		for k, v := range configMap {
			log.Println(k + " : " + v)
		}
		log.Println("---- End of config ----")
	}

	config := LlatConfig{
		ipAddress:    configMap["Address"],
		privateKey:   configMap["PrivateKey"],
		publicKey:    configMap["PublicKey"],
		presharedKey: configMap["PresharedKey"],
	}

	return &config, nil
}
