package logger

import (
	"github.com/pterm/pterm"
)

type Logger struct {
	logger *pterm.Logger
}

func NewLogger(isVerbose bool) *Logger {

	lvl := pterm.LogLevelWarn
	if isVerbose {

		lvl = pterm.LogLevelTrace

	}

	lg := pterm.DefaultLogger.WithLevel(lvl)

	return &Logger{
		logger: lg,
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *Logger) Trace(msg string) {
	l.logger.Trace(msg)
}

func (l *Logger) MoreInfo(msg string, args map[string]any) {
	l.logger.Info(msg, l.logger.ArgsFromMap(args))
}

func (l *Logger) MoreError(msg string, args map[string]any) {
	l.logger.Error(msg, l.logger.ArgsFromMap(args))
}

func (l *Logger) MoreFatal(msg string, args map[string]any) {
	l.logger.Fatal(msg, l.logger.ArgsFromMap(args))
}

func (l *Logger) MoreTrace(msg string, args map[string]any) {
	l.logger.Trace(msg, l.logger.ArgsFromMap(args))
}
