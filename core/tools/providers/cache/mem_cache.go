package cache

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

//MemCache struct for assign ICacheProvider
type MemCache struct {
	Client     *memcache.Client
	Expiration time.Duration
}

var singletonConnection *MemCache = nil

//New create a new instance
func New(host string) (*MemCache, error) {
	if singletonConnection == nil {
		client := memcache.New(host)
		if err := client.Ping(); err != nil {
			return nil, err
		}

		return &MemCache{
			Client: client,
		}, nil
	}
	return singletonConnection, nil
}

//Assign interface
var _ ICacheProvider = &MemCache{}

func (cache *MemCache) WithExpiration(expiration time.Duration) ICacheProvider {
	cache.Expiration = expiration
	return cache
}

func (cache *MemCache) Set(key string, value interface{}) error {
	valueBytes, err := cache.encode(value)
	if err != nil {
		return err
	}

	cache.Client.Set(&memcache.Item{
		Key:        key,
		Value:      valueBytes,
		Expiration: int32(cache.Expiration),
	})
	return nil
}

func (cache *MemCache) Get(key string, value interface{}) error {
	item, err := cache.Client.Get(key)
	if err != nil {
		return err
	}

	if err := cache.decode(item.Value, value); err != nil {
		return err
	}

	return nil
}

func (cache *MemCache) Delete(key string) error {
	return cache.Client.Delete(key)
}

func (cache *MemCache) Clear() error {
	return cache.Client.FlushAll()
}

func (cache *MemCache) encode(v interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (cache *MemCache) decode(data []byte, v interface{}) error {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)

	return decoder.Decode(v)
}
