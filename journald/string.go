package journald

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func init() {
	t := reflect.TypeOf(String(""))
	jh.SetInterfaceExt(t, 1, StringExt{})
}

// String is a journald string. journald log entries are represented as a map[string]string, except if the
// value contains invalid UTF-8 characters or ANSI escape codes, in that case the value is a JSON array of ints,
// which maps back to an array of bytes.
//
// Instead or keeping that array, pack is back into a string and scrub all the invalid chars.
type String string

// StringExt implements the codec.InterfaceExt for the journald.String
type StringExt struct{}

// ConvertExt converts a value into a simpler interface for easy encoding e.g. convert time.Time to int64.
func (x StringExt) ConvertExt(v interface{}) interface{} {
	str := v.(*String)
	return string(*str)
}

// UpdateExt updates a value from a simpler interface for easy decoding e.g. convert int64 to time.Time.
func (x StringExt) UpdateExt(dst interface{}, src interface{}) {
	v := dst.(*String)
	switch v2 := src.(type) {
	case string:
		str := String(v2)
		*v = str
	case []interface{}:
		bytes := make([]byte, len(v2))
		for i, j := range v2 {
			bytes[i] = byte(j.(uint64))
		}
		str := String(Scrub(bytes))
		*v = str
	default:
		panic(fmt.Sprintf("unsupported format for journald.String conversion: %T", v2))
	}
}

// Scrub replaces all invalid UTF-8 characters in a string with a replacement char.
func Scrub(bytes []byte) string {
	s := string(bytes)
	if utf8.ValidString(s) {
		return s
	}

	// scrub, this replaces all invalid chars with replacement chars
	rs := make([]rune, 0, len(s))
	for _, r := range s {
		rs = append(rs, r)
	}
	s = string(rs)

	return s
}
