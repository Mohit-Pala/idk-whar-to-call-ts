package thinkofanameforts

// remae ts when you think of a name lmao
import (
	"fmt"
	"os/exec"
)

// genertic method, to exec temrinal commands
func executeCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	// out,err := cmd.Output()
	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf("executeCommand threw err, err: %v, output: %s", err, string(out))
	}
	return out, nil
}
