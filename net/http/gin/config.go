/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import "github.com/gin-gonic/gin"

type Config gin.Engine

func (c *Config) New() *gin.Engine {
	// 内部循环引用,直接反序列化地址会变
	engine := gin.New()
	if c != nil {
		engine.RedirectTrailingSlash = c.RedirectTrailingSlash
		engine.RedirectFixedPath = c.RedirectFixedPath
		engine.HandleMethodNotAllowed = c.HandleMethodNotAllowed
		engine.ForwardedByClientIP = c.ForwardedByClientIP
		engine.AppEngine = c.AppEngine
		engine.UseRawPath = c.UseRawPath
		engine.UnescapePathValues = c.UnescapePathValues
		engine.RemoveExtraSlash = c.RemoveExtraSlash
		engine.RemoteIPHeaders = c.RemoteIPHeaders
		engine.TrustedPlatform = c.TrustedPlatform
		engine.MaxMultipartMemory = c.MaxMultipartMemory
		engine.UseH2C = c.UseH2C
		engine.ContextWithFallback = c.ContextWithFallback
	}
	return engine
}
