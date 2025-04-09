package binding

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

type User struct {
	ID    int    `uri:"id"`
	Name  string `json:"name"`
	Age   int    `header:"age"`
	Phone string `query:"phone"`
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
		assert.Equal(t, 1, u.ID)
		assert.Equal(t, "test", u.Name)
		done <- struct{}{}
	})
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(time.Second)
	req, err := http.NewRequest("POST", "http://localhost:8080/user/1?phone=123", bytes.NewBufferString(`{"name":"test"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Age", "16")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	<-done
}
