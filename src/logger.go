package main

import (
	"log"
	"os"
)

var LOG *Logger

type Logger struct {
	logfile *os.File
	logger  *log.Logger
	enabled bool
}

func NewLogger(enabled bool) *Logger {
	logfile, err := os.OpenFile(LOGFILENAME, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	return &Logger{logfile: logfile, logger: logger, enabled: enabled}
}

func (l *Logger) Close() {
	l.logger = nil
	l.enabled = false
	l.logfile.Close()
}

func (l *Logger) Printf(format string, v ...any) {
	if !l.enabled {
		return
	}
	l.logger.Printf(format, v...)
}

func InitLogger(enabled bool) {
	LOG = NewLogger(enabled)
}

func FiniLogger() {
	LOG.Close()
}
