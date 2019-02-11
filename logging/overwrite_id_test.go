package logging_test

import (
	"bytes"
	"testing"

	"github.com/yabslabs/utils/logging"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	_ "github.com/yabslabs/utils/logging/internal"
)

func TestOverwriteLogIDKey(t *testing.T) {
	buf := initBuf()
	logging.Log("UTILS-B7l7").Warn("check", "check")
	assert.Contains(t, buf.String(), "level=warning")
	assert.Contains(t, buf.String(), "msg=checkcheck")
	assert.Contains(t, buf.String(), "MyLoggingID=UTILS-B7l7")
	assert.NotContains(t, buf.String(), "LogID")
}

func initBuf() *bytes.Buffer {
	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	return &buf
}
