package generic

import (
	"path/filepath"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yabslabs/utils/config"
	"github.com/yabslabs/utils/config/json"
	"github.com/yabslabs/utils/config/toml"
	"github.com/yabslabs/utils/config/yaml"
)

var (
	errDifferentFileTypes     = status.Error(codes.InvalidArgument, "different file types")
	errNoConfigFiles          = status.Error(codes.InvalidArgument, "no config files provided")
	errNoMatchingConfigReader = status.Error(codes.InvalidArgument, "no matching config reader")

	extReaders = map[string]config.ConfigReaderFunc{
		".yaml": yaml.ConfigReader,
		".yml":  yaml.ConfigReader,
		".json": json.ConfigReader,
		".toml": toml.ConfigReader,
	}
)

func ReadConfig(obj interface{}, configFiles ...string) error {
	if len(configFiles) == 0 {
		return errNoConfigFiles
	}

	ext := strings.ToLower(filepath.Ext(configFiles[0]))
	cfgReader, ok := extReaders[ext]
	if !ok {
		return errNoMatchingConfigReader
	}

	if len(configFiles) > 1 {
		for _, f := range configFiles[1:] {
			if ext != strings.ToLower(filepath.Ext(f)) {
				return errDifferentFileTypes
			}
		}
	}

	return config.ReadConfig(cfgReader, obj, configFiles...)
}
