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
	for _, includePath := range config.Include {
		if len(includePath) > 0 {
			temp := utils.RemoveQuotes(includePath)
			result += "-I" + temp + " "
		}
	}
	for _, file := range config.Files {
		if len(config.Path) > 0 {
			result += filepath.Join(config.Path, utils.RemoveQuotes(file)) + " "
		} else {
			result += utils.RemoveQuotes(file) + " "
		}
	}
	return result
}

func ConstructFullBuildCommand(config config.Config) string {
	result := ""
	result += config.Compiler + " "
	if len(config.Std) > 0 {
		result += "-std=" + config.Std + " "
	}
	for _, includePath := range config.Include {
		if len(includePath) > 0 {
			temp := utils.RemoveQuotes(includePath)
			result += "-I" + temp + " "
		}
	}
	for _, file := range config.Files {
		if len(config.Path) > 0 {
			result += filepath.Join(config.Path, utils.RemoveQuotes(file)) + " "
		} else {
			result += utils.RemoveQuotes(file) + " "
		}
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
	for _, includePath := range config.Include {
		if len(includePath) > 0 {
			temp := utils.RemoveQuotes(includePath)
			result += "-I" + temp + " "
		}
	}
	if len(config.Path) > 0 {
		result += config.Path + "/*.o "
	} else {
		result += "*.o "
	}
	result += "-o " + config.Final
	return result
}
