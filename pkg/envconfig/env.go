package envconfig

type Env interface {
	GetVar(key string, defaultValue string) string
	AutoLoadDotEnvFile()
}
