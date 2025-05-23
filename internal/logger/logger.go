package logger

import (
	"fmt"
	"github.com/pterm/pterm"
)

type Logger struct {
	logger *pterm.Logger
}

func (l *Logger) Get() *pterm.Logger {
	return l.logger
}

func NewLogger(isVerbose bool) *Logger {

	lvl := pterm.LogLevelInfo
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

func (l *Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Error(msg)
	l.logger.Debug(err.Error())

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
func (l *Logger) MoreWarn(msg string, args map[string]any) {
	l.logger.Warn(msg, l.logger.ArgsFromMap(args))
}

func (l *Logger) Success(msg string) {
	pterm.DefaultBasicText.Println(
		fmt.Sprintf("✅ Success: %s", msg))
}
