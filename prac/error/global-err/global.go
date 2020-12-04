package main

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	initLog sync.Once
)

// Init - sets up the logger initially if run multiple times
// it return a error
func Init() error {
	err := errors.New("already initialized")

	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})

	return err
}

// SetLog sets the log
func SetLog(l *logrus.Logger) {
	log = l
}

// Debug exports the logs withfield connected
// to our global log
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// WithField exports the logs withfield conected to our global log
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}
