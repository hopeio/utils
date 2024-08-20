package msgpack

import "github.com/vmihailenco/msgpack/v5"

func Marshal(v any) ([]byte, error) {
	return msgpack.Marshal(v)
}
