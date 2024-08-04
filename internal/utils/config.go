package utils

import (
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

const (
	credentialFilePath = ".artifactsmmo/credentials/creds.json"
)

func ConfigPath(credFile bool) string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		Logger.Error("unable to retrieve users home directory", zap.Error(err))
	}
	if credFile {
		configFile := filepath.Join(homedir, credentialFilePath)
		return configFile
	}
	return ""
}
