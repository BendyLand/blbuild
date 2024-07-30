package main

import (
	"blbuild/config"
	"blbuild/full"
	"blbuild/single"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	configStr, _ := config.GetConfigFile()
	config := config.ConstructConfig(configStr)
	if len(os.Args) > 1 {
		if slices.Contains(os.Args, "compile") {
			if len(os.Args) == 2 {
				fmt.Println("Compiling all files...")
				compileAllFiles(config)
				fmt.Println("Files compiled successfully!")
			} else {
				compileSingleFile(os.Args[2], config)
			}
		}
	} else {
		command := full.ConstructFullBuildCommand(config)
		fmt.Println("Running:", command)
		cmd := exec.Command("sh", "-c", command)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("Error executing build command:", err)
			os.Exit(1)
		}
		fmt.Println("Project built successfully!")
	}
}

func compileAllFiles(config config.Config) {
	command := full.ConstructCompileAllFilesCommand(config)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	mvCmd := "mv *.o " + config.Path
	cmd = exec.Command("sh", "-c", mvCmd)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing move command:", err)
		os.Exit(1)
	}
}

func compileSingleFile(name string, config config.Config) {
	command := single.ConstructSingleFileCompilationCmd(name, config)
	fmt.Println(command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	filename := name[:strings.LastIndex(name, ".")]
	mvCmd := fmt.Sprintf("mv %s.o %s", filename, config.Path)
	cmd = exec.Command("sh", "-c", mvCmd)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing move command:", err)
		os.Exit(1)
	}
}
