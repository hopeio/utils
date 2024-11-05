package encoding

type Decoder interface {
	Decode(v interface{}) (err error)
}

type Encoder interface {
	Encode(v interface{}) (err error)
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

type Marshaler interface {
	Marshal(v any) ([]byte, error)
}
