package log

import (
	"testing"
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
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if l != c.Level {
				t.Fatalf("expected %v, got %v", c.Level, l)
			}
		})
	}

	t.Run("invalid", func(t *testing.T) {
		l, err := ParseLevel("something")
		if err != ErrInvalidLevel {
			t.Fatalf("expected %v, got %v", ErrInvalidLevel, err)
		}
		if l != InvalidLevel {
			t.Fatalf("expected %v, got %v", InvalidLevel, l)
		}
	})
}
