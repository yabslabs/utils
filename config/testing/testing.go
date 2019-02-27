package testing

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yabslabs/utils/config"
	"github.com/yabslabs/utils/config/json"
	"github.com/yabslabs/utils/config/toml"
	"github.com/yabslabs/utils/config/yaml"
)

type TestConfigReaderIdentifier int

const (
	TestTOMLIdentifier TestConfigReaderIdentifier = iota
	TestYAMLIdentifier
	TestJSONIdentifier
)

var (
	TestConfigReaders = map[TestConfigReaderIdentifier]config.ConfigReader{
		TestTOMLIdentifier: toml.ConfigReader,
		TestYAMLIdentifier: yaml.ConfigReader,
		TestJSONIdentifier: json.ConfigReader,
	}
)

func TestReadConfigForIdentifier(t *testing.T, identifier TestConfigReaderIdentifier, configContent string, expectedConfig interface{}) interface{} {
	configType := reflect.TypeOf(expectedConfig)
	if configType.Kind() != reflect.Ptr {
		t.Errorf("expected config is not a pointer: %s", configType.Kind())
	}
	config := reflect.New(configType.Elem())
	err := TestConfigReaders[identifier].Unmarshal([]byte(configContent), config.Interface())
	assert.NoError(t, err)
	return config.Interface()
}

func AssertConfigRead(t *testing.T, configs map[TestConfigReaderIdentifier]string, expectedConfig interface{}) map[TestConfigReaderIdentifier]interface{} {
	result := map[TestConfigReaderIdentifier]interface{}{}
	for id, c := range configs {
		result[id] = TestReadConfigForIdentifier(t, id, c, expectedConfig)
	}
	return result
}

func AssertConfigReadEquals(t *testing.T, configs map[TestConfigReaderIdentifier]string, expectedConfig interface{}) {
	r := AssertConfigRead(t, configs, expectedConfig)
	for _, result := range r {
		assert.Equal(t, expectedConfig, result)
	}
}
