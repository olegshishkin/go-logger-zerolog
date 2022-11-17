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
func From(log *zerolog.Logger) logger.Logger {
	return &wrapper{log}
}

func (w *wrapper) LogLevel(level logger.Level) error {
	zeroLevel, err := ToZeroLogLevel(level)
	if err != nil {
		return err
	}
	w.log.Level(zeroLevel)
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

// ToZeroLogLevel transforms Logger log level type to Zero log level type.
func ToZeroLogLevel(level logger.Level) (zerolog.Level, error) {
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
