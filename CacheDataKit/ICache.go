package CacheDataKit

import "fmt"

const V3 = 3

type ICache interface {
	Version() int
	StartAndGC(mpConfig map[string]interface{}) error
	SetCache(key string, value interface{}, timeout int) error
	GetCache(key string, to interface{}) error
	DelCache(key string) error
}

// Instance is a function create a new Cache Instance
type Instance func() ICache

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewCache(adapterName string, mpConfig map[string]interface{}) (adapter ICache, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.StartAndGC(mpConfig)
	if err != nil {
		adapter = nil
	}
	return
}
