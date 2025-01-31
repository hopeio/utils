/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package aes

import (
	"testing"
)

func Test_AES128Encrypt_AND_AES128Decrypt(t *testing.T) {
	expected := "helloworld"
	key := "8dv4byf8b9e6bc1x"
	iv := "xduio1f8a12348u4"
	encrypt, err := CBCEncrypt([]byte(expected), []byte(key), []byte(iv))
	if err != nil {
		t.Fatal(err)
	}
	decrypt, err := CBCDecrypt(encrypt, []byte(key), []byte(iv))
	if err != nil {
		t.Fatal(err)
	}
	de := string(decrypt)
	if de != expected {
		t.Fatalf("expected: %s, result: %s", expected, de)
	}
}
