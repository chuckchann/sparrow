package monitor

import "github.com/prometheus/client_golang/prometheus"

var metricsRequest *prometheus.CounterVec
var metricsLatency *prometheus.SummaryVec

func Init() {
	//for recording request info
	metricsRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_info",
			Help: "The total request number of this service_mng",
		}, []string{"namespace", "app", "protocol", "method", "resp_code"})

	//for recording request latency
	metricsLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "request_latency",
			Help:       "the latency of this service",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"namespace", "app", "protocol", "method", "resp_code"})

	prometheus.MustRegister(metricsRequest)
	prometheus.MustRegister(metricsLatency)
}

func GetRequestMetrics() *prometheus.CounterVec {
	return metricsRequest
}

func GetLatencyMetrics() *prometheus.SummaryVec {
	return metricsLatency
}
