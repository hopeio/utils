package os

import (
	"github.com/hopeio/gox/log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Split(line string) []string {
	var words []string
Words:
	for {
		line = strings.TrimLeft(line, " \t")
		if len(line) == 0 {
			break
		}
		if line[0] == '"' {
			for i := 1; i < len(line); i++ {
				c := line[i] // Only looking for ASCII so this is OK.
				switch c {
				case '\\':
					if i+1 == len(line) {
						log.Panic("bad backslash")
					}
					i++ // Absorb next byte (If it's a multibyte we'll get an error in Unquote).
				case '"':
					word, err := strconv.Unquote(line[0 : i+1])
					if err != nil {
						log.Panic("bad quoted string")
					}
					words = append(words, word)
					line = line[i+1:]
					// Check the next character is space or end of line.
					if len(line) > 0 && line[0] != ' ' && line[0] != '\t' {
						log.Panic("expect space after quoted argument")
					}
					continue Words
				}
			}
			log.Panic("mismatched quoted string")
		}
		i := strings.IndexAny(line, " \t")
		if i < 0 {
			i = len(line)
		}
		words = append(words, line[0:i])
		line = line[i:]
	}
	// Substitute command if required.

	// Substitute environment variables.
	for i, word := range words {
		words[i] = os.Expand(word, expandVar)
	}
	return words
}

var env = []string{
	"GOARCH=" + runtime.GOARCH,
	"GOOS=" + runtime.GOOS,
}

func expandVar(word string) string {
	w := word + "="
	for _, e := range env {
		if strings.HasPrefix(e, w) {
			return e[len(w):]
		}
	}
	return os.Getenv(word)
}
