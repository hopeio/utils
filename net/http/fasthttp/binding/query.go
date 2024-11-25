/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	"github.com/valyala/fasthttp"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	values := req.URI().QueryArgs()
	if err := mtos.MapFormByTag(obj, (*ArgsSource)(values), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
