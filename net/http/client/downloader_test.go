package client

import "testing"

func TestFetch(t *testing.T) {
	_, err := GetReader("")
	if err != nil {
		t.Log(err)
	}
}
