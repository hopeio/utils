package io

import (
	"bufio"
	"io"
)

func ReadReadCloser(readCloser io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(readCloser)
	if err != nil {
		return nil, err
	}
	readCloser.Close()
	return data, nil
}

type wrapCloser struct {
	io.Reader
}

func (*wrapCloser) Close() error {
	return nil
}

func WrapCloser(body io.Reader) io.ReadCloser {
	return &wrapCloser{body}
}

func ReadLines(reader io.Reader, f func(line string) bool) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if !f(scanner.Text()) {
			return nil
		}
	}
	return scanner.Err()
}
