/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobot(t *testing.T) {
	assert.Equal(t, `{"msgtype":"markdown","markdown":{"title":"aaa","text":"bbb"}}`, Format(&Markdown{
		Title: "aaa",
		Text:  "bbb",
		At:    nil,
	}))
	RobotSendTextMessageWithSecret("xx", "xx", "hello world")
	RobotSendMarkDownMessageWithSecret("xxx", "xx", "xxx", "hello world", nil)
}
