package usecase

import (
	"context"
	"log"
	"time"
	"wxpusher/internal/domain"
	"wxpusher/internal/repository/cache"
	"wxpusher/pkg/wxapi"

	"github.com/redis/go-redis/v9"
)

type WxAuthUsecase struct {
	loginTicketCache *cache.WxLoginTicketCache
	onlineUserCache  *cache.OnlineUserCache
}

func NewWxAuthUsecase(cmd redis.Cmdable) *WxAuthUsecase {
	return &WxAuthUsecase{
		loginTicketCache: cache.NewWxLoginTicketCache(cmd),
		onlineUserCache:  cache.NewOnlineUserCache(cmd),
	}
}

func (u *WxAuthUsecase) CacheTicket(ctx context.Context, ticket wxapi.Ticket, expireTime time.Duration) (err error) {
	err = u.loginTicketCache.Set(ctx, ticket.Value, domain.WaitScan, expireTime)
	if err != nil {
		log.Println("[usecase->auth] [ERROR] CacheTicket: ", err)
	}
	return
}

func (u *WxAuthUsecase) GetLoginStatusWithTicket(ctx context.Context, ticket string) (status domain.LoginStatus) {
	status, err := u.loginTicketCache.Get(ctx, ticket)
	if err != nil {
		log.Println("[usecase->auth] [ERROR] CheckLoginWithTicket: ", err)
		status = domain.IsInvalid
		return
	}
	return
}

func (u *WxAuthUsecase) DoLoginStatusWithTicket(ctx context.Context, ticket string, openId string) (err error) {
	// 修改为已登录
	err = u.loginTicketCache.Set(ctx, ticket, domain.IsLogin, 2*time.Minute)
	if err != nil {
		log.Println("[usecase->auth] [ERROR] DoLoginStatusWithTicket: ", err)
		return
	}
	// 更新在线用户列表
	err = u.onlineUserCache.Set(ctx, openId)
	return
}
func (u *WxAuthUsecase) UnSubscribeWithOpenId(ctx context.Context, openId string) (err error) {
	// 修改为已登录
	err = u.onlineUserCache.Delete(ctx, openId)
	if err != nil {
		log.Println("[usecase->auth] [ERROR] UnSubscribeWithOpenId: ", err)
		return
	}
	return
}
