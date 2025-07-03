package log_test

import (
	"testing"

	"github.com/caarlos0/log"
)

func TestFromContext(t *testing.T) {
	ctx := t.Context()

	logger := log.FromContext(ctx)
	if logger != log.Log {
		t.Fatalf("expected %v, got %v", log.Log, logger)
	}

	logs := log.WithField("foo", "bar")
	ctx = log.NewContext(ctx, logs)

	logger = log.FromContext(ctx)
	if logger != logs {
		t.Fatalf("expected %v, got %v", logs, logger)
	}
}
