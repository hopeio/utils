package console

import (
	"fmt"
	"strings"
)

func DrawProgressBar(prefix string, proportion float32, width int, suffix ...string) {
	pos := int(proportion * float32(width))
	fmt.Printf("[%s] %s%*s %6.2f%% %s\r",
		prefix, strings.Repeat("■", pos), width-pos, "", proportion*100, strings.Join(suffix, ""))
}
