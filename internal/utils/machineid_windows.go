package utils

import "go.uber.org/zap"

func getMachineId() string {
	registryKey, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		Logger.Error("failed to open registry item", zap.Error(err))
	}
	defer registryKey.Close()
	machineId, _, err := registryKey.GetStringValue("MachineGuid")
	if err != nil {
		Logger.Error("failed to get MachineGuid", zap.Error(err))
	}
	return machineId
}
