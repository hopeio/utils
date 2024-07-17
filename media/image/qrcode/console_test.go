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
