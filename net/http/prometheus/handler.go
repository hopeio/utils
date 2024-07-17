package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PromHandler() http.Handler {
	http.Handle("/metrics", promhttp.Handler())
	return http.DefaultServeMux
}
