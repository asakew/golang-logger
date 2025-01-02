package main

import (
	log "github.com/sirupsen/logrus"
)

// logRus is well suited for structured logging in JSON which — as JSON is a well-defined standard — makes it easy for external services to parse your logs and also makes the addition of context to a log message relatively straightforward through the use of fields, as shown below:

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")
}

// {"bar":"bar","foo":"foo","level":"info","msg":"Something happened","time":"2025-01-02T10:46:16+10:00"}
