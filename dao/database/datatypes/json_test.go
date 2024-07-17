package datatypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Foo struct {
	A int
	B string
}

func TestJSONArrayT(t *testing.T) {
	var jat JsonTArray[Foo]
	err := jat.Scan([]byte(`[{"A":1,"B":"1"},{"A":2,"B":"2"}]`))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, JsonTArray[Foo]{{1, "1"}, {2, "2"}}, jat)

}
