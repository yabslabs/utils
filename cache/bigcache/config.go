package bigcache

import (
	"time"

	"github.com/yabslabs/utils/cache"
)

type Config struct {
	MaxCacheSizeInMB int
	// CacheLifetimeSeconds if set, cache makes cleanup every minute
	CacheLifetimeSeconds time.Duration
}

func (c *Config) NewCache() (cache.Cache, error) {
	return NewBigcache(c)
}
