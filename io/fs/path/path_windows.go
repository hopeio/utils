//go:build windows

package path

func lastSlash(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] != '/' && s[i] != '\\' {
		i--
	}
	return i
}
