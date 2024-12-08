/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package apidoc

import (
	"github.com/pb33f/libopenapi"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hopeio/utils/log"
)

var Doc *v3high.Document

// 参数为路径和格式
func GetDoc(realPath, modName string) *v3high.Document {
	if Doc != nil {
		return Doc
	}
	if realPath == "" {
		realPath = "."
	}

	realPath = realPath + modName
	err := os.MkdirAll(realPath, os.ModePerm)
	if err != nil {
		log.Error(err)
	}

	realPath = filepath.Join(realPath, modName+SwaggerEXT)

	if _, err := os.Stat(realPath); os.IsNotExist(err) {
		return generate()
	} else {
		file, err := os.Open(realPath)
		if err != nil {
			log.Error(err)
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Error(err)
		}
		doc, _ := libopenapi.NewDocument(data)
		model, _ := doc.BuildV3Model()
		Doc = &model.Model
	}
	return Doc
}

func generate() *v3high.Document {
	doc, err := libopenapi.NewDocumentWithTypeCheck([]byte(`{"type":"openapi","version":"3.0"}`), true)
	if err != nil {
		log.Error(err)
	}
	model, errs := doc.BuildV3Model()
	if err != nil {
		log.Error(errs)
	}
	Doc = &model.Model
	Doc.Paths = &v3high.Paths{}
	Doc.Components = &v3high.Components{}
	Doc.Servers = append(Doc.Servers, &v3high.Server{
		URL: "http://localhost:80",
	})
	return Doc
}

func WriteToFile(realPath, modName string) {
	if Doc == nil {
		return
	}
	if realPath == "" {
		realPath = "."
	}

	realPath = realPath + modName
	err := os.MkdirAll(realPath, os.ModePerm)
	if err != nil {
		log.Error(err)
	}

	realPath = filepath.Join(realPath, modName+SwaggerEXT)

	if _, err := os.Stat(realPath); err == nil {
		os.Remove(realPath)
	}
	data, err := Doc.RenderJSON("")
	if err != nil {
		log.Error(err)
		return
	}
	err = os.WriteFile(realPath, data, 0666)
	if err != nil {
		log.Error(err)
		return
	}

	/*b, err := yaml.Marshal(swag.ToDynamicJSON(Doc))
	  if err != nil {
	  	log.Error(err)
	  }
	  if _, err := file.Write(b); err != nil {
	  	log.Error(err)
	  }*/

	Doc = nil
}

func NilDoc() {
	Doc = nil
}
