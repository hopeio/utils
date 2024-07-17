package strings

// Is c an ASCII lower-case letter?
func IsASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func IsASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsASCIILetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

// Is c an ASCII digit?
func IsASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// 有一个匹配成功就返回true
func HasPrefixes(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if len(s) >= len(prefix) && s[0:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

func IsASCIILowers(s string) bool {
	for _, c := range s {
		if 'a' < c || c > 'z' {
			return false
		}
	}
	return true
}

func IsASCIIUppers(s string) bool {
	for _, c := range s {
		if 'A' < c || c > 'Z' {
			return false
		}
	}
	return true
}

func IsASCIILetters(s string) bool {
	for _, c := range s {
		if c < 'A' || c > 'z' || (c > 'Z' && c < 'a') {
			return false
		}
	}
	return true
}
