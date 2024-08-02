package main

import (
	"blbld/config"
	"blbld/full"
	"blbld/single"
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
		} else if slices.Contains(os.Args, "print") {
			cmd := full.ConstructFullBuildCommand(config)
			fmt.Println(cmd)
		} else if slices.Contains(os.Args, "update") {
			if len(os.Args) == 2 {
				fmt.Println("Not enough arguments. Please provide a file.")
				return
			}
			compileSingleFile(os.Args[2], config)
			buildCompiledFiles(config)
		} else if slices.Contains(os.Args, "help") {
			printHelp()
		}
	} else {
		buildAllFiles(config)
	}
}

func printHelp() {
	commands := []string{"compile", "compile <file.c/cpp>", "build", "update <file.c/cpp>", "print", "help"}
	descriptions := []string{"Compiles all files using -c.", "Compiles the provided file using -c.", "Links every .o file into an executable.", "Compiles the provided file using -c then builds every .o file in one command.", "Prints the full command to the console, but doesn't run it.", "Shows this menu."}
	fmt.Println("USAGE: blbld <opt_command> <opt_file>")
	fmt.Println("Valid commands:")
	for i, command := range commands {
		fmt.Println(command, "-", descriptions[i])
	}
	fmt.Println("Running `blbld` with no arguments compiles everything directly to an executable.")
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
