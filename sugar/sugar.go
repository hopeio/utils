package sugar

func TernaryOperator[T any](v bool, a, b T) T {
	if v {
		return a
	}
	return b
}
