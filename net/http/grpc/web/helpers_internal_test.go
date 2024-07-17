package web

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGRPCEndpoint(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{input: "/", output: "/"},
		{input: "/resource", output: "/resource"},
		{input: "/test.TestService/PingEmpty", output: "/test.TestService/PingEmpty"},
		{input: "/test.TestService/PingEmpty/", output: "/test.TestService/PingEmpty"},
		{input: "/a/b/c/test.TestService/PingEmpty", output: "/test.TestService/PingEmpty"},
		{input: "/a/b/c/test.TestService/PingEmpty/", output: "/test.TestService/PingEmpty"},
	}

	for _, c := range cases {
		req := httptest.NewRequest("GET", c.input, nil)
		result := getGRPCEndpoint(req)

		assert.Equal(t, c.output, result)
	}
}
