package cache

import (
	"sync"
	"container/list"
	"github.com/pkg/errors"
	"fmt"
)

type LRUCache struct {
	lock *sync.RWMutex
	sortedlist *list.List
	container map[string]*list.Element
	capacity int
	size int
}

type entry struct {
	key string
	value interface{}
}

func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity <=0 {
		return nil, errors.New("the capacity of LRUCache should be positive number.")
	}
	return &LRUCache{
		lock: new(sync.RWMutex),
		size: 0,
		capacity: capacity,
		container: make(map[string]*list.Element, capacity),
		sortedlist: list.New(),
	}, nil
}


func (this *LRUCache) Cache(key string, data interface{}) (bool, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if element, exists := this.container[key]; exists {
		this.sortedlist.MoveToFront(element)
		element.Value.(*entry).value = data
		return true, nil
	}
	// add new entry
	element := &entry{
		key:key,
		value:data,
	}
	entry := this.sortedlist.PushFront(element)
	this.container[key] = entry
	if this.sortedlist.Len() > this.capacity {
		this.removeOldest()
	}
	this.size += 1
	return true, nil
}

func (this *LRUCache) removeOldest() {
	ent := this.sortedlist.Back()
	if ent != nil {
		this.sortedlist.Remove(ent)
		keyvalue := ent.Value.(*entry)
		delete(this.container, keyvalue.key)
		this.size -=1
	}
}


func (this *LRUCache) Size() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.size
}

func (this *LRUCache) Get(key string) (interface{}, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	// method Get need to update the order also.
	if element, exists := this.container[key]; exists == false {
		return nil, nil
	}else {
		this.sortedlist.MoveToFront(element)
		return element.Value.(*entry).value, nil
	}
}

func (this *LRUCache) Delete(key string) (bool, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	element := &list.Element{Value:key}
	if _, exists := this.container[key]; exists == false {
		return true, nil
	}else{
		this.sortedlist.Remove(element)
		delete(this.container, key)
		this.size -= 1
		return true, nil
	}
}

func (this *LRUCache) GetCachedData() (map[string]interface{}, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.size <= 0 {
		return nil, errors.New("current LRUCache is empty.")
	}
	// transfer to escape *list.Element
	ret := make(map[string]interface{})
	for key, element := range this.container {
		ret[key] = element
	}
	fmt.Println(this.sortedlist)
	return ret, nil
}