package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func GetLogger() *Logger {
	if logger == nil {
		logger = newLogger()
	}
	return logger
}

func newLogger() *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate | log.Ltime
	return &Logger{
		debug:   log.New(writer, "DEBUG: ", flags),
		info:    log.New(writer, "INFO: ", flags),
		warning: log.New(writer, "WARNING: ", flags),
		err:     log.New(writer, "ERROR: ", flags),
		writer:  writer,
	}
}

func (l *Logger) Debug(values ...interface{}) {
	l.debug.Println(values...)
}

func (l *Logger) Info(values ...interface{}) {
	l.info.Println(values...)
}

func (l *Logger) Warning(values ...interface{}) {
	l.warning.Println(values...)
}

func (l *Logger) Error(values ...interface{}) {
	l.err.Println(values...)
}

func (l *Logger) Debugf(message string, values ...interface{}) {
	l.debug.Printf(message, values...)
}

func (l *Logger) Infof(message string, values ...interface{}) {
	l.info.Printf(message, values...)
}

func (l *Logger) Warningf(message string, values ...interface{}) {
	l.warning.Printf(message, values...)
}

func (l *Logger) Errorf(message string, values ...interface{}) {
	l.err.Printf(message, values...)
}
