package example

import (
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/rs/zerolog"
)

// Common creates common sublogger.
func Common(l *zerolog.Logger) *zerolog.Logger {
	result := l.
		With().
		Str(glz.DefaultLogSourceKey, glz.DefaultLogCommonSource).
		Logger()

	return &result
}
