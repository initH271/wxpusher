package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	cmd    redis.Cmdable
	keyPrefix string
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:    cmd,
		keyPrefix: "user",
	}
}

func (c *UserCache) key(value string) string {
	return fmt.Sprintf("%s:%s", c.keyPrefix, value)
}
