package main

import (
	"flag"
	"log"
	"os"

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

	c := make(chan journald.JournalEntry, 2)
	errC := make(chan error)
	uri := loggly.GenerateUri(*token, *tag)
	go journald.CollectJournal(c, errC)
	go loggly.ProcessJournal(c, errC, uri)

	err = <-errC
	if err != nil {
		log.Fatalln(err)
	}
}
