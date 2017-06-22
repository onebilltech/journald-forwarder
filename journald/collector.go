package journald

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

// CollectJournal is the best
func CollectJournal(c chan<- Entry, errC chan<- error) {
	errC <- collectJournal(c)
}

func collectJournal(c chan<- Entry) error {
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
		var entry Entry
		err := Decode(msg, &entry)
		if err != nil {
			log.Println("journald: unmarshal error", err, msg)
		} else {
			c <- (entry)
		}
	}
	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("journald: scanner error: %v", err)
	}
	return nil
}
