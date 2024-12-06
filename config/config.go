package config

import (
	"blbld/utils"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"path/filepath"
	"strings"
)

type Config struct {
	Compiler string
	Std      string
	Path     string
	Files    []string
	Include  []string
	Final    string
}


func detectOS() string {
	os := runtime.GOOS
	var result string
	switch os {
	case "windows":
		result = "Windows"
	case "darwin":
		result = "MacOS"
	case "linux":
		result = "Linux"
	case "freebsd":
		result = "FreeBSD"
	case "openbsd":
		result = "OpenBSD"
	case "netbsd":
		result = "NetBSD"
	case "dragonfly":
		result = "DragonflyBSD"
	default:
		result = "Unknown"
	}
	return result
}

func platformTrim(str string, platform string) string {
	var result string
	switch platform {
	case "Windows":
		result = strings.Trim(str, "\r\n")
	default:
		result = strings.Trim(str, "\n")
	}
	return result
}

func getMissingConfigFields(platform string) []string {
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("What compiler would you like to use?")
	compiler, _ := stdin.ReadString('\n')
	compiler = platformTrim(compiler, platform)

	fmt.Println("What std would you like to use? You may leave this blank.")
	std, _ := stdin.ReadString('\n')
	std = platformTrim(std, platform)

	fmt.Println("What is the path from the root directory to where the files are located?")
	fmt.Println("(Keep blank for the root directory.)")
	path, _ := stdin.ReadString('\n')
	path = platformTrim(path, platform)

	fmt.Println("Please enter all of your files separated by spaces, and then a newline.")
	filesStr, _ := stdin.ReadString('\n')
	files := strings.Split(filesStr, " ")
	for i := range len(files) {
		files[i] = platformTrim(files[i], platform)
		files[i] = "\"" + files[i] + "\""
	}
	temp := strings.Trim(strings.Join(files, ", "), "\n")
	filesStr = "[" + temp + "]"

	fmt.Println("If you have additional include paths, please enter them here separated by spaces, otherwise leave it blank.")
	include, _ := stdin.ReadString('\n')
	includePaths := strings.Split(include, " ")
	for i := range len(includePaths) {
		includePaths[i] = platformTrim(includePaths[i], platform)
		includePaths[i] = "\"" + includePaths[i] + "\""
	}
	temp2 := strings.Trim(strings.Join(includePaths, ", "), "\n")
	if temp2 != "\"\"" {
		include = "[" + temp2 + "]"
	} else {
		include = "[]"
	}

	fmt.Println("Please enter the name you would like to use for the final executable.")
	final, _ := stdin.ReadString('\n')
	final = platformTrim(final, platform)

	lines := []string{compiler, std, path, filesStr, include, final}
	return lines
}

func createMissingConfigFile() string {
	var result string
	platform := detectOS()
	lines := getMissingConfigFields(platform)
	for i, line := range lines {
		switch i {
		case 0:
			result += "compiler = \"" + line + "\"\n"
		case 1:
			result += "std = \"" + line + "\"\n"
		case 2:
			result += "path = \"" + line + "\"\n"
		case 3:
			result += "files = " + line + "\n"
		case 4:
			result += "include = " + line + "\n"
		case 5:
			result += "final = \"" + line + "\"\n"
		}
	}
	file, err := os.Create("blbld.toml")
	if err != nil {
		fmt.Println("Error automatically creating build file.\nPlease make your own to avoid re-entering details each time.")
	} else {
		file.WriteString(result)
		file.Close()
	}
	return result
}

func GetConfigFile() (string, error) {
	path, err := os.Getwd()
	configPath := filepath.Join(path, "blbld.toml")
	_, err = os.Stat(configPath)
	if err == nil {
		result, err := os.ReadFile(configPath)
		if err != nil {
			e := fmt.Errorf("Error reading file:%s\n", err)
			return "", e
		}
		return string(result), nil
	}
	fmt.Println("No config file detected. Let's create one:")
	config := createMissingConfigFile()
	return config, nil
}

func ConstructConfig(config string) Config {
	lines := strings.Split(config, "\n")
	var result Config
	for i, line := range lines {
		switch i {
		case 0:
			result.Compiler = utils.ExtractConfigValue(line)
		case 1:
			result.Std = utils.ExtractConfigValue(line)
		case 2:
			result.Path = utils.ExtractConfigValue(line)
		case 3:
			temp := utils.ExtractConfigValue(line)
			temp = strings.Trim(temp, "[]")
			items := strings.Split(temp, ", ")
			for _, item := range items {
				result.Files = append(result.Files, item)
			}
		case 4:
			temp := utils.ExtractConfigValue(line)
			temp = strings.Trim(temp, "[]")
			items := strings.Split(temp, ", ")
			for _, item := range items {
				result.Include = append(result.Include, item)
			}
		case 5:
			result.Final = utils.ExtractConfigValue(line)
		}
	}
	return result
}

func ValidateCompiler(config Config) bool {
	validCompilers := []string{
		"gcc",
		"clang",
		"g++",
		"msvc",
		"icc",
		"scalac",
		"rustc",
		"javac",
		"gc",
		"gccgo",
		"swiftc",
		"fsc",
		"csc",
		"mcs",
		"ghc",
		"kotlinc",
	}
	for _, comp := range validCompilers {
		if config.Compiler == comp {
			return true
		}
	}
	return false
}