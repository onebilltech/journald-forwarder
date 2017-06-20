package journald

import (
	"testing"
	"unicode/utf8"
)

func TestScrub(t *testing.T) {
	var input []byte
	var output string

	input = []byte{104, 101, 108, 108, 111}
	output = Scrub(input)
	if output != "hello" {
		t.Errorf("should be 'hello': %v", output)
	}

	input = []byte{240, 40, 140, 188}
	output = Scrub(input)

	if output != "�(��" {
		t.Errorf("wrong format: %v %v", utf8.ValidString(output), output)
	}
}
