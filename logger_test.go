package log_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/caarlos0/log"
)

func TestLoggerOrdering(t *testing.T) {
	var l sync.Mutex
	var outs [][]byte
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var out bytes.Buffer
			log := log.New(&out)
			log.WithError(fmt.Errorf("here")).Info("a")
			log.Debug("debug")
			log.Debugf("warn %d", 1)
			log.Info("info")
			log.Infof("warn %d", 1)
			log.Warn("warn")
			log.Warnf("warn %d", 1)
			log.Error("error")
			log.Errorf("warn %d", 1)
			log.WithField("foo", "bar").Info("foo")
			log.IncreasePadding()
			log.Info("increased")
			log.WithoutPadding().WithField("foo", "bar").Info("without padding")
			log.Info("increased")
			log.ResetPadding()
			l.Lock()
			outs = append(outs, out.Bytes())
			l.Unlock()
		}()
	}
	wg.Wait()
	for i := 0; i < len(outs)-1; i++ {
		s1 := string(outs[i])
		s2 := string(outs[i+1])
		if s1 != s2 {
			t.Errorf("at least one of the outputs is different:\n%q\nvs\n%q\n", s1, s2)
		}
	}
	requireEqualOutput(t, outs[0])
}

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
		l.WithField("file", "sloth.png").
			WithField("type", "image/png").
			WithField("size", 1<<20).
			Info("upload")
	}
}

func BenchmarkLogger_large(b *testing.B) {
	var out bytes.Buffer
	l := log.New(&out)

	err := fmt.Errorf("boom")

	for i := 0; i < b.N; i++ {
		l.WithField("file", "sloth.png").
			WithField("type", "image/png").
			WithField("size", 1<<20).
			WithField("some", "more").
			WithField("data", "here").
			WithField("whatever", "blah blah").
			WithField("more", "stuff").
			WithField("context", "such useful").
			WithField("much", "fun").
			WithError(err).Error("upload failed")
	}
}
