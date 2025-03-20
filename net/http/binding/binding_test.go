package binding

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestBind(t *testing.T) {
	done := make(chan struct{}, 1)
	http.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		var u User
		err := Bind(r, &u)
		if err != nil {
			t.Error(err)
		}
		t.Log(u)
		assert.Equal(t, u.ID, 1)
		done <- struct{}{}
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(time.Second)
	_, err := http.Get("http://localhost:8080/user/1")
	if err != nil {
		t.Fatal(err)
	}
	<-done
}
