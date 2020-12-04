package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	initLog sync.Once
)

func Init() error {

}

func SetLog(l *logrus.Logger) {

}

func WithField(key string, value interface{}) *logrus.Entry {

}

func Debug(args ...interface{}) {

}
