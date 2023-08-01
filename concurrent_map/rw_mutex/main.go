package main

import (
	"github.com/davecgh/go-spew/spew"
	"sync"
)

type ConcurrentMap struct {
	sync.RWMutex
	items map[string]interface{}
}

func main() {
	c := New()
	c.Set("a1", 32)
	c.Set("a2", "Good Evening")
	spew.Dump(c.Get("a2"))
}

func New() *ConcurrentMap {
	cm := new(ConcurrentMap)
	cm.items = make(map[string]interface{})
	return cm
}

func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.Lock()
	defer cm.Unlock()
	cm.items[key] = value
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.RLock()
	defer cm.RUnlock()
	value, ok := cm.items[key]
	return value, ok
}
