package main

import (
	"blbuild/config"
	"blbuild/full"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	configStr, _ := config.GetConfigFile()
	config := config.ConstructConfig(configStr)
	if len(os.Args) > 1 {
		fmt.Println("Has CLArg!")
		//todo: handle single files
	} else {
		command := full.ConstructFullBuildCommand(config)
		fmt.Println("Building full project...")
		cmd := exec.Command("sh", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing build command:", err)
			os.Exit(1)
		}
		fmt.Println("Project built successfully!")
	}
}
