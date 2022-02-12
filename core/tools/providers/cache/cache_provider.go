package cache

import (
	"time"
)

//ICacheProvider interface of cache
type ICacheProvider interface {
	WithExpiration(time.Duration) ICacheProvider
	Set(key string, value interface{}) error
	Get(key string, value interface{}) error
	Delete(key string) error
	Clear() error
}
