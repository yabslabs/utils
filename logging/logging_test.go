package logging

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func initBuf() *bytes.Buffer {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	return &buf
}

func TestLogWarn(t *testing.T) {
	buf := initBuf()
	Log("UTILS-B7l7").Warn("check", "check")
	assert.Contains(t, buf.String(), "level=warning")
	assert.Contains(t, buf.String(), "msg=checkcheck")
	assert.Contains(t, buf.String(), "logID=UTILS-B7l7")
}

func TestLogError(t *testing.T) {
	buf := initBuf()
	Log("UTILS-Ld9V").OnError(fmt.Errorf("im an error")).Warn("error ocured")
	assert.Contains(t, buf.String(), "UTILS-Ld9V")
	assert.Contains(t, buf.String(), "level=warning")
	assert.Contains(t, buf.String(), "msg=\"error ocured\"")
	assert.Contains(t, buf.String(), "error=\"im an error\"")
}

func TestErrorLogStatus(t *testing.T) {
	buf := initBuf()
	s := status.Error(codes.NotFound, "something not found")
	Log("UTILS-457u").OnError(s).Warnln("lol")
	assert.Contains(t, buf.String(), "level=warning")
	assert.Contains(t, buf.String(), "msg=lol")
	assert.Contains(t, buf.String(), "error=\"something not found\"")
	assert.Contains(t, buf.String(), "grpcCode=NotFound")
	assert.Contains(t, buf.String(), "httpCode=404")
}

func TestErrorLog(t *testing.T) {
	buf := initBuf()
	Log("UTILS-457u").OnError(fmt.Errorf("big mistake")).Warnln("lol")
	assert.Contains(t, buf.String(), "level=warning")
	assert.Contains(t, buf.String(), "msg=lol")
	assert.Contains(t, buf.String(), "error=\"big mistake\"")
	assert.NotContains(t, buf.String(), "grpcCode")
	assert.NotContains(t, buf.String(), "httpCode")
}
