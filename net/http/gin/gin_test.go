/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRoute(t *testing.T) {
	i := gin.New()
	i.GET("/:id/:name/:path", func(context *gin.Context) { context.Writer.WriteString("/:id/:name/:path") })
	i.GET("/id/name/path", func(context *gin.Context) { context.Writer.WriteString("/id/name/path") })
	i.Run(":8080")
}
