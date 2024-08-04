package utils

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

const (
	dbusPath    = "/var/lib/dbus/machine-id"
	dbusPathEtc = "/var/machine-id"
)

func getMachineId() string {
	machineId, err := os.ReadFile(dbusPath)
	if err != nil {
		machineId, err = os.ReadFile(dbusPathEtc)
	}
	if err != nil {
		Logger.Error("failed to retrieve machine id", zap.Error(err))
	}
	return trim(string(machineId))
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}
