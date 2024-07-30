package full

import (
	"blbld/config"
	"blbld/utils"
	"path/filepath"
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
			result += utils.RemoveQuotes(file) + " "
		}
	}
	if len(config.Include) > 0 {
		result += "-I " + config.Include + " "
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
			result += utils.RemoveQuotes(file) + " "
		}
	}
	if len(config.Include) > 0 {
		result += "-I " + config.Include + " "
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
	if len(config.Path) > 0 {
		result += config.Path + "/*.o "
	} else {
		result += "*.o "
	}
	if len(config.Include) > 0 {
		result += "-I " + config.Include + " "
	}
	result += "-o " + config.Final
	return result
}
