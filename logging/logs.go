package logging

import (
	"bytes"
	"log"
)

type MatchLogger struct {
	name string
	*log.Logger
}

var myLogger *MatchLogger = nil

//GetLogger will create a new Logger or return the one we currently have
//this allows us to only make one instance of said logger.
func GetLogger() *MatchLogger {
	if myLogger == nil {
		var buf bytes.Buffer
		myLogger = &MatchLogger{Logger: log.New(&buf, "logger:", log.Lshortfile)}
	}
	return myLogger
}
