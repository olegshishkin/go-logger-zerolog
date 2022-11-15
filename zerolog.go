package zerolog

import (
	"fmt"
	"github.com/olegshishkin/go-logger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

const (
	LogSourceKey  = "logger"
	colorDarkGray = 90
)

type writerBuilder struct {
	console *io.Writer
	file    *io.Writer
}

type zeroLogger struct {
	log *zerolog.Logger
}

func (l *zeroLogger) LogLevel(level logger.Level) error {
	zeroLevel, err := toZeroLogLevel(level)
	if err != nil {
		return err
	}
	l.log.Level(zeroLevel)
	return nil
}

func (l *zeroLogger) Trace(msg string, vars ...any) {
	l.log.Trace().Msgf(msg, vars...)
}

func (l *zeroLogger) Debug(msg string, vars ...any) {
	l.log.Debug().Msgf(msg, vars...)
}

func (l *zeroLogger) Info(msg string, vars ...any) {
	l.log.Info().Msgf(msg, vars...)
}

func (l *zeroLogger) Warn(msg string, vars ...any) {
	l.log.Warn().Msgf(msg, vars...)
}

func (l *zeroLogger) Error(err error, msg string, vars ...any) {
	l.log.Error().Err(err).Msgf(msg, vars...)
}

func (l *zeroLogger) Fatal(err error, msg string, vars ...any) {
	l.log.Fatal().Err(err).Msgf(msg, vars...)
}

func Logger(writer io.Writer, level logger.Level) (logger.Logger, error) {
	zeroLevel, err := toZeroLogLevel(level)
	if err != nil {
		return nil, err
	}

	log := zerolog.New(writer).
		Level(zeroLevel).
		With().
		Timestamp().
		Stack().
		Caller().
		CallerWithSkipFrameCount(3).
		Str(LogSourceKey, colorize(colorDarkGray, "common")).
		Logger()
	return &zeroLogger{&log}, nil
}

func toZeroLogLevel(level logger.Level) (zerolog.Level, error) {
	switch level {
	case logger.Trace:
		return zerolog.TraceLevel, nil
	case logger.Debug:
		return zerolog.DebugLevel, nil
	case logger.Info:
		return zerolog.InfoLevel, nil
	case logger.Warn:
		return zerolog.WarnLevel, nil
	case logger.Error:
		return zerolog.ErrorLevel, nil
	case logger.Fatal:
		return zerolog.FatalLevel, nil
	default:
		return zerolog.Disabled, errors.Errorf("unknown log level: %v", level)
	}
}

func colorize(c int, m any) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, m)
}

func NewLogWriterBuilder() *writerBuilder {
	return &writerBuilder{}
}

func (b *writerBuilder) WithConsole() *writerBuilder {
	var writer io.Writer = zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    time.RFC3339Nano,
		FieldsExclude: []string{LogSourceKey},
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			LogSourceKey,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	}
	return b.WithConsoleWriter(writer)
}

func (b *writerBuilder) WithFile(file *os.File) *writerBuilder {
	var writer io.Writer = &lumberjack.Logger{
		Filename:   file.Name(),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return b.WithFileWriter(writer)
}

func (b *writerBuilder) WithConsoleWriter(writer io.Writer) *writerBuilder {
	b.console = &writer
	return b
}

func (b *writerBuilder) WithFileWriter(writer io.Writer) *writerBuilder {
	b.file = &writer
	return b
}

func (b *writerBuilder) Build() (io.Writer, error) {
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
