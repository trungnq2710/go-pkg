package manager

type Config interface {
	GetConfigName() string
	Validate() (err error)
	InitClient() (interface{}, error)
}
