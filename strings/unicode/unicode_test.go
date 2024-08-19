package unicode

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUnquote(t *testing.T) {
	var s = []byte(`{"ok":0,"errno":"100005","msg":"\u8bf7\u6c42\u8fc7\u4e8e\u9891\u7e41"}`)
	log.Println(string(s))
	log.Println(ToUtf8(s))
	s = []byte(`\u8bf7\u6c42\u8fc7\u4e8e\u9891\u7e41`)
	log.Println(ToUtf8(s))
	//log.Println(Unquote(s))
}

func TestTrimSymbol(t *testing.T) {
	assert.Equal(t, "Hello世界123", TrimSymbol("Hello, 世界! 123"))
	assert.Equal(t, "Hello世界123", TrimSymbol("Hello, 世界! 😊 123"))
	assert.Equal(t, "Hello, 世界!  123", TrimEmoji("Hello, 世界! 😊 123"))
	assert.Equal(t, "Hello世界123", TrimSymbol("Hello_世界_123"))
	assert.Equal(t, "是谁的小篮球", TrimSymbol("是谁的小篮球🏀？"))
	assert.Equal(t, "汉字567", RetainChineseAndAlphanumeric("૮𖥦აʚɞ汉字567"))
}
