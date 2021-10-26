package cache

import (
	"errors"
	"time"
)

type Cache struct {
	memory map[string]interface{}
}

func New() Cache {
	return Cache{
		make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.memory[key] = value
	go func() {
		time.Sleep(ttl)
		delete(c.memory, key)
	}()
}

func (c Cache) Get(key string) (interface{}, error) {
	value, ok := c.memory[key]
	if ok {
		return value, nil
	}
	return value, errors.New("value not found")

}

func (c Cache) Delete(key string) bool {
	_, ok := c.memory[key]
	if ok {
		delete(c.memory, key)
	}
	return ok
}
