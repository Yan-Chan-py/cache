package cache
import (
	"log"
)

type Cache struct {
	storage   map[string]interface{}


}

func New()Cache {
	return Cache{
		storage: map[string]interface{}{"0":"empty"} ,
	}
}


func inStorage(c Cache,key string) bool {
	_,exists := c.storage[key]
	return exists
	}

	
	
func(c *Cache) Set(key string, value interface{}) {
	c.storage[key] = value

}
func (cache *Cache) Get(key string) interface {} {
	if inStorage(*cache,key) {
		return cache.storage[key]
	}
	return 404
}


func (c *Cache) Delete(key string) {
	if inStorage(*c,key) {
		delete(c.storage,key)
	} else {
		log.Fatal("No such value")
	}

}

