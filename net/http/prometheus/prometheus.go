/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var AccessCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_requests_total",
	},
	[]string{"method", "uri"},
)

var QueueGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "queue_num_total",
	},
	[]string{"method", "uri"},
)

var HttpDurationsHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_durations_histogram_millisecond",
		Buckets: []float64{30, 60, 100, 200, 300, 500, 1000},
	},
	[]string{"method", "uri"},
)
var HttpDurations = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "http_durations_millisecond",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"method", "uri"},
)

func init() {
	prometheus.MustRegister(AccessCounter, QueueGauge, HttpDurationsHistogram, HttpDurations)
}
