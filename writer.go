package log

import (
	"bytes"
	"io"
	"sync"
)

type indentedLogWriter struct {
	indent     string
	writer     io.Writer
	needIndent bool
	mu         sync.Mutex
}

// Write implements [io.Writer].
func (w *indentedLogWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	n := len(p)
	endsNL := p[n-1] == '\n'
	indent := []byte(w.indent)
	nl := append([]byte{'\n'}, indent...)

	out := bytes.ReplaceAll(bytes.TrimSuffix(p, []byte{'\n'}), []byte{'\n'}, nl)
	if w.needIndent {
		out = append(indent, out...)
	}
	if endsNL {
		out = append(out, '\n')
	}
	w.needIndent = endsNL

	_, err := w.writer.Write(out)
	return n, err
}
