package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

func main() {
	file, err := readConfigFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	tokens := parseConfigFile(file)
	fmt.Println(tokens)
	commandStr, err := constructCommandString(tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(commandStr)
}

type Config struct {
	Compiler string
	Files    []string
	Flags    []string
	Options  [][]string
}

func constructCommandString(config Config) (string, error) {
	result := config.Compiler + " "
	files := ""
	for _, file := range config.Files {
		files += file + " "
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
