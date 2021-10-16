package cache

import "errors"

type Cache struct {
	memory map[string]interface{}
}

func New() Cache {
	return Cache{
		make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) error {
	_, ok := c.memory[key]
	if !ok {
		c.memory[key] = value
		return nil
	} else {
		return errors.New("value already exists")
	}
}

func (c Cache) Get(key string) interface{} {
	_, ok := c.memory[key]
	if ok {
		return c.memory[key]
	} else {
		return errors.New("value not found")
	}

}

func (c Cache) Delete(key string) bool {
	_, ok := c.memory[key]
	if ok {
		delete(c.memory, key)
	}
	return ok
}
