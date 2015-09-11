package log

import (
	"fmt"

	mailgunLog "github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/log"
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/mailgun/oxy/utils"
)

const (
	SeverityWarning = "WARN"
	SeverityError   = "ERROR"
	SeverityInfo    = "INFO"
	SeverityDebug   = "DEBUG"
)

type Logger interface {
	SetSeverity(s string) error
	GetSeverity() string

	Infof(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	// Fatalf(format string, args ...interface{})
}

type DefaultLogger struct {
	logSeverity string
}

var currentLogger Logger

func GetGlobalLogger() utils.Logger {
	return currentLogger
}

func NewDefaultLogger(name, severity string) (Logger, error) {
	err := mailgunLog.InitWithConfig(mailgunLog.Config{Name: name, Severity: severity})
	if err != nil {
		return nil, err
	}
	return &DefaultLogger{}, nil
}

func Infof(format string, args ...interface{}) {
	if currentLogger != nil {
		currentLogger.Infof(format, args...)
	}
}

func Warningf(format string, args ...interface{}) {
	if currentLogger != nil {
		currentLogger.Warningf(format, args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if currentLogger != nil {
		currentLogger.Errorf(format, args...)
	}
}

// func Fatalf(format string, args ...interface{}) {
// 	if currentLogger != nil {
// 		currentLogger.Fatalf(format, args...)
// 	}
// }

func GetSeverity() string {
	if currentLogger == nil {
		return ""
	}
	return currentLogger.GetSeverity()
}

func SetSeverity(s string) error {
	if currentLogger == nil {
		return fmt.Errorf("Logger is not set")
	}
	return currentLogger.SetSeverity(s)
}

func EnsureLoggerExist(name, severity string) (err error) {
	if currentLogger == nil {
		currentLogger, err = NewDefaultLogger(name, severity)
	}
	return
}

func SetLogger(l Logger) {
	currentLogger = l
}

func (d DefaultLogger) SetSeverity(s string) error {
	d.logSeverity = s
	lvl, err := mailgunLog.SeverityFromString(s)
	if err != nil {
		return err
	}
	mailgunLog.SetSeverity(lvl)
	return nil
}

func (d DefaultLogger) GetSeverity() string {
	return d.logSeverity
}

func (d DefaultLogger) Infof(format string, args ...interface{}) {
	mailgunLog.Infof(format, args...)
}
func (d DefaultLogger) Warningf(format string, args ...interface{}) {
	mailgunLog.Warningf(format, args...)
}
func (d DefaultLogger) Errorf(format string, args ...interface{}) {
	mailgunLog.Errorf(format, args...)
}

// func (d DefaultLogger) Fatalf(format string, args ...interface{}) {
// 	mailgunLog.Fatalf(format, args...)
// }
