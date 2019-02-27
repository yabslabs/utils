package config

import (
	"testing"

	conf_test "github.com/yabslabs/utils/config/testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type TestConfigLog struct {
	Log Log
}

func TestLog(t *testing.T) {
	yamlConfig := `Log:
  Level: 'info'
  Formatter: 'json'`

	jsonConfig := `{
	"Log": {
		"Level": "info",
		"Formatter": "json"
	}
}`

	tomlConfig := `Log="{\"Level\": \"info\", \"Formatter\": \"json\"}"`

	r := conf_test.AssertConfigRead(t, map[conf_test.TestConfigReaderIdentifier]string{
		conf_test.TestTOMLIdentifier: tomlConfig,
		conf_test.TestYAMLIdentifier: yamlConfig,
		conf_test.TestJSONIdentifier: jsonConfig,
	}, &TestConfigLog{})

	for _, result := range r {
		resultT := result.(*TestConfigLog)
		assert.Equal(t, logrus.InfoLevel, resultT.Log.Logger.Level)
		assert.Equal(t, &logrus.JSONFormatter{}, resultT.Log.Logger.Formatter)
	}

	assert.Equal(t, logrus.InfoLevel, logrus.GetLevel())
	assert.Equal(t, &logrus.JSONFormatter{}, logrus.StandardLogger().Formatter)

	LoggerApplyToGlobal = false

	conf_test.AssertConfigRead(t, map[conf_test.TestConfigReaderIdentifier]string{
		conf_test.TestJSONIdentifier: `{ "Log": { "Level": "info", "Formatter": "json" } }`,
	}, &TestConfigLog{})

	assert.Equal(t, logrus.InfoLevel, logrus.GetLevel())
	assert.Equal(t, &logrus.JSONFormatter{}, logrus.StandardLogger().Formatter)
}
