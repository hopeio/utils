/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package reconn

import (
	"reflect"
	"unsafe"

	"google.golang.org/grpc"
)

// 这里的key值应该取什么
var ReConnectMap = make(map[string]func() error)

func ReConnect(v interface{}, getEndPort func() string, opts []grpc.DialOption) func() error {
	value := reflect.ValueOf(v).Elem()
	ptr := value.Field(0).Pointer()
	conn := (*grpc.ClientConn)(unsafe.Pointer(ptr))
	return func() error {
		conn.Close()
		endPort := getEndPort()
		newConn, err := grpc.NewClient(endPort, opts...)
		if err != nil {
			return err
		}
		newConnPtr := (**grpc.ClientConn)(unsafe.Pointer(value.Addr().Pointer()))
		*newConnPtr = newConn
		return nil
	}
}
