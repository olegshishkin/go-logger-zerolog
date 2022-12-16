package zerolog

import (
	"github.com/olegshishkin/go-logger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Wrapper struct {
	log *zerolog.Logger
}

// From transforms source Zero logger to Logger type.
func From(log *zerolog.Logger) *Wrapper {
	return &Wrapper{log}
}

func (w *Wrapper) GetLevel() logger.Level {
	l, err := ToLogLevel(w.log.GetLevel())
	if err != nil {
		return logger.Fatal
	}

	return l
}

func (w *Wrapper) SetLevel(l logger.Level) error {
	level, err := ToZeroLogLevel(l)
	if err != nil {
		return err
	}

	w.log.Level(level)

	return nil
}

func (w *Wrapper) Trace(msg string, vars ...any) {
	w.log.Trace().Msgf(msg, vars...)
}

func (w *Wrapper) Debug(msg string, vars ...any) {
	w.log.Debug().Msgf(msg, vars...)
}

func (w *Wrapper) Info(msg string, vars ...any) {
	w.log.Info().Msgf(msg, vars...)
}

func (w *Wrapper) Warn(msg string, vars ...any) {
	w.log.Warn().Msgf(msg, vars...)
}

func (w *Wrapper) Error(err error, msg string, vars ...any) {
	w.log.Error().Err(err).Msgf(msg, vars...)
}

func (w *Wrapper) Fatal(err error, msg string, vars ...any) {
	w.log.Fatal().Err(err).Msgf(msg, vars...)
}

// ToLogLevel transforms Zero log level type to Logger log level type.
func ToLogLevel(level zerolog.Level) (logger.Level, error) {
	switch level {
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
	case zerolog.FatalLevel, zerolog.PanicLevel, zerolog.NoLevel, zerolog.Disabled:
		return logger.Fatal, nil
	default:
		return logger.Fatal, errors.Errorf("unknown log level: %v", level)
	}
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
