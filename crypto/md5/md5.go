/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func EncodeString(value string) string {
	md5 := md5.Sum([]byte(value))
	return hex.EncodeToString(md5[:])
}

func Encode(value string) []byte {
	md5 := md5.Sum([]byte(value))
	return md5[:]
}

func ToString(md5 []byte) string {
	return hex.EncodeToString(md5)
}

func EncodeReader(r io.Reader) ([]byte, error) {
	hash := md5.New()
	_, err := io.Copy(hash, r)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

func EncodeReaderString(r io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, r)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
