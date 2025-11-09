package services

import (
	"fmt"
)

func killProcess(pid int) (bool, error) {
	if pid <= 0 {
		return false, fmt.Errorf("invalid pid")
	}
	
}
