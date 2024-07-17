package math

import (
	"fmt"
	"testing"
)

func TestConvInt(t *testing.T) {
	t.Log(FormatInt(5102198557, 62))
	t.Log(ParseInt("1gk7tnzw", 36))
	t.Log(ParseInt("gk7tnzw", 36))
	t.Log(ParseInt("j53344mo7wk2", 36))
	t.Log(FormatInt(4389580, 36))
}

func TestConv(t *testing.T) {
	fmt.Println(ToBytes(333))
}
