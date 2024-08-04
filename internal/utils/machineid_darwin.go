package utils

import "go.uber.org/zap"

func extractId(lines string) string {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `""`)
			}
		}
	}
	return ""
}

func getMachineId() string {
	buffer := &bytes.Buffer{}
	err := runCommand(buffer, os.Stderr, "ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	if err != nil {
		Logger.Error("failed to run command to get machine id", zap.Error(err))
	}
	machineId := extractId(buffer.String())
	if machineId == "" {
		Logger.Error("failed to retrieve machine id", zap.Error(err))
	}
	return trim(machineId)
}

func runCommand(stdout, stderr io.Writer, cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = stdout
	command.Stderr = stderr
	return command.Run()
}

func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}
