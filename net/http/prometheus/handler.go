/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PromHandler() http.Handler {
	http.Handle("/metrics", promhttp.Handler())
	return http.DefaultServeMux
}
