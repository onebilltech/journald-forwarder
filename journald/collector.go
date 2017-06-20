package journald

import (
	"fmt"
	"os/exec"
)

func CollectJournal() (*Decoder, error) {
	cmd := exec.Command("journalctl", "--output", "json", "--follow")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("journald: could not run journalctl: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("journald: could not run journalctl(2): %v", err)
	}

	return NewDecoder(stdout), nil
}
