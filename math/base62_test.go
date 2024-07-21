package math

import (
	"fmt"
	"testing"
)

func TestConvInt(t *testing.T) {
	t.Log(FormatUint(5102198557, 62))
	t.Log(ParseUint("1gk7tnzw", 36))
	t.Log(ParseUint("gk7tnzw", 36))
	t.Log(ParseUint("j53344mo7wk2", 36))
	t.Log(FormatUint(4389580, 36))
}

func TestConv(t *testing.T) {
	fmt.Println(ToBytes(333))
}
