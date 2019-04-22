package yaml

import (
	"github.com/ghodss/yaml"

	"github.com/yabslabs/utils/config"
)

var ConfigReader = config.ConfigReaderFunc(yamlUnmarshalNoOpts)

func yamlUnmarshalNoOpts(y []byte, o interface{}) error {
	return yaml.Unmarshal(y, o)
}

func ReadConfig(obj interface{}, configFiles ...string) error {
	return config.ReadConfig(ConfigReader, obj, configFiles...)
}
