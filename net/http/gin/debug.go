/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/debug"
)

func Debug(r *gin.Engine) {
	r.Any("/debug/*path", Wrap(debug.Handler()))
}
