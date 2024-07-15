package domain

import (
	"context"
	"time"
	"wxpusher/pkg/wxapi"
)

type LoginStatus int

const (
	WaitScan LoginStatus = iota
	IsLogin
	IsInvalid
)

type IWxAuthUsecase interface {
	// 缓存凭证
	CacheTicket(ctx context.Context, ticket wxapi.Ticket, expireTime time.Duration) error
	// 通过ticket获取登录状态
	GetLoginStatusWithTicket(ctx context.Context, ticket string) LoginStatus
	// 通过ticket已登录
	DoLoginStatusWithTicket(ctx context.Context, ticket string, openId string) error
	// 取消关注
	UnSubscribeWithOpenId(ctx context.Context, openId string) error
}
