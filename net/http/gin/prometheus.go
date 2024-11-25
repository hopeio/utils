/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Prom(r *gin.Engine) {
	// Register Metrics metrics handler.
	r.Any("/metrics", Wrap(promhttp.Handler()))
}
