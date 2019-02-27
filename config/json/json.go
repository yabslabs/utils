package json

import (
	"encoding/json"

	"github.com/yabslabs/utils/config"
)

var ConfigReader = config.ConfigReaderFunc(json.Unmarshal)

func ReadConfig(obj interface{}, configFiles ...string) error {
	return config.ReadConfig(ConfigReader, obj, configFiles...)
}
