package race

import (
	"flag"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
)

const (
	RaceNoHandle = iota
	RaceLock
	RaceAtomic
)

type DataRace struct {
	value1 int
	value2 int32
	Arr    []byte
	Type   byte
	sync.WaitGroup
	sync.RWMutex
}

func (d *DataRace) Run() {
	log.Println("test start....")
	flag.Parse()
	t, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		t = RaceNoHandle
	}
	switch byte(t) {
	case RaceNoHandle:
		log.Println("RaceNoHandle...")
	case RaceLock:
		log.Println("RaceLock...")
	case RaceAtomic:
		log.Println("RaceAtomic...")
	default:
		panic("invalid handle")
	}
	d.Type = byte(t)
	for i := 0; i < 10; i++ {
		d.Add(1)
		go d.Do()
	}
	d.Wait()
	switch byte(t) {
	case RaceNoHandle:
		log.Println("RaceNoHandle...")
		log.Println(d.value1)
	case RaceLock:
		log.Println(d.value1)
	case RaceAtomic:
		log.Println(d.value2)
	default:
		panic("invalid handle")
	}
}

func (d *DataRace) Do() {
	for i := 0; i < 10000; i++ {
		d.Counter()
	}
	d.Done()
}

func (d *DataRace) Value() int {
	d.Lock()
	defer d.Unlock()
	return d.value1
}

func (d *DataRace) Counter() {
	switch d.Type {
	case RaceLock:
		d.Lock()
		defer d.Unlock()
		d.value1 += 1
	case RaceNoHandle:
		d.value1++
	case RaceAtomic:
		atomic.AddInt32(&d.value2, 1)
	default:
		panic("invalid handle")
	}
}
