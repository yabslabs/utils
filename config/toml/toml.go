package toml

import (
	"github.com/BurntSushi/toml"

	"github.com/yabslabs/utils/config"
)

var ConfigReader = config.ConfigReaderFunc(toml.Unmarshal)

func ReadConfig(obj interface{}, configFiles ...string) error {
	return config.ReadConfig(ConfigReader, obj, configFiles...)
}
