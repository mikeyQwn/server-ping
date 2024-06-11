package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type logger struct {
	log *zerolog.Logger
}

func New(cfg *LoggerConfig) Logger {
	log := zerolog.New(os.Stdout)
	return &logger{
		log: &log,
	}
}

func (l *logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *logger) Infof(format string, v any) {
	l.log.Info().Msgf(format, v)
}

func (l *logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *logger) Debugf(format string, v any) {
	l.log.Debug().Msgf(format, v)
}

func (l *logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *logger) Warnf(format string, v any) {
	l.log.Warn().Msgf(format, v)
}

func (l *logger) Error(err error, msg string) {
	l.log.Error().Err(err).Msg(msg)
}

func (l *logger) Errorf(err error, format string, v any) {
	l.log.Error().Err(err).Msgf(format, v)
}

func (l *logger) Fatal(err error, msg string) {
	l.log.Fatal().Err(err).Msg(msg)
}

func (l *logger) Fatalf(err error, format string, v any) {
	l.log.Fatal().Err(err).Msgf(format, v)
}
