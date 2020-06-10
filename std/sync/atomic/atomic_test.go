package atomic

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Cache struct {
	data sync.Map
	size int32
}

func (c *Cache) store(key string, value interface{}) {
	c.data.Store(key, value)
	atomic.AddInt32(&c.size, 1)
}

func (c *Cache) length() int {
	return int(atomic.LoadInt32(&c.size))
}
func (c *Cache) delete(key string) {
	if _, ok := c.data.Load(key); ok {
		c.data.Delete(key)
		atomic.AddInt32(&c.size, -1)
	}
}

func (c *Cache) ranges() {
	c.data.Range(func(key, value interface{}) bool {
		log.Println(key, value)
		return true
	})
}

func TestName(t *testing.T) {
	count := int32(0)
	atomic.AddInt32(&count, 1)
	log.Println(atomic.LoadInt32(&count))
	atomic.AddInt32(&count, 10)
	log.Println(atomic.LoadInt32(&count))
	atomic.AddInt32(&count, -5)
	log.Println(atomic.LoadInt32(&count))
	log.Println(count)
	log.Println(time.Time{}.String())
	log.Println(time.Now().Sub(time.Time{}))
}

func TestCache(t *testing.T) {
	cache := Cache{}
	cache.store("1", 1)
	cache.store("2", 2)
	cache.store("3", 3)
	log.Println(cache.length())
	cache.delete("1")
	cache.delete("2")
	log.Println(cache.length())
	cache.ranges()
	cache.store("11", 11)
	cache.store("12", 11)
	cache.store("13", 11)
	cache.store("14", 11)
	cache.store("15", 11)
	cache.ranges()
	log.Println(cache.length())
	cache.delete("11")
	log.Println(cache.length())
	log.Println(cache.size)
}
