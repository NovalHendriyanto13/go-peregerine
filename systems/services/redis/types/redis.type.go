package RedisType

type SettingService interface {
	Set(group, key, value string) error
	Get(key, value string) (string, error)
	GetAll(group string) (map[string] string, error)
}

// type RedisHandler 