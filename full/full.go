package full

import (
	"blbuild/config"
	"path/filepath"
	"blbuild/utils"
)

func ConstructCompileAllFilesCommand(config config.Config) string {
	result := ""
	result += config.Compiler + " "
	if len(config.Std) > 0 {
		result += "-std=" + config.Std + " -c "
	}
	for _, file := range config.Files {
		if len(config.Path) > 0 {
			result += filepath.Join(config.Path, utils.RemoveQuotes(file)) + " "
		} else {
			result += file + " "
		}
	}
	if len(config.Include) > 0 {
		result += config.Include + " "
	}
	return result
}

func ConstructFullBuildCommand(config config.Config) string {
	result := ""
	result += config.Compiler + " "
	if len(config.Std) > 0 {
		result += "-std=" + config.Std + " "
	}
	for _, file := range config.Files {
		if len(config.Path) > 0 {
			result += filepath.Join(config.Path, utils.RemoveQuotes(file)) + " "
		} else {
			result += file + " "
		}
	}
	if len(config.Include) > 0 {
		result += config.Include + " "
	}
	result += "-o " + config.Final
	return result
}

func ConstructBuildCompiledFilesCmd(config config.Config) string {
	result := ""
	result += config.Compiler + " "
	if len(config.Std) > 0 {
		result += "-std=" + config.Std + " "
	}
	result += config.Path + "/*.o "
	if len(config.Include) > 0 {
		result += config.Include + " "
	}
	result += "-o " + config.Final
	return result
}
