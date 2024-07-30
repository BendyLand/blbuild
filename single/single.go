package single

import (
	"blbld/config"
)

func ConstructSingleFileCompilationCmd(name string, config config.Config) string {
	result := ""
	result += config.Compiler + " "
	if len(config.Std) > 0 {
		result += "-std=" + config.Std + " -c "
	}
	if len(config.Include) > 0 {
		result += "-I " + config.Include + " "
	}
	if len(config.Path) > 0 {
		line := config.Path + "/" + name
		result += line
	} else {
		result += name
	}
	return result
}

