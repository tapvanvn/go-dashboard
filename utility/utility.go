package utility

import "github.com/tapvanvn/goutil"

func GetGeneralConfigPath(rootPath string) string {
	configFile := goutil.GetEnv("CONFIG")

	if configFile == "" {

		configFile = "config.jsonc"
	}
	configPath := rootPath + "/config/" + configFile
	return configPath
}
