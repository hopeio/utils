/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package protobuf

import "google.golang.org/protobuf/proto"

func Unmarshal(body []byte, obj interface{}) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return Validate(obj), but util now we can't add
	// `binding:""` to the struct which automatically generate by gen-proto
	return nil
	// return Validate(obj)
}
