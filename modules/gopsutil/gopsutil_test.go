package gopsutil_test

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

const (
	K = 1024
	M = K * 1024
	T = M * 1024
)

func TestName(t *testing.T) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(partitions)
	usage, err := disk.Usage(partitions[0].Mountpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(usage.Total / T)
	log.Println(usage.Free / T)
	log.Println(usage.Free * 100 / usage.Total)

}

func TestVirtualMemory(t *testing.T) {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}

func TestAbs(t *testing.T) {
	abs, err := filepath.Abs(".")
	log.Println(abs, err)
}
