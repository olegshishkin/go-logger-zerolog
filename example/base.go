package example

import (
	"github.com/olegshishkin/go-logger"
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/rs/zerolog"
	"io"
)

// Base creates initial logger instance.
func Base(writer io.Writer, level logger.Level) *zerolog.Logger {
	zeroLevel, err := glz.ToZeroLogLevel(level)
	if err != nil {
		panic(err)
	}

	result := zerolog.New(writer).
		Level(zeroLevel).
		With().
		Timestamp().
		Stack().
		Caller().
		CallerWithSkipFrameCount(3).
		Logger()
	return &result
}
