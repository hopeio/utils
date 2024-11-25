/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package qrcode

import (
	"github.com/boombuler/barcode/qr"
	"testing"
)

func TestName(t *testing.T) {
	qrcode, err := qr.Encode("hello world", qr.H, qr.Unicode)
	if err != nil {
		t.Error(err)
	}
	ConsolePrint(qrcode)
}
