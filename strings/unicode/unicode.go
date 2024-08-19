package unicode

import (
	"github.com/hopeio/utils/slices"
	stringsi "github.com/hopeio/utils/strings"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// [。；，：“”（）、？《》]
var HanPunctuation = []rune{
	'\u3002', '\uff1b', '\uff0c', '\uff1a', '\u201c', '\u201d', '\uff08', '\uff09', '\u3001', '\uff1f', '\u300a', '\u300b',
}

func HasHan(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) || slices.In(r, HanPunctuation) {
			return true
		}
	}
	return false
}

// Getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func Getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	var r rune
	for _, c := range s[2:6] {
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}
		r = r*16 + rune(c)
	}
	return r
}

func ToUtf8(s []byte) string {
	if len(s) < 6 {
		return stringsi.BytesToString(s)
	}
	b := make([]byte, len(s)+2*utf8.UTFMax)
	begin, bbegin := 0, 0
	for i := 0; i+6 <= len(s); {
		if s[i] == '\\' && s[i+1] == 'u' {
			bbegin += copy(b[bbegin:], s[begin:i])
			rr := Getu4(s[i:])
			if rr < 0 {
				return stringsi.BytesToString(s)
			}
			i += 6
			if utf16.IsSurrogate(rr) {
				rr1 := Getu4(s[i:])
				if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
					// A valid pair; consume.
					i += 6
					bbegin += utf8.EncodeRune(b[bbegin:], dec)
					break
				}
				// Invalid surrogate; fall back to replacement rune.
				rr = unicode.ReplacementChar
			}
			begin = i
			bbegin += utf8.EncodeRune(b[bbegin:], rr)
		} else {
			i++
		}
	}
	bbegin += copy(b[bbegin:], s[begin:])
	return stringsi.BytesToString(b[:bbegin])
}

func ToLowerFirst(s string) string {
	if len(s) > 0 {
		return string(unicode.ToLower(rune(s[0]))) + s[1:]
	}
	return ""
}

func TrimSymbol(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

var emojiReg = regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}]`)

func TrimEmoji(s string) string {
	return emojiReg.ReplaceAllString(s, "")
}

var pattern = regexp.MustCompile(`[\p{Han}a-zA-Z0-9]+`)

func RetainChineseAndAlphanumeric(inputStr string) string {
	// 使用正则表达式匹配中文字符和字母数字
	pattern := regexp.MustCompile(`[\p{Han}a-zA-Z0-9]+`)
	filteredStr := pattern.FindAllString(inputStr, -1)
	return joinStrings(filteredStr)
}

func joinStrings(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}
