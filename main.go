package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/exec"
	"path/filepath"
)

// todo: refactor to use positional arguments and groups, rather than "official" names
func main() {
	configPath, err := findConfigFile()
	if err != nil {
		fmt.Println("Error finding config file:", err)
		return
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	tokens := parseConfigFile(string(file))
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
	Extras   []string
}

func constructCommandString(config Config) (string, error) {
	result := config.Compiler + " "
	files := ""
	for _, file := range config.Files {
		fullPath := filepath.Join(config.Path, file)
		files += fullPath + " "
	}
	result += files
	for _, extra := range config.Extras {
		result += extra
	}
	return result, nil
}

func parseConfigFile(file string) Config {
	var config Config
	toml.Decode(file, &config)
	return config
}

func findConfigFile() (string, error) {
	// Look for blbuild.toml in the current directory and its parents
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		configPath := filepath.Join(dir, "blbuild.toml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return "", fmt.Errorf("No config file found. Please create `blbuild.toml` in your project directory or its parent directories.")
}
