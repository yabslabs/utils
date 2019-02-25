package fastcache

import (
	"encoding/json"

	"github.com/VictoriaMetrics/fastcache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yabslabs/utils/logging"
)

type Fastcache struct {
	cache *fastcache.Cache
}

func NewFastcache(config *Config) (*Fastcache, error) {
	return &Fastcache{
		cache: fastcache.New(config.MaxCacheSizeInByte),
	}, nil
}

func (fc *Fastcache) Set(key string, object interface{}) error {
	marshalled, err := json.Marshal(object)
	if err != nil {
		logging.Log("FASTC-uCEEDB").WithError(err).Debug("unable to marshall object into json")
		return status.Error(codes.InvalidArgument, "unable to marshall object into json")
	}
	fc.cache.Set([]byte(key), marshalled)
	return nil
}

func (fc *Fastcache) Get(key string, ptrToObject interface{}) error {
	data := fc.cache.Get(nil, []byte(key))
	if len(data) == 0 {
		return status.Error(codes.NotFound, "not in cache")
	}
	return json.Unmarshal(data, ptrToObject)
}

func (fc *Fastcache) Delete(key string) error {
	fc.cache.Del([]byte(key))
	return nil
}
