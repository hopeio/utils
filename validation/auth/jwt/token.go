/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package jwti

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	Parser          = jwt.NewParser()
	ErrInvalidToken = errors.New("invalid token")
)

func SetOptions(options ...jwt.ParserOption) {
	for _, option := range options {
		option(Parser)
	}
}

// 如果只存一个id，jwt的意义在哪呢，跟session_id有什么区别
// jwt应该存放一些用户不能更改的信息，所以不能全存在jwt里
// 或者说用户每更改一次信息就刷新token（貌似可行）
// 有泛型这里多好写
type Claims[T any] struct {
	Auth T `json:"auth,omitempty"`
	jwt.RegisteredClaims
}

func (c *Claims[T]) GenerateToken(secret interface{}) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(secret)
}

func NewClaims[T any](data T, maxAge int64, sign string) *Claims[T] {
	now := time.Now()
	exp := now.Add(time.Duration(maxAge))
	return &Claims[T]{
		Auth: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: exp},
			IssuedAt:  &jwt.NumericDate{Time: now},
			Issuer:    sign,
		},
	}
}

func GenerateToken(claims jwt.Claims, secret interface{}) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

func ParseToken(claims jwt.Claims, token string, secret []byte) (*jwt.Token, error) {
	return Parser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}

func ParseTokenWithKeyFunc(claims jwt.Claims, token string, f jwt.Keyfunc) (*jwt.Token, error) {
	return Parser.ParseWithClaims(token, claims, f)
}
