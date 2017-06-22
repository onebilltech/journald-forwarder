package journald

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"
)

const DefaultSocket = "/var/run/journald-test.sock"

func CollectJournal(c chan<- JournalEntry, errC chan<- error) {
	errC <- collectJournal(c)
}

func collectJournal(c chan<- JournalEntry) error {
	cmd := exec.Command("journalctl", "--output", "json", "--follow")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("journald: could not run journalctl: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("journald: could not run journalctl(2): %v", err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		msg := scanner.Text()
		var entry JournalEntry
		err := json.Unmarshal([]byte(msg), &entry)
		if err != nil {
			return fmt.Errorf("journald: unmarshal error: %v: %v", err, msg)
		}
		c <- (entry)
	}
	return nil
}
