package example

import (
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/rs/zerolog"
)

// Storage creates sublogger for storage actions logging.
func Storage(l *zerolog.Logger) *zerolog.Logger {
	result := l.
		With().
		Str(glz.DefaultLogSourceKey, glz.Colorize(glz.DefaultLogSourceColor, glz.DefaultLogStorageSource)).
		Logger()
	return &result
}
