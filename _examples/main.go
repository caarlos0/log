package main

import (
	"errors"

	"github.com/caarlos0/log"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.
		WithField("field5", "value5").
		WithField("field2", "value2").
		WithField("field1", "value1").
		WithField("field4", "value4").
		WithFields(log.Fields{
			"FOO": "bar",
		}).
		WithField("field3", "value3").
		Info("AQUI")
	log.WithField("foo", "bar").Debug("debug")
	log.WithField("foo", "bar").Info("info")
	log.WithField("foo", "bar").Warn("warn")
	log.WithFields(log.Fields{
		"multiple": "fields",
		"yes":      true,
	}).Info("a longer line in this particular log")
	log.IncreasePadding()
	log.WithField("foo", "bar").Info("info with increased padding")
	log.IncreasePadding()
	log.WithoutPadding().WithField("foo", "bar").Info("info without padding")
	log.WithField("foo", "bar").Info("info with a more increased padding")
	log.ResetPadding()
	log.WithError(errors.New("some error")).Error("error")
	log.WithError(errors.New("some fatal error")).Fatal("fatal")
}
