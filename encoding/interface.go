package encoding

type Decoder interface {
	Decode(v interface{}) (err error)
}

type Encoder interface {
	Encode(v interface{}) (err error)
}
