package log_test

import (
	"context"
	"testing"

	"github.com/caarlos0/log"
	"github.com/stretchr/testify/require"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()

	logger := log.FromContext(ctx)
	require.Equal(t, log.Log, logger)

	logs := log.WithField("foo", "bar")
	ctx = log.NewContext(ctx, logs)

	logger = log.FromContext(ctx)
	require.Equal(t, logs, logger)
}
