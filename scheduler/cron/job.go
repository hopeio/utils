/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package cron

import (
	"github.com/hopeio/gox/log"
)

type RedisTo struct {
}

func (RedisTo) Run() {
	log.Info("定时任务执行")
}
