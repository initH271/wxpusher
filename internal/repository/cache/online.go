package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type OnlineUserCache struct {
	cmd        redis.Cmdable
	keyPrefix  string
	expireTime time.Duration
}

func NewOnlineUserCache(cmd redis.Cmdable) *OnlineUserCache {
	return &OnlineUserCache{
		cmd:        cmd,
		keyPrefix:  "online_user",
		expireTime: time.Hour * 72,
	}
}

func (c *OnlineUserCache) key(value string) string {
	return fmt.Sprintf("%s:%s", c.keyPrefix, value)
}

func (c *OnlineUserCache) Set(ctx context.Context, token string) (err error) {
	err = c.cmd.Set(ctx, c.key(token), 1, c.expireTime).Err()
	return
}

func (c *OnlineUserCache) Get(ctx context.Context, token string) (res int, err error) {
	res, err = c.cmd.Get(ctx, c.key(token)).Int()
	return
}

func (c *OnlineUserCache) Delete(ctx context.Context, token string) (err error) {
	err = c.cmd.Del(ctx, c.key(token)).Err()
	return
}
