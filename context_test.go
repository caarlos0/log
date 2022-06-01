package log_test

import (
	"context"
	"testing"

	"github.com/caarlos0/log"
	"github.com/matryer/is"
)

func TestFromContext(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	logger := log.FromContext(ctx)
	is.Equal(log.Log, logger)

	logs := log.WithField("foo", "bar")
	ctx = log.NewContext(ctx, logs)

	logger = log.FromContext(ctx)
	is.Equal(logs, logger)
}
