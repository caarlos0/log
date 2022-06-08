package log

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/charmbracelet/lipgloss"
)

// Styles mapping.
var Styles = [...]lipgloss.Style{
	DebugLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("15")),
	InfoLevel:  lipgloss.NewStyle().Foreground(lipgloss.Color("12")),
	WarnLevel:  lipgloss.NewStyle().Foreground(lipgloss.Color("11")),
	ErrorLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
	FatalLevel: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
}

// Strings mapping.
var Strings = [...]string{
	DebugLevel: "•",
	InfoLevel:  "•",
	WarnLevel:  "•",
	ErrorLevel: "⨯",
	FatalLevel: "⨯",
}

// CLI implementation.
type CLI struct {
	mu     sync.Mutex
	Writer io.Writer

	Padding int
}

const defaultPadding = 2

func newCLI(w io.Writer) *CLI {
	if f, ok := w.(*os.File); ok {
		return &CLI{
			Writer:  f,
			Padding: defaultPadding,
		}
	}

	return &CLI{
		Writer:  w,
		Padding: defaultPadding,
	}
}

// ResetPadding resets the padding to default.
func (h *CLI) ResetPadding() {
	h.Padding = defaultPadding
}

// IncreasePadding increases the padding 1 times.
func (h *CLI) IncreasePadding() {
	h.Padding += defaultPadding
}

// DecreasePadding decreases the padding 1 times.
func (h *CLI) DecreasePadding() {
	h.Padding -= defaultPadding
}

// HandleLog implements Handler.
func (h *CLI) HandleLog(e *Entry) error {
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

func (h *CLI) padding(m string) int {
	l := h.Padding + 25 - len(m)
	if l >= defaultPadding {
		return l
	}
	return defaultPadding
}
