package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseLevel(t *testing.T) {
	cases := []struct {
		String string
		Level  Level
		Num    int
	}{
		{"debug", DebugLevel, 0},
		{"info", InfoLevel, 1},
		{"warn", WarnLevel, 2},
		{"warning", WarnLevel, 3},
		{"error", ErrorLevel, 4},
		{"fatal", FatalLevel, 5},
	}

	for _, c := range cases {
		t.Run(c.String, func(t *testing.T) {
			l, err := ParseLevel(c.String)
			require.NoError(t, err) // no parse err
			require.Equal(t, c.Level, l)
		})
	}

	t.Run("invalid", func(t *testing.T) {
		l, err := ParseLevel("something")
		require.Equal(t, ErrInvalidLevel, err)
		require.Equal(t, InvalidLevel, l)
	})
}
