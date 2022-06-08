package log_test

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/caarlos0/log"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func init() {
	lipgloss.SetColorProfile(termenv.ANSI256)
}

type Pet struct {
	Name string
	Age  int
}

func (p *Pet) Fields() log.Fields {
	return log.Fields{
		"name": p.Name,
		"age":  p.Age,
	}
}

func TestRootLogOptions(t *testing.T) {
	var out bytes.Buffer
	log.Log = log.New(&out)
	log.SetLevel(log.DebugLevel)
	log.SetLevelFromString("info")
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
	log.ResetPadding()
	pet := &Pet{"Tobi", 3}
	log.WithFields(pet).Info("add pet")
	log.SetHandler(nil)
	requireEqualOutput(t, out.Bytes())
}

// Unstructured logging is supported, but not recommended since it is hard to query.
func Example_unstructured() {
	log.Infof("%s logged in", "Tobi")
}

// Structured logging is supported with fields, and is recommended over the formatted message variants.
func Example_structured() {
	log.WithField("user", "Tobo").Info("logged in")
}

// Errors are passed to WithError(), populating the "error" field.
func Example_errors() {
	err := errors.New("boom")
	log.WithError(err).Error("upload failed")
}

// Multiple fields can be set, via chaining, or WithFields().
func Example_multipleFields() {
	log.WithFields(log.Fields{
		"user": "Tobi",
		"file": "sloth.png",
		"type": "image/png",
	}).Info("upload")
}

var update = flag.Bool("update", false, "update .golden files")

func requireEqualOutput(tb testing.TB, in []byte) {
	tb.Helper()

	bts := useLinuxEOL(in)
	golden := "testdata/" + tb.Name() + ".golden"
	if *update {
		if err := os.MkdirAll(filepath.Dir(golden), 0o755); err != nil {
			tb.Fatal(err)
		}
		if err := os.WriteFile(golden, bts, 0o600); err != nil {
			tb.Fatal(err)
		}
	}

	gbts, err := os.ReadFile(golden)
	if err != nil {
		tb.Fatal(err)
	}
	gbts = useLinuxEOL(gbts)

	if !bytes.Equal(bts, gbts) {
		sg := format(string(gbts))
		so := format(string(bts))
		tb.Fatalf("output do not match:\ngot:\n%s\n\nexpected:\n%s\n\n", so, sg)
	}
}

func useLinuxEOL(bts []byte) []byte {
	return bytes.ReplaceAll(bts, []byte("\r\n"), []byte("\n"))
}

func format(str string) string {
	return strings.NewReplacer(
		"\x1b", "\\x1b",
		"\n", "\\n",
		"\r", "\\r",
		"\t", "\\t",
	).Replace(str)
}
