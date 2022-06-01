package log_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/caarlos0/log"
	"github.com/caarlos0/log/handlers/discard"
	"github.com/caarlos0/log/handlers/memory"
	"github.com/matryer/is"
)

func TestLogger_printf(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}

	l.Infof("logged in %s", "Tobi")

	e := h.Entries[0]
	is := is.New(t)
	is.Equal(e.Message, "logged in Tobi")
	is.Equal(e.Level, log.InfoLevel)
}

func TestLogger_levels(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}

	l.Debug("uploading")
	l.Info("upload complete")

	is := is.New(t)
	is.Equal(1, len(h.Entries))

	e := h.Entries[0]
	is.Equal(e.Message, "upload complete")
	is.Equal(e.Level, log.InfoLevel)
}

func TestLogger_WithFields(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}

	ctx := l.WithFields(log.Fields{"file": "sloth.png"})
	ctx.Debug("uploading")
	ctx.Info("upload complete")

	is := is.New(t)
	is.Equal(1, len(h.Entries))

	e := h.Entries[0]
	is.Equal(e.Message, "upload complete")
	is.Equal(e.Level, log.InfoLevel)
	is.Equal(log.Fields{"file": "sloth.png"}, e.Fields)
}

func TestLogger_WithField(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
		Level:   log.InfoLevel,
	}

	ctx := l.WithField("file", "sloth.png").WithField("user", "Tobi")
	ctx.Debug("uploading")
	ctx.Info("upload complete")

	is := is.New(t)
	is.Equal(1, len(h.Entries))

	e := h.Entries[0]
	is.Equal(e.Message, "upload complete")
	is.Equal(e.Level, log.InfoLevel)
	is.Equal(log.Fields{"file": "sloth.png", "user": "Tobi"}, e.Fields)
}

func TestLogger_HandlerFunc(t *testing.T) {
	h := memory.New()
	f := func(e *log.Entry) error {
		return h.HandleLog(e)
	}

	l := &log.Logger{
		Handler: log.HandlerFunc(f),
		Level:   log.InfoLevel,
	}

	l.Infof("logged in %s", "Tobi")

	e := h.Entries[0]
	is := is.New(t)
	is.Equal(e.Message, "logged in Tobi")
	is.Equal(e.Level, log.InfoLevel)
}

func BenchmarkLogger_small(b *testing.B) {
	l := &log.Logger{
		Handler: discard.New(),
		Level:   log.InfoLevel,
	}

	for i := 0; i < b.N; i++ {
		l.Info("login")
	}
}

func BenchmarkLogger_medium(b *testing.B) {
	l := &log.Logger{
		Handler: discard.New(),
		Level:   log.InfoLevel,
	}

	for i := 0; i < b.N; i++ {
		l.WithFields(log.Fields{
			"file": "sloth.png",
			"type": "image/png",
			"size": 1 << 20,
		}).Info("upload")
	}
}

func BenchmarkLogger_large(b *testing.B) {
	l := &log.Logger{
		Handler: discard.New(),
		Level:   log.InfoLevel,
	}

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

func isType(tb testing.TB, a, b any) {
	tb.Helper()
	is.New(tb).Equal(reflect.TypeOf(a), reflect.TypeOf(b))
}
