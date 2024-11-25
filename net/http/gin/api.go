/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/fs"
	"github.com/hopeio/utils/net/http/apidoc"
	_ "github.com/ugorji/go"
)

func OpenApi(mux *gin.Engine, uriPrefix, dir string) {
	if dir != "" {
		if b := dir[len(dir)-1:]; b == "/" || b == "\\" {
			apidoc.Dir = dir
		} else {
			apidoc.Dir = dir + fs.PathSeparator
		}
	}
	if uriPrefix != "" {
		apidoc.UriPrefix = uriPrefix
	}

	mux.GET(apidoc.UriPrefix, Wrap(apidoc.DocList))
	mux.GET(apidoc.UriPrefix+"/markdown/*file", Wrap(apidoc.Markdown))
	mux.GET(apidoc.UriPrefix+"/swagger/*file", Wrap(apidoc.Swagger))
}
