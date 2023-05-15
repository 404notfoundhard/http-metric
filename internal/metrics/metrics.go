package metrics

import (
	"math/rand"
	"runtime"
)

type Metrics struct {
	GCCPUFraction float64
	Alloc         uint64
	BuckHashSys   uint64
	Frees         uint64
	GCSys         uint64
	HeapAlloc     uint64
	HeapIdle      uint64
	HeapInuse     uint64
	HeapObjects   uint64
	HeapReleased  uint64
	HeapSys       uint64
	LastGC        uint64
	Lookups       uint64
	MCacheInuse   uint64
	MCacheSys     uint64
	MSpanInuse    uint64
	MSpanSys      uint64
	Mallocs       uint64
	NextGC        uint64
	NumForcedGC   uint32
	NumGC         uint32
	OtherSys      uint64
	PauseTotalNs  uint64
	StackInuse    uint64
	StackSys      uint64
	Sys           uint64
	TotalAlloc    uint64
	PollCount     uint64
	RandomValue   uint32
}

func (m Metrics) ReadMetrics() Metrics {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	return Metrics{
		GCCPUFraction: rtm.GCCPUFraction,
		Alloc:         rtm.Alloc,
		BuckHashSys:   rtm.BuckHashSys,
		Frees:         rtm.Frees,
		GCSys:         rtm.GCSys,
		HeapAlloc:     rtm.HeapAlloc,
		HeapIdle:      rtm.HeapIdle,
		HeapInuse:     rtm.HeapInuse,
		HeapObjects:   rtm.HeapObjects,
		HeapReleased:  rtm.HeapReleased,
		HeapSys:       rtm.HeapSys,
		LastGC:        rtm.LastGC,
		Lookups:       rtm.Lookups,
		MCacheInuse:   rtm.MCacheInuse,
		MCacheSys:     rtm.MCacheSys,
		MSpanInuse:    rtm.MSpanInuse,
		MSpanSys:      rtm.MSpanSys,
		Mallocs:       rtm.Mallocs,
		NextGC:        rtm.NextGC,
		NumForcedGC:   rtm.NumForcedGC,
		NumGC:         rtm.NumGC,
		OtherSys:      rtm.OtherSys,
		PauseTotalNs:  rtm.PauseTotalNs,
		StackInuse:    rtm.StackInuse,
		StackSys:      rtm.StackSys,
		Sys:           rtm.Sys,
		TotalAlloc:    rtm.TotalAlloc,
		RandomValue:   uint32(rand.Int31()),
	}

}
