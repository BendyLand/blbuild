package main

import (
	"blbld/config"
	"os"
	"blbld/commands"
)

func main() {
	configStr, _ := config.GetConfigFile()
	config := config.ConstructConfig(configStr)
	if len(os.Args) > 1 {
		commands.ExecuteSpecialCommand(config)
	} else {
		commands.BuildAllFiles(config, false, false)
	}
}
