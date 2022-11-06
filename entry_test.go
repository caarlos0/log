package log

import (
	"fmt"
	"io"
	"testing"

	"github.com/matryer/is"
)

func TestEntry_WithFields(t *testing.T) {
	is := is.New(t)
	a := NewEntry(New(io.Discard))
	is.Equal(a.Fields, nil)

	b := a.WithFields(Fields{"foo": "bar"})
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{"foo": "bar"}, b.mergedFields())

	c := a.WithFields(Fields{"foo": "hello", "bar": "world"})

	e := c.finalize(InfoLevel, "upload")
	is.Equal(e.Message, "upload")
	is.Equal(e.Fields, Fields{"foo": "hello", "bar": "world"})
	is.Equal(e.Level, InfoLevel)
}

func TestEntry_WithField(t *testing.T) {
	is := is.New(t)
	a := NewEntry(New(io.Discard))
	b := a.WithField("foo", "bar")
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{"foo": "bar"}, b.mergedFields())
}

func TestEntry_WithError(t *testing.T) {
	is := is.New(t)
	a := NewEntry(New(io.Discard))
	b := a.WithError(fmt.Errorf("boom"))
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{"error": "boom"}, b.mergedFields())
}

func TestEntry_WithError_fields(t *testing.T) {
	is := is.New(t)
	a := NewEntry(New(io.Discard))
	b := a.WithError(errFields("boom"))
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{
		"error":  "boom",
		"reason": "timeout",
	}, b.mergedFields())
}

func TestEntry_WithError_nil(t *testing.T) {
	is := is.New(t)
	a := NewEntry(New(io.Discard))
	b := a.WithError(nil)
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{}, b.mergedFields())
}

func TestEntry_WithoutPadding(t *testing.T) {
	is := is.New(t)
	log := New(io.Discard)

	a := NewEntry(log)
	is.Equal(defaultPadding, a.Padding)

	log.IncreasePadding()
	b := NewEntry(log)
	is.Equal(defaultPadding+2, b.Padding)

	c := b.WithoutPadding()
	is.Equal(defaultPadding, c.Padding)
}

type errFields string

func (ef errFields) Error() string {
	return string(ef)
}

func (ef errFields) Fields() Fields {
	return Fields{"reason": "timeout"}
}
