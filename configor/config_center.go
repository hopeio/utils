package configor

type ConfigCenter interface {
	NewClient() Client
}

type Client interface {
	GetConfig() ([]byte, error)
	SetConfig(func([]byte)) error
	Listener(func([]byte)) error
}
