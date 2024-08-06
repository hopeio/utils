package encoding

type Format string

const (
	Json     Format = "json"
	Yaml     Format = "yaml"
	Toml     Format = "toml"
	Yml      Format = "yml"
	Protobuf Format = "protobuf"
	Xml      Format = "xml"
	Text     Format = "text"
	Base64   Format = "base64"
)
