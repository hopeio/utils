/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package middleware

import (
	log2 "github.com/hopeio/utils/net/http/log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/log"
)

func SetLog(app *gin.Engine, logger2 *log.Logger, errHandle bool) {
	app.Use(LoggerWithFormatter(log2.DefaultLogFormatter, logger2, errHandle))

}

// ErrorLogger returns a handlerfunc for any error type.
func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

// ErrorLoggerT returns a handlerfunc for a given error type.
func ErrorLoggerT(typ gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.ByType(typ)
		if len(errors) > 0 {
			c.JSON(-1, errors)
		}
	}
}

// Logger instances a Logger middleware that will write the logs to gin.DefaultWriter.
// By default gin.DefaultWriter = os.Stdout.
func Logger() gin.HandlerFunc {
	return LoggerWithConfig(log2.LoggerConfig{})
}

// LoggerWithFormatter instance a Logger middleware with the specified log format function.
func LoggerWithFormatter(f log2.Formatter, logger *log.Logger, hasErr bool) gin.HandlerFunc {
	return LoggerWithConfig(log2.LoggerConfig{
		Formatter: f,
		Logger:    logger,
		ErrHandle: hasErr,
	})
}

// LoggerWithConfig instance a Logger middleware with config.
func LoggerWithConfig(conf log2.LoggerConfig) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = log2.DefaultLogFormatter
	}

	logger := conf.Logger
	if logger == nil {
		logger = log.Default()
	}

	notlogged := conf.SkipPaths

	errHandle := conf.ErrHandle

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// ErrorLog only when path is not being skipped
		for _, ext := range notlogged {
			if strings.HasSuffix(path, ext) {
				return
			}
		}
		param := log2.FormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		if errHandle {
			logger.Warn(formatter(param))
		} else {
			logger.Info(formatter(param))
		}
	}
}
