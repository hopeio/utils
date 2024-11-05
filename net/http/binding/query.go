// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/hopeio/utils/encoding"
	"net/http"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	if err := encoding.MapFormByTag(obj, encoding.KVsSource(values), Tag); err != nil {
		return err
	}
	return Validate(obj)
}
