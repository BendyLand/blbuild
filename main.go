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
				compileAllFiles(config)
			} else {
				compileSingleFile(os.Args[2], config)
			}
		} else if slices.Contains(os.Args, "build") {
			buildCompiledFiles(config)
		}
	} else {
		buildAllFiles(config)
	}
}

func buildAllFiles(config config.Config) {
	command := full.ConstructFullBuildCommand(config)
	fmt.Println("Running:", command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if len(config.Path) > 0 {
		mvCmd := fmt.Sprintf("mv %s %s", config.Final, config.Path)
		cmd = exec.Command("sh", "-c", mvCmd)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing move command:", err)
			os.Exit(1)
		}
	}
	fmt.Println("Project built successfully!")
}

func compileAllFiles(config config.Config) {
	command := full.ConstructCompileAllFilesCommand(config)
	fmt.Println("Running:", command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if len(config.Path) > 0 {
		mvCmd := "mv *.o " + config.Path
		cmd = exec.Command("sh", "-c", mvCmd)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing move command:", err)
			os.Exit(1)
		}
	}
	fmt.Println("Files compiled successfully!")
}

func compileSingleFile(name string, config config.Config) {
	command := single.ConstructSingleFileCompilationCmd(name, config)
	fmt.Println(command)
	fmt.Printf("Compiling '%s'...\n", name)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if len(config.Path) > 0 {
		filename := name[:strings.LastIndex(name, ".")]
		mvCmd := fmt.Sprintf("mv %s.o %s", filename, config.Path)
		cmd = exec.Command("sh", "-c", mvCmd)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing move command:", err)
			os.Exit(1)
		}
	}
	fmt.Println("File compiled successfully!")
}

func buildCompiledFiles(config config.Config) {
	command := full.ConstructBuildCompiledFilesCmd(config)
	fmt.Println("Running:", command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if len(config.Path) > 0 {
		mvCmd := fmt.Sprintf("mv %s %s", config.Final, config.Path)
		cmd = exec.Command("sh", "-c", mvCmd)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing move command:", err)
			os.Exit(1)
		}
	}
	fmt.Println("Project built successfully!")
}
