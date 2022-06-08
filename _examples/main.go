package main

import (
	"errors"

	"github.com/caarlos0/log"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.WithField("foo", "bar").Debug("debug")
	log.WithField("foo", "bar").Info("info")
	log.WithField("foo", "bar").Warn("warn")
	log.WithError(errors.New("some error")).Error("error")
	log.WithError(errors.New("some fatal error")).Fatal("fatal")
}
