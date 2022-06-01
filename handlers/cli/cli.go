// Package cli implements a colored text handler suitable for command-line interfaces.
package cli

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/caarlos0/log"
	"github.com/charmbracelet/lipgloss"
)

// Default handler outputting to stderr.
var Default = New(os.Stderr)

// Styles mapping.
var Styles = [...]lipgloss.Style{
	log.DebugLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("15")),
	log.InfoLevel:  lipgloss.NewStyle().Foreground(lipgloss.Color("12")),
	log.WarnLevel:  lipgloss.NewStyle().Foreground(lipgloss.Color("11")),
	log.ErrorLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
	log.FatalLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
}

// Strings mapping.
var Strings = [...]string{
	log.DebugLevel: "•",
	log.InfoLevel:  "•",
	log.WarnLevel:  "•",
	log.ErrorLevel: "⨯",
	log.FatalLevel: "⨯",
}

// Handler implementation.
type Handler struct {
	mu     sync.Mutex
	Writer io.Writer

	Padding int
}

const defaultPadding = 2

// New handler.
func New(w io.Writer) *Handler {
	if f, ok := w.(*os.File); ok {
		return &Handler{
			Writer:  f,
			Padding: defaultPadding,
		}
	}

	return &Handler{
		Writer:  w,
		Padding: defaultPadding,
	}
}

// ResetPadding resets the padding to default.
func (h *Handler) ResetPadding() {
	h.Padding = defaultPadding
}

// IncreasePadding increases the padding 1 times.
func (h *Handler) IncreasePadding() {
	h.Padding += defaultPadding
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	style := Styles[e.Level]
	level := Strings[e.Level]
	names := e.Fields.Names()

	h.mu.Lock()
	defer h.mu.Unlock()

	fmt.Fprintf(
		h.Writer,
		"%s %s",
		style.Bold(true).PaddingLeft(h.Padding).Render(level),
		e.Message,
	)
	if len(names) > 0 {
		pad := h.padding(e.Message)
		fmt.Fprint(h.Writer, lipgloss.NewStyle().PaddingLeft(pad).Render(""))
	}
	for _, name := range names {
		fmt.Fprintf(h.Writer, " %s=%v", style.Render(name), e.Fields.Get(name))
	}

	fmt.Fprintln(h.Writer)

	return nil
}

func (h *Handler) padding(m string) int {
	l := h.Padding + 25 - len(m)
	if l >= defaultPadding {
		return l
	}
	return defaultPadding
}
