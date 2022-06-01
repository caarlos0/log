package log_test

import (
	"errors"
	"testing"

	"github.com/caarlos0/log"
	"github.com/caarlos0/log/handlers/memory"
	"github.com/matryer/is"
)

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

func TestInfo(t *testing.T) {
	h := memory.New()
	log.SetHandler(h)

	log.Infof("logged in %s", "Tobi")

	e := h.Entries[0]
	is := is.New(t)
	is.Equal(e.Message, "logged in Tobi")
	is.Equal(e.Level, log.InfoLevel)
}

func TestFielder(t *testing.T) {
	h := memory.New()
	log.SetHandler(h)

	pet := &Pet{"Tobi", 3}
	log.WithFields(pet).Info("add pet")

	e := h.Entries[0]
	is := is.New(t)
	is.Equal(log.Fields{"name": "Tobi", "age": 3}, e.Fields)
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

// Trace can be used to simplify logging of start and completion events,
// for example an upload which may fail.
func Example_trace() {
	fn := func() (err error) {
		defer log.Trace("upload").Stop(&err)
		return
	}

	fn()
	return
}
