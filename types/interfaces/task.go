package interfaces

type FuncContinue interface {
	Do(times uint) (retry bool)
}
