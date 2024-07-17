package markdown

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// MarkdownToHTML 将markdown 转换为 html
func MarkdownToHTML(input []byte) []byte {
	unsafe := blackfriday.Run(input)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}
