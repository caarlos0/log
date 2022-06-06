package log

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestEntry_WithFields(t *testing.T) {
	is := is.New(t)
	a := NewEntry(nil)
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
	a := NewEntry(nil)
	b := a.WithField("foo", "bar")
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{"foo": "bar"}, b.mergedFields())
}

func TestEntry_WithError(t *testing.T) {
	is := is.New(t)
	a := NewEntry(nil)
	b := a.WithError(fmt.Errorf("boom"))
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{"error": "boom"}, b.mergedFields())
}

func TestEntry_WithError_fields(t *testing.T) {
	is := is.New(t)
	a := NewEntry(nil)
	b := a.WithError(errFields("boom"))
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{
		"error":  "boom",
		"reason": "timeout",
	}, b.mergedFields())
}

func TestEntry_WithError_nil(t *testing.T) {
	is := is.New(t)
	a := NewEntry(nil)
	b := a.WithError(nil)
	is.Equal(Fields{}, a.mergedFields())
	is.Equal(Fields{}, b.mergedFields())
}

type errFields string

func (ef errFields) Error() string {
	return string(ef)
}

func (ef errFields) Fields() Fields {
	return Fields{"reason": "timeout"}
}
