package aop

import (
	"log"
	"testing"
)

var foo1 = func() {
	log.Println("foo1")
}

func foo2() {
	log.Println("foo1")
}
func before() { log.Println("before") }
func after()  { log.Println("after") }

func TestAop(t *testing.T) {
	Invoke(before, &foo1, after)
	foo1()

	log.Println("----------------------------------------")
	aop(before, foo2, after)
	foo2()
	
}
