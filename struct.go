package appmonitor

import (
	"github.com/wgliang/metrics"
)

// 运行生产环境下监控的应用指标
type RuntimeMetrics struct {
	MemStats struct {
		Alloc         metrics.Gauge
		BuckHashSys   metrics.Gauge
		DebugGC       metrics.Gauge
		EnableGC      metrics.Gauge
		Frees         metrics.Gauge
		HeapAlloc     metrics.Gauge
		HeapIdle      metrics.Gauge
		HeapInuse     metrics.Gauge
		HeapObjects   metrics.Gauge
		HeapReleased  metrics.Gauge
		HeapSys       metrics.Gauge
		LastGC        metrics.Gauge
		Lookups       metrics.Gauge
		Mallocs       metrics.Gauge
		MCacheInuse   metrics.Gauge
		MCacheSys     metrics.Gauge
		MSpanInuse    metrics.Gauge
		MSpanSys      metrics.Gauge
		NextGC        metrics.Gauge
		NumGC         metrics.Gauge
		GCCPUFraction metrics.GaugeFloat64
		PauseNs       metrics.Histogram
		PauseTotalNs  metrics.Gauge
		StackInuse    metrics.Gauge
		StackSys      metrics.Gauge
		Sys           metrics.Gauge
		TotalAlloc    metrics.Gauge
	}
	NumCgoCall   metrics.Gauge
	NumGoroutine metrics.Gauge
	ReadMemStats metrics.Histogram
}

// 调试模式下的应用指标
type DebugMetrics struct {
	GCStats struct {
		LastGC     metrics.Gauge
		NumGC      metrics.Gauge
		Pause      metrics.Histogram
		PauseTotal metrics.Gauge
	}
	ReadGCStats metrics.Histogram
}