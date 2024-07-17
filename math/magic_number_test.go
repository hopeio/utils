package math

import (
	"testing"
	"time"
)

func TestMagicNumber(t *testing.T) {
	key := SecondKey()
	t.Log(key)
	t.Log(ValidateSecondKey(key))
	t.Log(ValidateSecondKey(1 ^ magicNumber))
	t.Log(ValidateSecondKey(2 ^ magicNumber))
	t.Log(ValidateSecondKey(3 ^ magicNumber))
	t.Log(ValidateSecondKey(time.Now().Unix() - 1 ^ magicNumber))
}
