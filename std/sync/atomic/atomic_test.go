package atomic

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Data struct {
	value string
	time  time.Time
}
type Cache struct {
	data sync.Map
	size int32
}

func (c *Cache) watch() {
	ticket := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticket.C:
			log.Println("before", c.size)
			c.data.Range(func(key, value interface{}) bool {
				k := key.(string)
				v := value.(Data)
				if time.Now().Sub(v.time) > time.Second*5 {
					c.delete(k)
				}
				return true
			})
			log.Println("end", c.size)
		}
	}
}

func (c *Cache) store(key string, value Data) {
	if _, ok := c.data.Load(key); !ok {
		c.data.Store(key, value)
		atomic.AddInt32(&c.size, 1)
		log.Println("存储", key, "长度", c.size)
	}
}

func (c *Cache) length() int {
	return int(atomic.LoadInt32(&c.size))
}
func (c *Cache) delete(key string) {
	if _, ok := c.data.Load(key); ok {
		c.data.Delete(key)
		atomic.AddInt32(&c.size, -1)
		log.Println("删除", key, "长度", c.size)
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
	go cache.watch()
	cache.store("1", Data{value: "1111", time: time.Now()})
	time.Sleep(time.Second * 2)
	cache.store("2", Data{value: "2222", time: time.Now()})
	time.Sleep(time.Second * 2)
	cache.store("3", Data{value: "3333", time: time.Now()})
	time.Sleep(time.Second * 3)
	cache.store("4", Data{value: "4444", time: time.Now()})
	cache.store("4", Data{value: "4444", time: time.Now()})
	time.Sleep(time.Second)
	cache.store("5", Data{value: "5555", time: time.Now()})
	cache.store("5", Data{value: "5555", time: time.Now()})
	time.Sleep(time.Second * 5)
	cache.store("6", Data{value: "6666", time: time.Now()})
	time.Sleep(time.Second * 10)
}
