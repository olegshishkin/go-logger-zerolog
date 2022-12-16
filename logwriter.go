package zerolog

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io"
)

type Builder struct {
	console *io.Writer
	file    *io.Writer
}

// NewLogWriterBuilder creates a builder for writer building.
func NewLogWriterBuilder() *Builder {
	return &Builder{}
}

// WithConsole appends a console writer for the final writer.
func (b *Builder) WithConsole(writer io.Writer) *Builder {
	b.console = &writer

	return b
}

// WithFile appends a file writer for the final writer.
func (b *Builder) WithFile(writer io.Writer) *Builder {
	b.file = &writer

	return b
}

// Build builds a final writer.
func (b *Builder) Build() (io.Writer, error) {
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
