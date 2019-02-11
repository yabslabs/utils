package logging

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/yabslabs/utils/pairs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Entry struct {
	*logrus.Entry
	err     error
	onError bool
}

var idKey = "logID"

// SetIDKey key of id in logentry
func SetIDKey(key string) {
	idKey = key
}

// Log creates a new entry with an id
func Log(id string) *Entry {
	return &Entry{Entry: logrus.WithField(idKey, id), onError: false}
}

// LogWithFields creates a new entry with an id and the given fields
func LogWithFields(id string, fields ...interface{}) *Entry {
	logFields := pairs.Pairs(fields...)
	logFields[idKey] = id
	return &Entry{Entry: logrus.WithFields(logFields)}
}

// SetFields sets the given fields on the entry. It panics if length of fields is odd
func (e *Entry) SetFields(fields ...interface{}) *Entry {
	logFields := pairs.Pairs(fields...)
	e.WithFields(logFields)
	return e
}

// OnError sets the error. The log will only be printed if err is not nil
func (e *Entry) OnError(err error) *Entry {
	e.err = err
	e.onError = true
	return e
}

func (e *Entry) Debug(args ...interface{}) {
	e.Log(logrus.DebugLevel, args...)
}

func (e *Entry) Debugln(args ...interface{}) {
	e.Logln(logrus.DebugLevel, args...)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.Logf(logrus.DebugLevel, format, args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.Log(logrus.InfoLevel, args...)
}

func (e *Entry) Infoln(args ...interface{}) {
	e.Logln(logrus.InfoLevel, args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.Logf(logrus.InfoLevel, format, args...)
}

func (e *Entry) Trace(args ...interface{}) {
	e.Log(logrus.TraceLevel, args...)
}

func (e *Entry) Traceln(args ...interface{}) {
	e.Logln(logrus.TraceLevel, args...)
}

func (e *Entry) Tracef(format string, args ...interface{}) {
	e.Logf(logrus.TraceLevel, format, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.Log(logrus.WarnLevel, args...)
}

func (e *Entry) Warnln(args ...interface{}) {
	e.Logln(logrus.WarnLevel, args...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.Logf(logrus.WarnLevel, format, args...)
}

func (e *Entry) Warning(args ...interface{}) {
	e.Log(logrus.WarnLevel, args...)
}

func (e *Entry) Warningln(args ...interface{}) {
	e.Logln(logrus.WarnLevel, args...)
}

func (e *Entry) Warningf(format string, args ...interface{}) {
	e.Logf(logrus.WarnLevel, format, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.Log(logrus.ErrorLevel, args...)
}

func (e *Entry) Errorln(args ...interface{}) {
	e.Logln(logrus.ErrorLevel, args...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.Logf(logrus.ErrorLevel, format, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.Log(logrus.FatalLevel, args...)
}

func (e *Entry) Fatalln(args ...interface{}) {
	e.Logln(logrus.FatalLevel, args...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.Logf(logrus.FatalLevel, format, args...)
}

func (e *Entry) Panic(args ...interface{}) {
	e.Log(logrus.PanicLevel, args...)
}

func (e *Entry) Panicln(args ...interface{}) {
	e.Logln(logrus.PanicLevel, args...)
}

func (e *Entry) Panicf(format string, args ...interface{}) {
	e.Logf(logrus.PanicLevel, format, args...)
}

func (e *Entry) Log(level logrus.Level, args ...interface{}) {
	if e.onError && e.err != nil {
		e.errorLog().Log(level, args...)
	}
	if !e.onError {
		e.Entry.Log(level, args...)
	}
}

func (e *Entry) Logf(level logrus.Level, format string, args ...interface{}) {
	if e.onError && e.err != nil {
		e.errorLog().Logf(level, format, args...)
	}
	if !e.onError {
		e.Entry.Logf(level, format, args...)
	}
}

func (e *Entry) Logln(level logrus.Level, args ...interface{}) {
	if e.onError && e.err != nil {
		e.errorLog().Logln(level, args...)
	}
	if !e.onError {
		e.Entry.Logln(level, args...)
	}
}

func (e *Entry) WithError(err error) *Entry {
	e.err = err
	e.Entry = e.errorLog()
	return e
}

func (e *Entry) errorLog() *logrus.Entry {
	if s, ok := errorToStatus(e.err); ok {
		return e.WithFields(logrus.Fields{logrus.ErrorKey: s.Message(), "httpCode": runtime.HTTPStatusFromCode(s.Code()), "grpcCode": s.Code()})
	}
	return e.Entry.WithError(e.err)
}

func errorToStatus(err error) (*status.Status, bool) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.FromContextError(err)
		return s, s.Code() != codes.Unknown
	}
	return s, true
}
