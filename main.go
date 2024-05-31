package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"strings"
)

func main() {
	file, err := readConfigFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	tokens := parseConfigFile(file)
	commandStr, err := constructCommandString(tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Running: `%s`\n", commandStr)
	cmd := exec.Command("sh", "-c", commandStr)
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compiled successfully!")
}

type Config struct {
	Compiler string
	Path     string
	Files    []string
	Flags    []string
	Options  [][]string
}

func constructCommandString(config Config) (string, error) {
	result := config.Compiler + " "
	files := ""
	for _, file := range config.Files {
		files += config.Path + "/" + file + " "
	}
	result += files
	var pairs []string
	if len(config.Flags) != len(config.Options) {
		err := fmt.Errorf("Wrong number of options provided.")
		return "", err
	}
	for i, flag := range config.Flags {
		pair := flag + " " + strings.Join(config.Options[i], " ")
		pairs = append(pairs, pair)
	}
	flagsOptions := strings.Join(pairs, " ")
	result += flagsOptions
	return result, nil
}

func parseConfigFile(file string) Config {
	var config Config
	toml.Decode(file, &config)
	return config
}

func readConfigFile() (string, error) {
	file, err := os.ReadFile("./blbuild.toml")
	if err != nil {
		myErr := fmt.Errorf("No config file found.\nPlease create `blbuild.toml` in your root directory.\n")
		return "", myErr
	}
	return string(file), nil
}
