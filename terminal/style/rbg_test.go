package style

import (
	"fmt"
	"testing"
)

func TestGradient(t *testing.T) {
	fmt.Println(Gradient("test:这是一个长长的渐变色的字符串", colorRGB{r: 0x3B, g: 0xD1, b: 0x91}, colorRGB{r: 0x2B, g: 0x4C, b: 0xEE}))
}

func TestGradientRandom(t *testing.T) {
	fmt.Println(GradientRandom("test:这是一个长长的随机的渐变色的字符串\n还添加了换行\n还添加了unicode符号:☔☕♈♉♊\nunicode后还有字符，会影响渐变吗"))
}

func TestGradientMultiLine(t *testing.T) {
	fmt.Println(GradientMultiLine("test:这是一个长长的随机的渐变色的字符串\n还添加了换行\n还添加了unicode符号:☔☕♈♉♊\nunicode后还有字符，会影响渐变吗", colorRGB{r: 0x3B, g: 0xD1, b: 0x91}, colorRGB{r: 0x2B, g: 0x4C, b: 0xEE}))
}

func TestGradientMultiLineRandom(t *testing.T) {
	fmt.Println(GradientMultiLineRandom("test:这是一个长长的随机的渐变色的字符串\n还添加了换行\n还添加了unicode符号:☔☕♈♉♊\nunicode后还有字符，会影响渐变吗"))
}

func TestRainbow(t *testing.T) {
	fmt.Println(Rainbow("test:这是一个长长的随机的渐变色的字符串\n还添加了换行\n还添加了unicode符号:☔☕♈♉♊\nunicode后还有字符，会影响渐变吗"))
}

func TestRainbowMultiLine(t *testing.T) {
	fmt.Println(RainbowMultiLine("test:这是一个长长的随机的渐变色的字符串\n还添加了换行\n还添加了unicode符号:☔☕♈♉♊\nunicode后还有字符，会影响渐变吗"))
}
