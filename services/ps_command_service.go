package services

import (
	"fmt"
	"strconv"
	"github.com/Mohit-Pala/idk-whar-to-call-ts"
)

func killProcess(pid int) (string, error) {
	if pid <= 0 {
		return "", fmt.Errorf("invalid pid")
	}

	out, err := thinkofanameforts.ExecuteCommand("kill", strconv.Itoa(pid))
	if err != nil {
		return "", err
	}
	return string(out), nil
}
