package single

import (
	"blbld/config"
	"blbld/utils"
)

func ConstructSingleFileCompilationCmd(name string, config config.Config) string {
	result := ""
	result += config.Compiler + " "
	if len(config.Std) > 0 {
		result += "-std=" + config.Std + " -c "
	} else {
		result += "-c "
	}
	for _, includePath := range config.Include {
		if len(includePath) > 0 {
			temp := utils.RemoveQuotes(includePath)
			result += "-I" + temp + " "
		}
	}
	if len(config.Path) > 0 {
		line := config.Path + "/" + name
		result += line
	} else {
		result += name
	}
	return result
}

