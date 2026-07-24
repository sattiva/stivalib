package sys

import (
	"runtime"
	"time"
)

type Metrics struct {
	Goroutines int
	AllocBytes uint64
	TotalAlloc uint64
	SysBytes   uint64
	NumGC      uint32
}

type Watchdog struct {
	MaxGoroutines int
	MaxMemoryMB   uint64
	Interval      time.Duration
	OnExceed      func(Metrics)
}

func NewWatchdog(maxGoroutines int, maxMemoryMB uint64, interval time.Duration, onExceed func(Metrics)) *Watchdog {
	w := &Watchdog{
		MaxGoroutines: maxGoroutines,
		MaxMemoryMB:   maxMemoryMB,
		Interval:      interval,
		OnExceed:      onExceed,
	}
	go w.run()
	return w
}

func GetMetrics() Metrics {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return Metrics{
		Goroutines: runtime.NumGoroutine(),
		AllocBytes: m.Alloc,
		TotalAlloc: m.TotalAlloc,
		SysBytes:   m.Sys,
		NumGC:      m.NumGC,
	}
}

func (w *Watchdog) run() {
	for range time.Tick(w.Interval) {
		m := GetMetrics()
		memMB := m.AllocBytes / (1024 * 1024)
		if (w.MaxGoroutines > 0 && m.Goroutines > w.MaxGoroutines) || (w.MaxMemoryMB > 0 && memMB > w.MaxMemoryMB) {
			if w.OnExceed != nil {
				w.OnExceed(m)
			}
		}
	}
}
