package metrics

import "runtime"

// Register metrics for the Go runtime statistics exported in
// runtime.MemStats.  The metrics are named by their fully-qualified
// Go symbols, i.e. runtime.MemStatsAlloc.  In addition to
// runtime.MemStats, register the return value of runtime.Goroutines()
// as runtime.Goroutines.
func RegisterRuntimeMemStats(r Registry) {

	r.Register("runtime.MemStats.Alloc", NewGauge())
	r.Register("runtime.MemStats.TotalAlloc", NewGauge())
	r.Register("runtime.MemStats.Sys", NewGauge())
	r.Register("runtime.MemStats.Lookups", NewGauge())
	r.Register("runtime.MemStats.Mallocs", NewGauge())
	r.Register("runtime.MemStats.Frees", NewGauge())

	r.Register("runtime.MemStats.HeapAlloc", NewGauge())
	r.Register("runtime.MemStats.HeapSys", NewGauge())
	r.Register("runtime.MemStats.HeapIdle", NewGauge())
	r.Register("runtime.MemStats.HeapInuse", NewGauge())
	r.Register("runtime.MemStats.HeapObjects", NewGauge())

	r.Register("runtime.MemStats.StackInuse", NewGauge())
	r.Register("runtime.MemStats.StackSys", NewGauge())
	r.Register("runtime.MemStats.MSpanInuse", NewGauge())
	r.Register("runtime.MemStats.MSpanSys", NewGauge())
	r.Register("runtime.MemStats.MCacheInuse", NewGauge())
	r.Register("runtime.MemStats.MCacheSys", NewGauge())
	r.Register("runtime.MemStats.BuckHashSys", NewGauge())

	r.Register("runtime.MemStats.NextGC", NewGauge())
	r.Register("runtime.MemStats.PauseTotalNs", NewGauge())
	r.Register("runtime.MemStats.PauseNs",
		NewHistogram(NewExpDecaySample(1028, 0.015)))
	r.Register("runtime.MemStats.NumGC", NewGauge())
	r.Register("runtime.MemStats.EnableGC", NewGauge())
	r.Register("runtime.MemStats.DebugGC", NewGauge())

	r.Register("runtime.NumCgoCall", NewGauge())
	r.Register("runtime.NumGoroutine", NewGauge())

}

// Capture new values for the Go runtime statistics exported in
// runtime.MemStats.  This is designed to be called in a background
// goroutine.  Giving a registry which has not been given to
// RegisterRuntimeMemStats will panic.  If the second parameter is
// false, the counters will be left to the lazy updates provided by
// the runtime.
func CaptureRuntimeMemStats(r Registry, readMemStats bool) {
	var m runtime.MemStats
	if readMemStats {
		runtime.ReadMemStats(&m)
	}

	r.Get("runtime.MemStats.Alloc").(Gauge).Update(int64(m.Alloc))
	r.Get("runtime.MemStats.TotalAlloc").(Gauge).Update(int64(m.TotalAlloc))
	r.Get("runtime.MemStats.Sys").(Gauge).Update(int64(m.Sys))
	r.Get("runtime.MemStats.Lookups").(Gauge).Update(int64(m.Lookups))
	r.Get("runtime.MemStats.Mallocs").(Gauge).Update(int64(m.Mallocs))
	r.Get("runtime.MemStats.Frees").(Gauge).Update(int64(m.Frees))

	r.Get("runtime.MemStats.HeapAlloc").(Gauge).Update(int64(m.HeapAlloc))
	r.Get("runtime.MemStats.HeapSys").(Gauge).Update(int64(m.HeapSys))
	r.Get("runtime.MemStats.HeapIdle").(Gauge).Update(int64(m.HeapIdle))
	r.Get("runtime.MemStats.HeapInuse").(Gauge).Update(int64(m.HeapInuse))
	r.Get("runtime.MemStats.HeapObjects").(Gauge).Update(int64(m.HeapObjects))

	r.Get("runtime.MemStats.StackInuse").(Gauge).Update(int64(m.StackInuse))
	r.Get("runtime.MemStats.StackSys").(Gauge).Update(int64(m.StackSys))
	r.Get("runtime.MemStats.MSpanInuse").(Gauge).Update(int64(m.MSpanInuse))
	r.Get("runtime.MemStats.MSpanSys").(Gauge).Update(int64(m.MSpanSys))
	r.Get("runtime.MemStats.MCacheInuse").(Gauge).Update(int64(m.MCacheInuse))
	r.Get("runtime.MemStats.MCacheSys").(Gauge).Update(int64(m.MCacheSys))
	r.Get("runtime.MemStats.BuckHashSys").(Gauge).Update(int64(m.BuckHashSys))

	r.Get("runtime.MemStats.NextGC").(Gauge).Update(int64(m.NextGC))
	r.Get("runtime.MemStats.PauseTotalNs").(Gauge).Update(int64(m.PauseTotalNs))
	r.Get("runtime.MemStats.PauseNs").(Histogram).Update(int64(m.PauseNs[0]))
	r.Get("runtime.MemStats.NumGC").(Gauge).Update(int64(m.NumGC))
	if m.EnableGC {
		r.Get("runtime.MemStats.EnableGC").(Gauge).Update(1)
	} else {
		r.Get("runtime.MemStats.EnableGC").(Gauge).Update(0)
	}
	if m.EnableGC {
		r.Get("runtime.MemStats.DebugGC").(Gauge).Update(1)
	} else {
		r.Get("runtime.MemStats.DebugGC").(Gauge).Update(0)
	}

	r.Get("runtime.NumCgoCall").(Gauge).Update(int64(runtime.NumCgoCall()))
	r.Get("runtime.NumGoroutine").(Gauge).Update(int64(runtime.NumGoroutine()))

}
