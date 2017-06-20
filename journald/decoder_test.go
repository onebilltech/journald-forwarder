package journald

import (
	// "io"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	var sample string
	var e Entry
	var err error

	e = make(Entry)
	sample = `{"MESSAGE": [104, 101, 108, 108, 111], "__CURSOR": "woot"}`
	err = Decode(sample, &e)
	if err != nil {
		t.Error(err)
	}
	if "hello" != e["MESSAGE"] {
		t.Errorf("should be hello: %v", e)
	}
	fmt.Println(e)

	e = make(Entry)
	sample = `{"MESSAGE": "test"}`
	err = Decode(sample, &e)
	if err != nil {
		t.Error(err)
	}
	if "test" != e["MESSAGE"] {
		t.Errorf("should be test: %v", e)
	}
	fmt.Println(e)
}
