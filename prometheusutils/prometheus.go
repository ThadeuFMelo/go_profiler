package prometheusutil

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Metrics
	SearchQueryCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "search_query_counter",
		Help: "The total number of search queries",
	}, []string{"query"})

	//CPU usage per process
	ProcessCPUUsage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "process_cpu_usage",
		Help: "CPU usage by the process",
	}, []string{"process_name"})

	//Memory usage per process
	ProcessMemoryUsage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "process_memory_usage",
		Help: "Memory usage by the process",
	}, []string{"process_name"})
)

// Register and starts the prometheus
func Register(endpoint string) {
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(endpoint, nil) // nolint:errcheck
	// Initialize the counter on start so that it gets picked up for availability reporting.
	SearchQueryCount.WithLabelValues("").Add(0)
	ProcessCPUUsage.WithLabelValues("").Add(0)
	ProcessMemoryUsage.WithLabelValues("").Add(0)
}
