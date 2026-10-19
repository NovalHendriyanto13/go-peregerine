package RedisType

import (
	"context"
	"errors"
)

// ErrNotFound is returned when the requested key does not exist
var ErrNotFound = errors.New("redis: key not found")

type SettingService interface {
	Set(ctx context.Context, group, key, value string) error
	Get(ctx context.Context, key, value string) (string, error)
	GetAll(ctx context.Context, group string) (map[string] string, error)
}
