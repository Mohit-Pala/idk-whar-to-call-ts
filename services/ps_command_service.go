package services

import (
	"fmt"
	"strconv"
	thinkofanameforts "github.com/Mohit-Pala/idk-whar-to-call-ts"
)

func killProcess(pid int) (bool, error) {
	if pid <= 0 {
		return false, fmt.Errorf("invalid pid")
	}

	_, err := thinkofanameforts.ExecuteCommand("kill", strconv.Itoa(pid))
	if err != nil {
		return false, err
	}
	return true, nil
}
