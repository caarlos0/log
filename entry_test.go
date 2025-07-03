package log

import (
	"fmt"
	"io"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntry_WithField(t *testing.T) {
	a := NewEntry(New(io.Discard))
	b := a.WithField("foo", "bar")
	require.Empty(t, a.Fields)
	require.Equal(t, []string{"foo"}, slices.Collect(maps.Keys(b.Fields)))
}

func TestEntry_WithError(t *testing.T) {
	a := NewEntry(New(io.Discard))
	b := a.WithError(fmt.Errorf("boom"))
	require.Empty(t, a.Fields)
	require.Equal(t, []string{"error"}, slices.Collect(maps.Keys(b.Fields)))
}

func TestEntry_WithError_nil(t *testing.T) {
	a := NewEntry(New(io.Discard))
	b := a.WithError(nil)
	require.Empty(t, slices.Collect(maps.Keys(a.Fields)))
	require.Empty(t, slices.Collect(maps.Keys(b.Fields)))
}

func TestEntry_WithoutPadding(t *testing.T) {
	log := New(io.Discard)

	a := NewEntry(log)
	require.Equal(t, defaultPadding, a.Padding)

	log.IncreasePadding()
	b := NewEntry(log)
	require.Equal(t, defaultPadding+2, b.Padding)

	c := b.WithoutPadding()
	require.Equal(t, defaultPadding, c.Padding)
}
