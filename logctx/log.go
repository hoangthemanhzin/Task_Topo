package logctx

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

var _logger *loggerWrapper

type loggerWrapper struct {
	*logrus.Entry
	instance *logrus.Logger
}

func init() {
	instance := logrus.New()
	instance.SetLevel(logrus.Level(DEFAULT_LOG_LEVEL))
	_logger = &loggerWrapper{
		instance: instance,
		Entry:    instance.WithField("logger", "generic"), //default root
	}
}

//initialize logger, set fields for application root writer
func Init(f *os.File, level LogLevel, fields Fields) (err error) {
	_logger.instance.SetLevel(logrus.Level(level))
	if f != nil {
		_logger.instance.SetOutput(f)
	}
	_logger.Entry = _logger.instance.WithFields(logrus.Fields(fields)) //updated root
	return
}

//new writer inheriting fields from the application root logger
func WithFields(fields Fields) LogWriter {
	return &writer{
		Entry: _logger.WithFields(logrus.Fields(fields)),
	}
}

type writer struct {
	*logrus.Entry
}

//add fields to writer
func (w *writer) WithFields(fields Fields) LogWriter {
	return &writer{
		Entry: w.Entry.WithFields(logrus.Fields(fields)),
	}
}

type LogWriter interface {
	Trace(args ...interface{})
	Debug(args ...interface{})
	Print(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Traceln(args ...interface{})
	Debugln(args ...interface{})
	Println(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	WithFields(Fields) LogWriter
}
