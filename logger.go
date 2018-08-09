package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger is a default implementation of the Logger interface.
type Logger struct {
	log *log.Logger

	debug     bool
	calldepth int
}

// New returns a new logger with the given writer as the output
func New(w io.Writer) *Logger {
	out := w
	if out == nil {
		out = os.Stderr
	}

	return &Logger{
		calldepth: 2,
		log:       log.New(out, "", log.LstdFlags|log.Lmicroseconds),
	}
}

// SetCallDepth sets the call depth of the logger
func (l *Logger) SetCallDepth(d int) {
	l.calldepth = d
}

// EnableDebug enables debug messages to print.
func (l *Logger) EnableDebug(v bool) {
	if v {
		l.log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	}
	l.debug = v
}

// Debug implements the Logger interface.
func (l *Logger) Debug(v ...interface{}) {
	if l.debug {
		l.log.Output(l.calldepth, header("DBG", fmt.Sprint(v...)))
	}
}

// Debugf implements the Logger interface.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.log.Output(l.calldepth, header("DBG", fmt.Sprintf(format, v...)))
	}
}

// Info implements the Logger interface.
func (l *Logger) Info(v ...interface{}) {
	l.log.Output(l.calldepth, header("INF", fmt.Sprint(v...)))
}

// Infof implements the Logger interface.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.log.Output(l.calldepth, header("INF", fmt.Sprintf(format, v...)))
}

// Error implements the Logger interface.
func (l *Logger) Error(v ...interface{}) {
	l.log.Output(l.calldepth, header("ERR", fmt.Sprint(v...)))
}

// Errorf implements the Logger interface.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log.Output(l.calldepth, header("ERR", fmt.Sprintf(format, v...)))
}

// Fatal implements the Logger interface.
func (l *Logger) Fatal(v ...interface{}) {
	l.log.Output(l.calldepth, header("FTL", fmt.Sprint(v...)))
	os.Exit(1)
}

// Fatalf implements the Logger interface.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log.Output(l.calldepth, header("FTL", fmt.Sprintf(format, v...)))
	os.Exit(1)
}

// Panic implements the Logger interface.
func (l *Logger) Panic(v ...interface{}) {
	l.log.Panic(v)
}

// Panicf implements the Logger interface.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.log.Panicf(format, v...)
}

func header(lvl, msg string) string {
	return lvl + ": " + msg
}
