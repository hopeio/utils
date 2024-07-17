//go:build unix

package path

// lastSlash(s) is strings.LastIndex(s, "/") but we can't import strings.
func lastSlash(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] != '/' {
		i--
	}
	return i
}
