package zerolog

import (
	"github.com/olegshishkin/go-logger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type wrapper struct {
	log *zerolog.Logger
}

// From transforms source Zero logger to Logger type.
func From(log *zerolog.Logger) *wrapper {
	return &wrapper{log}
}

func (w *wrapper) GetLevel() logger.Level {
	l, err := ToLogLevel(w.log.GetLevel())
	if err != nil {
		return logger.Fatal
	}
	return l
}

func (w *wrapper) SetLevel(l logger.Level) error {
	level, err := ToZeroLogLevel(l)
	if err != nil {
		return err
	}
	w.log.Level(level)
	return nil
}

func (w *wrapper) Trace(msg string, vars ...any) {
	w.log.Trace().Msgf(msg, vars...)
}

func (w *wrapper) Debug(msg string, vars ...any) {
	w.log.Debug().Msgf(msg, vars...)
}

func (w *wrapper) Info(msg string, vars ...any) {
	w.log.Info().Msgf(msg, vars...)
}

func (w *wrapper) Warn(msg string, vars ...any) {
	w.log.Warn().Msgf(msg, vars...)
}

func (w *wrapper) Error(err error, msg string, vars ...any) {
	w.log.Error().Err(err).Msgf(msg, vars...)
}

func (w *wrapper) Fatal(err error, msg string, vars ...any) {
	w.log.Fatal().Err(err).Msgf(msg, vars...)
}

// ToLogLevel transforms Zero log level type to Logger log level type.
func ToLogLevel(l zerolog.Level) (logger.Level, error) {
	switch l {
	case zerolog.TraceLevel:
		return logger.Trace, nil
	case zerolog.DebugLevel:
		return logger.Debug, nil
	case zerolog.InfoLevel:
		return logger.Info, nil
	case zerolog.WarnLevel:
		return logger.Warn, nil
	case zerolog.ErrorLevel:
		return logger.Error, nil
	case zerolog.FatalLevel:
		return logger.Fatal, nil
	default:
		return logger.Fatal, errors.Errorf("unknown log level: %v", l)
	}
}

// ToZeroLogLevel transforms Logger log level type to Zero log level type.
func ToZeroLogLevel(l logger.Level) (zerolog.Level, error) {
	switch l {
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
		return zerolog.Disabled, errors.Errorf("unknown log level: %v", l)
	}
}
