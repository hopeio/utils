package reflect

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	var a int
	t.Log(reflect.ValueOf(a).IsValid())
}

func TestInitPtr(t *testing.T) {
	var a *int
	var b = 1
	a = &b
	reflect.ValueOf(a).Elem().Set(reflect.ValueOf(2))
	assert.Equal(t, 2, *a)
	var c *int
	var d = 2
	reflect.ValueOf(&c).Elem().Set(reflect.ValueOf(&d))
	assert.Equal(t, 2, *c)

	var e *int
	v := InitPtr(reflect.ValueOf(&e))
	v.Set(reflect.ValueOf(1))
	assert.Equal(t, 1, *e)

	var f int
	v = InitPtr(reflect.ValueOf(&f))
	v.Set(reflect.ValueOf(3))
	assert.Equal(t, 3, f)

	var g ********int
	v = InitPtr(reflect.ValueOf(&g))
	v.Set(reflect.ValueOf(3))
	assert.Equal(t, 3, ********g)
}

func TestDerefInterfaceValue(t *testing.T) {
	var a any
	v := reflect.ValueOf(&a)
	t.Log(v.Kind())
	v1 := v.Elem()
	t.Log(v1.Kind())
	v2 := v1.Elem()
	t.Log(v2.Kind())
	a = 1
	v = reflect.ValueOf(&a)
	t.Log(v.Kind())
	v1 = v.Elem()
	t.Log(v1.Kind())
	t.Log(v1.Type().Kind())
	v2 = v1.Elem()
	t.Log(v2.Kind())
	t.Log(v2.Type().Kind())

	var b any
	t.Log(DerefValue(reflect.ValueOf(&b)).Kind())
	b = 1
	t.Log(DerefValue(reflect.ValueOf(&b)).Kind())
}
