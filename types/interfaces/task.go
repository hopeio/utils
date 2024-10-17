package interfaces

type TaskRetry interface {
	Do(times uint) (retry bool)
}
