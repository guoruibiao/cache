package cache


type Cache interface {
	Cache(key string, data interface{}) (bool, error)
	Get(key string)(interface{}, error)
	Delete(key string) error
}