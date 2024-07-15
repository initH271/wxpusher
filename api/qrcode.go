package api

import (
	"log"
	"net/http"
	"time"
	"wxpusher/config"
	"wxpusher/internal/domain"
	"wxpusher/internal/usecase"
	"wxpusher/pkg/wxapi"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/skip2/go-qrcode"
)

type qrCodeController struct {
	authUsecase domain.IWxAuthUsecase
}

func NewQrCodeRouter(router *gin.RouterGroup, rdb redis.Cmdable) {
	controller := &qrCodeController{
		authUsecase: usecase.NewWxAuthUsecase(rdb),
	}
	qrCodeRouter := router.Group("qrCode")
	qrCodeRouter.GET("/createImage", controller.CreateWxQrCodeImage)
	qrCodeRouter.GET("/create", controller.CreateWxQrCode)
}

func (controller *qrCodeController) CreateWxQrCode(ctx *gin.Context) {
	ticket, err := wxapi.GetTicket(config.AppConfig.WxConfig.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取凭证失败",
		})
		return
	}
	// 缓存凭证到redis
	controller.authUsecase.CacheTicket(ctx, ticket, 2*time.Minute)
	ctx.JSON(http.StatusOK, ticket)

}

func (controller *qrCodeController) CreateWxQrCodeImage(ctx *gin.Context) {
	ticket, err := wxapi.GetTicket(config.AppConfig.WxConfig.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取二维码图片失败",
		})
		return
	}
	// 生成二维码
	qr, err := qrcode.New(ticket.Url, qrcode.Medium)
	if err != nil {
		log.Println("qrcode.New: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成二维码失败"})
		return
	}
	controller.authUsecase.CacheTicket(ctx, ticket, 2*time.Minute)

	// 将二维码图片写入响应
	ctx.Header("Content-Type", "image/png")
	b, _ := qr.PNG(256)
	ctx.Writer.Write(b)

}
