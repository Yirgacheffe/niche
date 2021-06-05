package mockgen

type GetSetter interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}
