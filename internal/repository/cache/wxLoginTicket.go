package cache

import (
	"context"
	"fmt"
	"time"
	"wxpusher/internal/domain"

	"github.com/redis/go-redis/v9"
)

type WxLoginTicketCache struct {
	cmd       redis.Cmdable
	keyPrefix string
}

func NewWxLoginTicketCache(cmd redis.Cmdable) *WxLoginTicketCache {
	return &WxLoginTicketCache{
		cmd:       cmd,
		keyPrefix: "wx_qrcode_login_status",
	}
}

func (c *WxLoginTicketCache) Set(ctx context.Context, ticket string, status domain.LoginStatus, expireTime time.Duration) error {
	return c.cmd.Set(ctx, c.key(ticket), int(status), expireTime).Err()
}

func (c *WxLoginTicketCache) Get(ctx context.Context, ticket string) (status domain.LoginStatus, err error) {
	key := c.key(ticket)
	v, err := c.cmd.Get(ctx, key).Int()
	if err != nil {
		return
	}
	status = domain.LoginStatus(v)
	return
}

func (c *WxLoginTicketCache) Delete(ctx context.Context, ticket string) (err error) {
	key := c.key(ticket)
	err = c.cmd.Del(ctx, key).Err()
	return
}

func (c *WxLoginTicketCache) key(ticket string) string {
	return fmt.Sprintf("%s:%s", c.keyPrefix, ticket)
}
