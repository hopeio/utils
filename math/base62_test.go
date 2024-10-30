package math

import (
	"testing"
)

func TestConvInt(t *testing.T) {
	t.Log(FormatUint(5102198557, 62))
	t.Log(ParseUint("5ziiV7", 62, 64))
	t.Log(FormatInt(-5102198557, 62))
	t.Log(ParseInt("-5ziiV7", 62, 64))
	t.Log(ParseUint("1gk7tnzw", 36, 64))
	t.Log(ParseUint("gk7tnzw", 36, 64))
	t.Log(ParseUint("j53344mo7wk2", 36, 64))
	t.Log(FormatUint(4389580, 36))
}
