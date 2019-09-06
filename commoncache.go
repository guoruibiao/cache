package cache

import (
	"sync"
	"github.com/pkg/errors"
)

// It seems useless, because we can not always put all the data into this cache container.
type CommonCache struct {
	lock *sync.RWMutex
	container map[string]interface{}
}

func (this *CommonCache) GetCachedData() (map[string]interface{}, error) {
	if this.container != nil {
		return this.container, nil
	}
	return nil, errors.New("current cache container is empty.")
}


func NewCommonCacher() *CommonCache {
	return &CommonCache{
		container: make(map[string]interface{}),
		lock: new(sync.RWMutex),
	}
}

func (this *CommonCache) Cache(key string, data interface{}) (bool, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.container[key] = data
	return true, nil
}

func (this *CommonCache) Get(key string) (interface{}, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	data, exist := this.container[key]
	if exist == false {
		return nil, errors.New("no such data for key: " + key)
	}
	return data, nil
}

func (this *CommonCache) Delete(key string) (bool, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	// just an useless judgement
	if _, exists := this.container[key]; exists == false {
		return true, nil
	} else {
		delete(this.container, key)
		return true, nil
	}
}