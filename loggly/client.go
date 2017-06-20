package loggly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is used to talk to Loggly
type Client struct {
	URI string
	*http.Client
}

// NewClient creates a new instance of Client
func NewClient(token string, tag string) *Client {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &Client{sendURI(token, tag), c}
}

type response struct {
	Response string `json:"response"`
}

// SendEvent sends the given event to be logged by loggly
func (s *Client) SendEvent(e *Event) error {
	entry, err := e.MarshalJSON()
	if err != nil {
		return err
	}

	resp, err := s.Post(s.URI, "application/x-www-form-urlencoded", ioutil.NopCloser(bytes.NewReader(entry)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}
	if r.Response != "ok" {
		return fmt.Errorf("invalid response, was: %s", r.Response)
	}

	return nil
}

func sendURI(token string, tag string) string {
	uri := ""
	switch tag {
	case "":
		uri = fmt.Sprintf("https://logs-01.loggly.com/inputs/%s", token)
	default:
		uri = fmt.Sprintf("https://logs-01.loggly.com/inputs/%s/tag/%s/", token, tag)
	}

	return uri
}
