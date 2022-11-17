package example

import (
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/rs/zerolog"
)

// Web creates sublogger for web-server actions logging.
func Web(l *zerolog.Logger) *zerolog.Logger {
	result := l.
		With().
		Str(glz.DefaultLogSourceKey, glz.Colorize(glz.DefaultLogSourceColor, glz.DefaultLogWebSource)).
		Logger()
	return &result
}
