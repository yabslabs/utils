package fastcache

import "github.com/yabslabs/utils/cache"

type Config struct {
	MaxCacheSizeInByte int
}

func (c *Config) NewCache() (cache.Cache, error) {
	return NewFastcache(c)
}
