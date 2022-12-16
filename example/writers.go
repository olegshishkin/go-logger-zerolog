package example

import (
	glz "github.com/olegshishkin/go-logger-zerolog"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

const (
	logFileMaxSize    = 10
	logFileMaxBackups = 5
	logFileMaxAge     = 30
)

// ConsoleWriter creates an example of the writer instance for logging to the console.
func ConsoleWriter() io.Writer {
	return &zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    time.RFC3339Nano,
		FieldsExclude: []string{glz.DefaultLogSourceKey},
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			glz.DefaultLogSourceKey,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	}
}

// FileWriter creates an example of the writer instance for logging to the file.
func FileWriter(file *os.File) io.Writer {
	return &lumberjack.Logger{
		Filename:   file.Name(),
		MaxSize:    logFileMaxSize,
		MaxBackups: logFileMaxBackups,
		MaxAge:     logFileMaxAge,
		Compress:   true,
	}
}
