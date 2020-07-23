package runtime

import (
	"runtime"
	"testing"
)

func TestReadMemStats(t *testing.T) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	t.Log("NumGC", stats.NumGC)
	t.Log("TotalAlloc", stats.TotalAlloc/1024)
	t.Log("Alloc", stats.Alloc/1024)
	t.Log("Sys", stats.Sys/1024)
	t.Log("HeapReleased", stats.HeapReleased/1024)
	t.Log("BuckHashSys", stats.BuckHashSys)
	t.Log("Lookups", stats.Lookups)
	t.Log("Mallocs", stats.Mallocs)
	t.Log("Frees", stats.Frees)
	t.Log("HeapAlloc", stats.HeapAlloc)
	t.Log("HeapSys", stats.HeapSys)
	t.Log("HeapIdle", stats.HeapIdle)
	t.Log("HeapInuse", stats.HeapInuse)
	t.Log("HeapReleased", stats.HeapReleased)
	t.Log("HeapObjects", stats.HeapObjects)
	t.Log("StackInuse", stats.StackInuse)
	t.Log("StackSys", stats.StackSys)

	t.Log("PauseTotalNs", stats.PauseTotalNs)
	t.Log("GCCPUFraction", stats.GCCPUFraction)
	t.Log("EnableGC", stats.EnableGC)
}
