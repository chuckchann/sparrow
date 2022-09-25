package monitor

import (
	"log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"net/http"
)

func PrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	promAddr := ":" + viper.GetString("prometheus.port")
	log.Fatal(http.ListenAndServe(promAddr, nil))
}
