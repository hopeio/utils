/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package validator

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/hopeio/gox/log"
)

var (
	Validator *validator.Validate
	trans     ut.Translator
)

func init() {
	zhcn := zh.New()
	uni := ut.New(zhcn)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("zh")

	Validator = validator.New()
	zh_translations.RegisterDefaultTranslations(Validator, trans)
	Validator.RegisterTagNameFunc(func(sf reflect.StructField) string {
		if comment := sf.Tag.Get("comment"); comment != "" {
			return comment
		}
		if json := sf.Tag.Get("json"); json != "" {
			return json
		}
		return sf.Name
	})
	Validator.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(phonePattern, fl.Field().String())
		return match
	})
	Validator.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "{0}必须是一个有效的手机号!", true)
	}, translateFunc)
}

func TransError(err error) string {
	if err == nil {
		return ""
	}
	var msg []string
	var ve validator.ValidationErrors
	ok := errors.As(err, &ve)
	if !ok {
		return err.Error()
	}
	for _, v := range ve.Translate(trans) {
		msg = append(msg, v)
	}
	return strings.Join(msg, ",")
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Error("translate err:", fe)
		return fe.(error).Error()
	}

	return t
}
