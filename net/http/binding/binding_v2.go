package binding

import "io"

type Peek interface {
	Peek(key string) string
}

type Source interface {
	Uri() Peek
	Query() Peek
	Header() Peek
	BodyBind(io.Reader) error
}
