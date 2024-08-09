package json

import (
	"github.com/hopeio/utils/log"
	"testing"
)

func TestUnquote(t *testing.T) {
	s := []byte(`"\u8bf7\u6c42\u8fc7\u4e8e\u9891\u7e41"`)
	log.Println(Unquote(s))
}
