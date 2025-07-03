package log

import (
	"fmt"
	"io"
	"testing"
)

func TestEntry_WithField(t *testing.T) {
	a := NewEntry(New(io.Discard))
	b := a.WithField("foo", "bar")
	if len(a.Fields.Keys()) != 0 {
		t.Fatalf("expected empty fields, got %v", a.Fields.Keys())
	}
	keys := b.Fields.Keys()
	if len(keys) != 1 || keys[0] != "foo" {
		t.Fatalf("expected [foo], got %v", keys)
	}
}

func TestEntry_WithError(t *testing.T) {
	a := NewEntry(New(io.Discard))
	b := a.WithError(fmt.Errorf("boom"))
	if len(a.Fields.Keys()) != 0 {
		t.Fatalf("expected empty fields, got %v", a.Fields.Keys())
	}
	keys := b.Fields.Keys()
	if len(keys) != 1 || keys[0] != "error" {
		t.Fatalf("expected [error], got %v", keys)
	}
}

func TestEntry_WithError_nil(t *testing.T) {
	a := NewEntry(New(io.Discard))
	b := a.WithError(nil)
	if len(a.Fields.Keys()) != 0 {
		t.Fatalf("expected empty fields, got %v", a.Fields.Keys())
	}
	if len(b.Fields.Keys()) != 0 {
		t.Fatalf("expected empty fields, got %v", b.Fields.Keys())
	}
}

func TestEntry_WithoutPadding(t *testing.T) {
	log := New(io.Discard)

	a := NewEntry(log)
	if a.Padding != defaultPadding {
		t.Fatalf("expected %d, got %d", defaultPadding, a.Padding)
	}

	log.IncreasePadding()
	b := NewEntry(log)
	if b.Padding != defaultPadding+2 {
		t.Fatalf("expected %d, got %d", defaultPadding+2, b.Padding)
	}

	c := b.WithoutPadding()
	if c.Padding != defaultPadding {
		t.Fatalf("expected %d, got %d", defaultPadding, c.Padding)
	}
}
