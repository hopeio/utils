package strings

import (
	"fmt"
	"testing"
)

func TestStrconv(t *testing.T) {
	s := "test"
	b := StringToBytes(s)
	s2 := BytesToString(b)
	fmt.Println(b, s2)
}
