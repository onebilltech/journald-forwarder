package loggly

import (
	"encoding/json"
	"time"
)

// Event represents a loggly log entry
type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Data      map[string]string
}

// MarshalJSON implements the json.Marshaler interface. Note that e.Data gets mutated during that operation.
func (e *Event) MarshalJSON() ([]byte, error) {
	e.Data["timestamp"] = e.Timestamp.Format(time.RFC3339Nano)
	return json.Marshal(e.Data)
}
