package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/uswitch/journald-forwarder/journald"
	"github.com/uswitch/journald-forwarder/loggly"
)

var token = flag.String("token", "", "Loggly Token")
var logFile = flag.String("logfile", "/var/log/journald-forwarder.log", "Path to log file to write (use \"-\" for stdout)")
var tag = flag.String("tag", "", "What tag to use on Loggly")

func main() {
	flag.Parse()

	if *token == "" {
		log.Fatalf("-token is required.")
	}

	var err error
	var f *os.File

	if *logFile == "-" {
		f = os.Stdout
	} else {
		f, err = os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
	}
	log.SetOutput(f)

	forwardLogs(*token, *tag)
}

func forwardLogs(token, tag string) {
	collector, err := journald.CollectJournal()
	if err != nil {
		log.Fatal(err)
	}

	c := loggly.NewClient(token, tag)
	log.Println("Will send data to", c.URI)

	var (
		entry journald.Entry
		event loggly.Event
	)

	for {
		entry = make(journald.Entry, 10) // make a fresh entry, usually has 10+ fields in it
		err = collector.Decode(&entry)
		if err != nil {
			log.Fatal(err)
		}

		// Add timestamp to the event
		err = convert(entry, &event)
		if err != nil {
			log.Println("conversion error:", err)
		}

		err = c.SendEvent(&event)
		if err != nil {
			log.Println("loggly error:", err)
		}
	}
}

func convert(entry journald.Entry, event *loggly.Event) error {
	// Add timestamp to the event
	ts, err := strconv.ParseInt(string(entry["__REALTIME_TIMESTAMP"]), 10, 64)
	if err != nil {
		return err
	}

	event.Timestamp = time.Unix(0, ts*1000)
	event.Data = make(map[string]string, len(entry))
	for k, v := range entry {
		event.Data[k] = string(v)
	}
	return nil
}
