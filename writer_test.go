package log_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/caarlos0/log"
)

func TestWriter_singleLine(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	w := l.Writer()
	fmt.Fprint(w, "hello\n")
	requireEqualOutput(t, buf.Bytes())
}

func TestWriter_multiLine(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	w := l.Writer()
	fmt.Fprint(w, "a\nb\nc\n")
	requireEqualOutput(t, buf.Bytes())
}

func TestWriter_noTrailingNewline(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	w := l.Writer()
	fmt.Fprint(w, "hello")
	requireEqualOutput(t, buf.Bytes())
}

func TestWriter_consecutiveWrites(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	w := l.Writer()
	fmt.Fprint(w, "first\n")
	fmt.Fprint(w, "second\n")
	requireEqualOutput(t, buf.Bytes())
}

func TestWriter_consecutiveWritesNoNewline(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	w := l.Writer()
	fmt.Fprint(w, "first")
	fmt.Fprint(w, "second\n")
	requireEqualOutput(t, buf.Bytes())
}

func TestWriter_withPadding(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	l.IncreasePadding()
	w := l.Writer()
	fmt.Fprint(w, "indented\n")
	requireEqualOutput(t, buf.Bytes())
}

func TestWriter_concurrent(t *testing.T) {
	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("COLORTERM", "truecolor")
	var buf bytes.Buffer
	l := log.New(&buf)
	w := l.Writer()

	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Fprint(w, "line\n")
		}()
	}
	wg.Wait()

	if buf.Len() == 0 {
		t.Fatal("expected output, got none")
	}
}
