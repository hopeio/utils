package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestConvert(t *testing.T) {
	t.Log(StringConverterArrays)
}

func TestSizeof(t *testing.T) {
	t.Log(unsafe.Sizeof(1))
}

func TestStringConvertBasicFor(t *testing.T) {

	t.Run("int8", func(t *testing.T) {
		got, err := StringConvertBasicFor[int8]("123")
		assert.Nil(t, err)
		assert.Equal(t, int8(123), got)
	})
	t.Run("int", func(t *testing.T) {
		got, err := StringConvertBasicFor[int]("123456789")
		assert.Nil(t, err)
		assert.Equal(t, 123456789, got)
	})
	t.Run("uint", func(t *testing.T) {
		got, err := StringConvertBasicFor[uint]("123456789")
		assert.Nil(t, err)
		assert.Equal(t, uint(123456789), got)
	})
	t.Run("bool", func(t *testing.T) {
		got, err := StringConvertBasicFor[bool]("1")
		assert.Nil(t, err)
		assert.Equal(t, true, got)
	})
	t.Run("float32", func(t *testing.T) {
		got, err := StringConvertBasicFor[float32]("1.23")
		assert.Nil(t, err)
		assert.Equal(t, float32(1.23), got)
	})
}
