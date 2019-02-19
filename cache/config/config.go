package config

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cmn_cache "github.com/yabslabs/utils/cache"
	cmn_cache_bc "github.com/yabslabs/utils/cache/bigcache"
	cmn_cache_fc "github.com/yabslabs/utils/cache/fastcache"
)

type CacheConfig struct {
	ID     string
	Type   string
	Config cmn_cache.Config
}

var caches = map[string]func() cmn_cache.Config{
	"bigcache":  func() cmn_cache.Config { return &cmn_cache_bc.Config{} },
	"fastcache": func() cmn_cache.Config { return &cmn_cache_fc.Config{} },
}

func (c *CacheConfig) UnmarshalJSON(data []byte) error {
	var rc struct {
		ID     string
		Type   string
		Config json.RawMessage
	}

	if err := json.Unmarshal(data, &rc); err != nil {
		return status.Errorf(codes.Internal, "%v parse config: %v", "CACHE-vmjS", err)
	}

	c.Type = rc.Type
	c.ID = rc.ID

	var err error
	c.Config, err = newCacheConfig(c.Type, rc.Config)
	if err != nil {
		return status.Errorf(codes.Internal, "%v parse config: %v", "CACHE-Ws9E", err)
	}

	return nil
}

func newCacheConfig(cacheType string, configData []byte) (cmn_cache.Config, error) {
	t, ok := caches[cacheType]
	if !ok {
		return nil, status.Errorf(codes.Internal, "%v No config: %v", "CACHE-HMEJ", cacheType)
	}

	cacheConfig := t()
	if len(configData) == 0 {
		return cacheConfig, nil
	}

	if err := json.Unmarshal(configData, cacheConfig); err != nil {
		return nil, status.Errorf(codes.Internal, "%v Could not read conifg: %v", "CACHE-1tSS", err)
	}

	return cacheConfig, nil
}
