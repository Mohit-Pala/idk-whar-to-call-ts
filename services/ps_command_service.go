package services

import (
	"fmt"
	"github.com/Mohit-Pala/idk-whar-to-call-ts"
	"strconv"
)

type PsCommandService struct{}

func (s *PsCommandService) KillProcess(pid int) (string, error) {
	if pid <= 0 {
		return "", fmt.Errorf("invalid pid")
	}

	out, err := thinkofanameforts.ExecuteCommand("kill", strconv.Itoa(pid))
	if err != nil {
		return "", err
	}
	return string(out), nil
}
