package journald

import (
	"io"

	"github.com/ugorji/go/codec"
)

var jh codec.JsonHandle

type Decoder struct {
	*codec.Decoder
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{codec.NewDecoder(r, &jh)}
}

func (d *Decoder) Decode(e *Entry) error {
	return d.Decoder.Decode(e)
}

// Decode reads a JSON line into a journald.Entry
func Decode(line string, e *Entry) error {
	dec := &Decoder{codec.NewDecoderBytes([]byte(line), &jh)}
	return dec.Decode(e)
}
