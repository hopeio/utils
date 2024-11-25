/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	url2 "github.com/hopeio/utils/net/url"
	"net/url"
	"testing"
)

func TestResolveURL(t *testing.T) {
	testURL := "http://www.example.com/test/index.m3m8"
	u, err := url.Parse(testURL)
	if err != nil {
		t.Error(err)
	}

	result := url2.ResolveURL(u, "videos/111111.ts")
	expected := "http://www.example.com/test/videos/111111.ts"
	if result != expected {
		t.Fatalf("wrong URL, expected: %s, result: %s", expected, result)
	}

	result = url2.ResolveURL(u, "/videos/2222222.ts")
	expected = "http://www.example.com/videos/2222222.ts"
	if result != expected {
		t.Fatalf("wrong URL, expected: %s, result: %s", expected, result)
	}

	result = url2.ResolveURL(u, "https://test.com/11111.key")
	expected = "https://test.com/11111.key"
	if result != expected {
		t.Fatalf("wrong URL, expected: %s, result: %s", expected, result)
	}
}
