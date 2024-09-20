package commands

import (
	"blbld/config"
	"blbld/full"
	"blbld/single"
	"blbld/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"slices"
)

func ExecuteSpecialCommand(config config.Config) {
		move := slices.Contains(os.Args, "mv")
		debug := slices.Contains(os.Args, "debug")
		if slices.Contains(os.Args, "compile") {
			if len(os.Args) == 2 {
				compileAllFiles(config, debug)
			} else {
				compileSingleFile(os.Args[2], config, debug)
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
			compileSingleFile(os.Args[2], config, debug)
			fmt.Println()
			buildCompiledFiles(config, move)
		} else if slices.Contains(os.Args, "help") {
			printHelp()
		} else if slices.Contains(os.Args, "make") {
			makeFiles(config, move, debug)
		} else {
			BuildAllFiles(config, debug, move)
		}
}

func makeFiles(config config.Config, move bool, debug bool) {
	var wg sync.WaitGroup
	for _, file := range config.Files {
		file := utils.RemoveQuotes(file)
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			compileSingleFile(file, config, debug)
		}(file)
	}
	wg.Wait()
	fmt.Println("\nAll files compiled successfully!")
	fmt.Println()
	buildCompiledFiles(config, move)
}

func printHelp() {
	commands := []string{"compile", "compile <file.c/cpp>", "build", "update <file.c/cpp>", "print", "make", "debug", "mv", "help"}
	descriptions := []string{"Compiles all files using -c.", "Compiles the provided file using -c.", "Links every .o file into an executable.", "Compiles the provided file using -c then builds every .o file in one command.", "Prints the full command to the console, but doesn't run it.", "Concurrently compiles each file independently with -c, then links the .o files together.", "Include this command to compile the files with -g. May be used with other commands.", "Include this command to move the resulting binary to the source directory. May be used with other commands.", "Shows this menu."}
	fmt.Println("USAGE: blbld <opt_command> <opt_file>")
	fmt.Println("Valid commands:")
	for i, command := range commands {
		fmt.Println(command, "-", descriptions[i])
	}
	fmt.Println("Running `blbld` with no arguments compiles everything directly to an executable. You may include modifier commands like `debug` and `mv`.")
}

func BuildAllFiles(configFile config.Config, debug bool, move bool) {
	if !config.ValidateCompiler(configFile) {
		fmt.Println("Error: Invalid compiler.")
		os.Exit(1)
	}
	command := full.ConstructFullBuildCommand(configFile)
	command = utils.Sanitize(command)
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
		if len(configFile.Path) > 0 {
			mvCmd := fmt.Sprintf("mv %s %s", configFile.Final, configFile.Path)
			cmd = exec.Command("sh", "-c", mvCmd)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println("Error executing move command:", err)
				os.Exit(1)
			}
		}
	}
	if debug {
		fmt.Println("\nDebug project built successfully!")
	} else {
		fmt.Println("\nProject built successfully!")
	}
}

func compileAllFiles(configFile config.Config, debug bool) {
	if !config.ValidateCompiler(configFile) {
		fmt.Println("Error: Invalid compiler.")
		os.Exit(1)
	}
	command := full.ConstructCompileAllFilesCommand(configFile)
	if debug {
		command += " -g"
	}
	command = utils.Sanitize(command)
	fmt.Println(command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if len(configFile.Path) > 0 {
		mvCmd := "mv *.o " + configFile.Path
		cmd = exec.Command("sh", "-c", mvCmd)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing move command:", err)
			os.Exit(1)
		}
	}
	fmt.Println("\nFiles compiled successfully!")
}

func compileSingleFile(name string, configFile config.Config, debug bool) {
	if !config.ValidateCompiler(configFile) {
		fmt.Println("Error: Invalid compiler.")
		os.Exit(1)
	}
	command := single.ConstructSingleFileCompilationCmd(name, configFile)
	if debug {
		command += " -g"
	}
	command = utils.Sanitize(command)
	fmt.Println(command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if len(configFile.Path) > 0 {
		filename := name[:strings.LastIndex(name, ".")]
		mvCmd := fmt.Sprintf("mv %s.o %s", filename, configFile.Path)
		cmd = exec.Command("sh", "-c", mvCmd)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing move command:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("\n'%s' compiled successfully!\n", name)
}

func buildCompiledFiles(configFile config.Config, move bool) {
	if !config.ValidateCompiler(configFile) {
		fmt.Println("Error: Invalid compiler.")
		os.Exit(1)
	}
	command := full.ConstructBuildCompiledFilesCmd(configFile)
	command = utils.Sanitize(command)
	fmt.Println(command)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing build command:", err)
		os.Exit(1)
	}
	if move {
		if len(configFile.Path) > 0 {
			mvCmd := fmt.Sprintf("mv %s %s", configFile.Final, configFile.Path)
			cmd = exec.Command("sh", "-c", mvCmd)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println("Error executing move command:", err)
				os.Exit(1)
			}
		}
	}
	fmt.Println("\nProject built successfully!")
}
