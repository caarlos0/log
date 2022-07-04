package main

import (
	"errors"
	"os"

	"github.com/caarlos0/log"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func main() {
	if os.Getenv("CI") != "" {
		lipgloss.SetColorProfile(termenv.TrueColor)
		lipgloss.SetOutput(termenv.NewOutput(os.Stderr))
		log.Log = log.New(os.Stderr)
	}

	log.SetLevel(log.DebugLevel)
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
	log.WithField("foo", "bar").Info("info with a more increased padding")
	log.ResetPadding()
	log.WithError(errors.New("some error")).Error("error")
	log.WithError(errors.New("some fatal error")).Fatal("fatal")
}
