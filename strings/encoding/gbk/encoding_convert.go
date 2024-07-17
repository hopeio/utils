package gbk

import (
	"bytes"
	strings2 "github.com/hopeio/utils/strings"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

// GBK 转 UTF-8
func GBKToUTF8(s string) (string, error) {
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	b, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return strings2.ToString(b), nil
}

func GBKBytesToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(reader)
}

// UTF-8 转 GBK

func UTF8ToGBK(s string) (string, error) {
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	b, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return strings2.ToString(b), nil
}

func UTF8BytesToGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(reader)
}
