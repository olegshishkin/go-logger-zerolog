package zerolog

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io"
)

type builder struct {
	console *io.Writer
	file    *io.Writer
}

// NewLogWriterBuilder creates a builder for writer building.
func NewLogWriterBuilder() *builder {
	return &builder{}
}

// WithConsole appends a console writer for the final writer.
func (b *builder) WithConsole(writer io.Writer) *builder {
	b.console = &writer
	return b
}

// WithFile appends a file writer for the final writer.
func (b *builder) WithFile(writer io.Writer) *builder {
	b.file = &writer
	return b
}

// Build builds a final writer.
func (b *builder) Build() (io.Writer, error) {
	if b.console != nil && b.file != nil {
		return zerolog.MultiLevelWriter(*b.console, *b.file), nil
	}
	if b.console != nil {
		return *b.console, nil
	}
	if b.file != nil {
		return *b.file, nil
	}
	return nil, errors.Errorf("no log writers selected")
}
