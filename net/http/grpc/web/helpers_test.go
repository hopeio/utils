//Copyright 2017 Improbable. All Rights Reserved.
// See LICENSE for licensing terms.

package web

import (
	"sort"
	"testing"

	testproto "github.com/hopeio/gox/net/http/grpc/web/test"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestListGRPCResources(t *testing.T) {
	server := grpc.NewServer()
	testproto.RegisterTestServiceServer(server, &testServiceImpl{})
	expected := []string{
		"/test.TestService/PingEmpty",
		"/test.TestService/Ping",
		"/test.TestService/PingError",
		"/test.TestService/PingList",
		"/test.TestService/Echo",
		"/test.TestService/PingPongBidi",
		"/test.TestService/PingStream",
	}
	actual := ListGRPCResources(server)
	sort.Strings(expected)
	sort.Strings(actual)
	assert.EqualValues(t,
		expected,
		actual,
		"list grpc resources must provide an exhaustive list of all registered handlers")
}
