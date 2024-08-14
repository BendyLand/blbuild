package main

import (
	"blbld/config"
	"blbld/full"
	"blbld/single"
	"blbld/utils"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"sync"
)

func main() {
	configStr, _ := config.GetConfigFile()
	config := config.ConstructConfig(configStr)
	if len(os.Args) > 1 {
		move := slices.Contains(os.Args, "mv")
		if slices.Contains(os.Args, "compile") {
			if len(os.Args) == 2 {
				compileAllFiles(config)
			} else {
				compileSingleFile(os.Args[2], config)
			}
		} else if slices.Contains(os.Args, "build") {
			buildCompiledFiles(config, move)
		} else if slices.Contains(os.Args, "print") {
			cmd := full.ConstructFullBuildCommand(config)
			if slices.Contains(os.Args, "debug") {
				cmd += " -g"
			}
			fmt.Println(cmd)
		} else if slices.Contains(os.Args, "update") {
			if len(os.Args) == 2 {
				fmt.Println("Not enough arguments. Please provide a file.")
				return
			}
			compileSingleFile(os.Args[2], config)
			buildCompiledFiles(config, move)
		} else if slices.Contains(os.Args, "help") {
			printHelp()
		} else if slices.Contains(os.Args, "make") {
			makeFiles(config, move)
		} else if slices.Contains(os.Args, "debug") {
			buildAllFiles(config, true, move)
		} else if slices.Contains(os.Args, "mv") {
			buildAllFiles(config, false, move)
		}
	} else {
		buildAllFiles(config, false, false)
	}
}

func makeFiles(config config.Config, move bool) {
	var wg sync.WaitGroup
	for _, file := range config.Files {
		file := utils.RemoveQuotes(file)
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			compileSingleFile(file, config)
		}(file)
	}
	wg.Wait()
	fmt.Println("All files compiled successfully!")
	buildCompiledFiles(config, move)
}

func printHelp() {
	commands := []string{"compile", "compile <file.c/cpp>", "build", "update <file.c/cpp>", "print", "make", "help"}
	descriptions := []string{"Compiles all files using -c.", "Compiles the provided file using -c.", "Links every .o file into an executable.", "Compiles the provided file using -c then builds every .o file in one command.", "Prints the full command to the console, but doesn't run it.", "Concurrently compiles each file independently with -c, then links the .o files together.", "Shows this menu."}
	fmt.Println("USAGE: blbld <opt_command> <opt_file>")
	fmt.Println("Valid commands:")
	for i, command := range commands {
		fmt.Println(command, "-", descriptions[i])
	}
	fmt.Println("Running `blbld` with no arguments compiles everything directly to an executable.")
}

func buildAllFiles(config config.Config, debug bool, move bool) {
	command := full.ConstructFullBuildCommand(config)
	if debug {
		command += " -g"
	}
	fmt.Println(command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if move {
		if len(config.Path) > 0 {
			mvCmd := fmt.Sprintf("mv %s %s", config.Final, config.Path)
			cmd = exec.Command("sh", "-c", mvCmd)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println("Error executing move command:", err)
				os.Exit(1)
			}
		}
	}
	if debug {
		fmt.Println("Debug project built successfully!")
	} else {
		fmt.Println("Project built successfully!")
	}
}

func compileAllFiles(config config.Config) {
	command := full.ConstructCompileAllFilesCommand(config)
	fmt.Println(command)
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
	fmt.Printf("'%s' compiled successfully!\n", name)
}

func buildCompiledFiles(config config.Config, move bool) {
	command := full.ConstructBuildCompiledFilesCmd(config)
	fmt.Println(command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if move {
		if len(config.Path) > 0 {
			mvCmd := fmt.Sprintf("mv %s %s", config.Final, config.Path)
			cmd = exec.Command("sh", "-c", mvCmd)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println("Error executing move command:", err)
				os.Exit(1)
			}
		}
	}
	fmt.Println("Project built successfully!")
}

