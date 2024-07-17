package types

func PairOf[T1 any, T2 any](f T1, s T2) *Pair[T1, T2] {
	return &Pair[T1, T2]{f, s}
}

func PairEmpty[T1 any, T2 any]() *Pair[T1, T2] {
	return &Pair[T1, T2]{}
}

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func (a *Pair[T1, T2]) Val() (T1, T2) {
	return a.First, a.Second
}

func TupleOf[T1 any, T2 any, T3 any](f T1, s T2, t T3) *Tuple[T1, T2, T3] {
	return &Tuple[T1, T2, T3]{f, s, t}
}

type Tuple[T1 any, T2 any, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

func (a *Tuple[T1, T2, T3]) Val() (T1, T2, T3) {
	return a.First, a.Second, a.Third
}
