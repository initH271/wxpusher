package api

import (
	"log"
	"net/http"
	"wxpusher/config"
	"wxpusher/pkg/wxapi"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type qrCodeController struct {
}

func NewQrCodeRouter(router *gin.RouterGroup) {
	controller := &qrCodeController{}
	qrCodeRouter := router.Group("qrCode")
	qrCodeRouter.GET("/create", controller.CreateWxQrCode)
}

func (controller *qrCodeController) CreateWxQrCode(ctx *gin.Context) {
	ticket, err := wxapi.GetTicket(config.AppConfig.WxConfig.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "获取二维码",
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
	// 将二维码图片写入响应
	ctx.Header("Content-Type", "image/png")
	b, _ := qr.PNG(256)
	ctx.Writer.Write(b)

}
