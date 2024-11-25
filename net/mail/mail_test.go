/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package mail

import (
	"fmt"
	"testing"
)

func TestGenMsg(t *testing.T) {
	var msg = Mail{
		FromName: "liov",
		From:     "liov@github.com",
		Subject:  "测试",
		Content:  "邮件",
		To:       []string{"test1@mail.com", "test2@mail.com"},
	}
	bytes, _ := msg.GenMsg()
	fmt.Println(string(bytes))
}
