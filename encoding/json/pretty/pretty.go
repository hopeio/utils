package pretty

import "github.com/tidwall/pretty"

func Pretty(json []byte) []byte { return pretty.PrettyOptions(json, nil) }
func PrettyOptions(json []byte, options *pretty.Options) []byte {
	return pretty.PrettyOptions(json, options)
}
