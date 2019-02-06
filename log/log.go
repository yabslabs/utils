package log

import (
	"github.com/sirupsen/logrus"
	"github.com/yabslabs/utils/pairs"
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
	return &Entry{Entry: logrus.WithField(idKey, id)}
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
	e.log(e.Debug, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.log(e.Warn, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.log(e.Error, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.log(e.Fatal, args...)
}

func (e *Entry) Panic(args ...interface{}) {
	e.log(e.Panic, args...)
}

func (e *Entry) log(log func(...interface{}), args ...interface{}) {
	if e.onError && e.err != nil {
		log(args...)
	}
	if !e.onError {
		log(args...)
	}
}
