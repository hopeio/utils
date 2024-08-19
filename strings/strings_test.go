package strings

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type test struct {
	input, expected string
}

var camelToSnakeTests = []test{
	{"", ""},
	{"camelCase", "camel_case"},
	{"snakeCase", "snake_case"},
	{"PascalCase", "pascal_case"},
	{"kebab-case", "kebab-case"}, // No change for other cases

}

func TestCamelToSnake(t *testing.T) {
	for _, test := range camelToSnakeTests {
		result := CamelToSnake(test.input)
		if result != test.expected {
			t.Errorf("CamelToSnake('%s') = '%s', expected '%s'", test.input, result, test.expected)
		}
	}
}

func TestFormatLen(t *testing.T) {
	s := "post"
	log.Println(FormatLen(s, 10), "test")
	s = "AutoCommit"
	log.Println(CamelToSnake(s))
}

func TestReplaceBytes(t *testing.T) {
	s := "p我o爱s中t"
	log.Println(ReplaceBytes(s, []byte{'o'}, '-'))
	log.Println(ReplaceBytes(s, []byte{'o', 's'}, '-'))
	log.Println(ReplaceBytes(s, []byte{'o', 't'}, '-'))
	log.Println(ReplaceBytes(s, []byte{'p', 't'}, '-'))
}

func TestReplaceRunesEmpty(t *testing.T) {
	s := "p我o爱s中t"
	log.Println(RemoveRunes(s, 'o'))
	log.Println(RemoveRunes(s, 'o', 's'))
	log.Println(RemoveRunes(s, 'o', 't'))
	log.Println(RemoveRunes(s, '中', 't'))
}

// TODO
func TestCountdownCutoff(t *testing.T) {
	log.Println(ReverseCutPart("", "%2F"))
	log.Println(ReverseCutPart("", "/"))
	log.Println(CutPart("", "wx1"))
	log.Println(ReverseCutPart(CutPartContain("", "mp4"), "/"))
	log.Println(CutPart("6108162447_4848748796058856_20221220134741_006Fne59ly1h9a87sb8d7j52802yo4qr2.jpg", "?"))
	baseUrl := ReverseCutPart("", "/")
	log.Println(baseUrl)
	baseUrl = CutPart(baseUrl, "?")
	log.Println(baseUrl)
}

var upperCaseFirstTests = []test{
	{"local", "Local"},
	{"dev", "Dev"},
	{"prod", "Prod"},
	{"1prod", "1prod"},
}

func TestUpperCaseFirst(t *testing.T) {
	for _, tt := range upperCaseFirstTests {
		out := UpperCaseFirst(tt.input)
		if tt.expected != out {
			t.Errorf("UpperCaseFirst(%q) = %q, want %q", tt.input, out, tt.expected)
		}
	}
}

func FuzzUpperCaseFirst(f *testing.F) {
	for _, tt := range upperCaseFirstTests {
		f.Add(tt.input)
	}
	f.Fuzz(func(t *testing.T, str string) {
		UpperCaseFirst(str)
	})
}

func TestBracketsIntervals(t *testing.T) {
	tests := []struct {
		s          string
		tokenBegin rune
		tokenEnd   rune
		expected   string
		expected2  bool
	}{
		{"(test)", '(', ')', "(test)", true},
		{"(test)", '[', ']', "", false},
		{"[(test)]", '(', ')', "(test)", true},
		{"[(test)]", '[', ']', "[(test)]", true},
		{"((test))", '(', ')', "((test))", true},
		{"((test))", '[', ']', "", false},
		{"", '(', ')', "", false},
	}

	for _, test := range tests {
		result, result2 := BracketsIntervals(test.s, test.tokenBegin, test.tokenEnd)
		if result != test.expected {
			t.Errorf("BracketsIntervals(%s, %v, %v) = %s; want %s", test.s, test.tokenBegin, test.tokenEnd, result, test.expected)
		}
		if result2 != test.expected2 {
			t.Errorf("BracketsIntervals(%s, %v, %v) = %t; want %t", test.s, test.tokenBegin, test.tokenEnd, result2, test.expected2)
		}
	}
}

func TestConvert(t *testing.T) {
	for i := 'a'; i <= 'z'; i++ {
		t.Log(string(i^' '), string(i^' '^' '))
	}
}

func TestTrimSymbol(t *testing.T) {
	assert.Equal(t, "Hello世界123", RemoveSymbol("Hello, 世界! 123"))
	assert.Equal(t, "Hello世界123", RemoveSymbol("Hello, 世界! 😊 123"))
	assert.Equal(t, "Hello, 世界!  123", RemoveEmoji("Hello, 世界! 😊 123"))
	assert.Equal(t, "Hello世界123", RemoveSymbol("Hello_世界_123"))
	assert.Equal(t, "是谁的小篮球", RemoveSymbol("是谁的小篮球🏀？"))
	assert.Equal(t, "汉字567_.mp4\\\"", RetainHanAndASCIIGt32("૮𖥦აʚɞ汉字567_.mp4\n\t\\\a\b\f\r\v\""))
	assert.Equal(t, "", RetainHanAndASCIIGt32("ᴿᵃʳᵉⁱⁿᵗʰᵉʷᵒʳˡᵈʸᵒᵘʳᵉʷᵒʳᵗʰⁱᵗ"))
	assert.Equal(t, "#@!~$%&^*(){}[]", RetainHanAndASCIIGt32("#@!！~·￥$%&^*()【】{}[]"))
}
