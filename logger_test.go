package log_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/caarlos0/log"
)

func TestLogger_printf(t *testing.T) {
	var out bytes.Buffer
	l := log.New(&out)
	l.Infof("logged in %s", "Tobi")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_levels(t *testing.T) {
	var out bytes.Buffer
	l := log.New(&out)

	l.Debug("uploading")
	l.Info("upload complete")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_WithFields(t *testing.T) {
	var out bytes.Buffer
	l := log.New(&out)

	ctx := l.WithFields(log.Fields{"file": "sloth.png"})
	ctx.Debug("uploading")
	ctx.Info("upload complete")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_WithField(t *testing.T) {
	var out bytes.Buffer
	l := log.New(&out)

	ctx := l.WithField("file", "sloth.png").WithField("user", "Tobi")
	ctx.Debug("uploading")
	ctx.Info("upload complete")
	requireEqualOutput(t, out.Bytes())
}

func TestLogger_HandlerFunc(t *testing.T) {
	var out bytes.Buffer
	l := log.New(&out)

	l.Infof("logged in %s", "Tobi")
	requireEqualOutput(t, out.Bytes())
}

func BenchmarkLogger_small(b *testing.B) {
	var out bytes.Buffer
	l := log.New(&out)

	for i := 0; i < b.N; i++ {
		l.Info("login")
	}
}

func BenchmarkLogger_medium(b *testing.B) {
	var out bytes.Buffer
	l := log.New(&out)

	for i := 0; i < b.N; i++ {
		l.WithFields(log.Fields{
			"file": "sloth.png",
			"type": "image/png",
			"size": 1 << 20,
		}).Info("upload")
	}
}

func BenchmarkLogger_large(b *testing.B) {
	var out bytes.Buffer
	l := log.New(&out)

	err := fmt.Errorf("boom")

	for i := 0; i < b.N; i++ {
		l.WithFields(log.Fields{
			"file": "sloth.png",
			"type": "image/png",
			"size": 1 << 20,
		}).
			WithFields(log.Fields{
				"some":     "more",
				"data":     "here",
				"whatever": "blah blah",
				"more":     "stuff",
				"context":  "such useful",
				"much":     "fun",
			}).
			WithError(err).Error("upload failed")
	}
}
