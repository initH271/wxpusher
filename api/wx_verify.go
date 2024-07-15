package api

import (
	"log"
	"net/http"
	"wxpusher/config"
	"wxpusher/pkg/wxapi"

	"github.com/gin-gonic/gin"
)

type wxVerifyController struct {
}

func NewWxVerifyRouter(router *gin.RouterGroup) {
	controller := &wxVerifyController{}
	// 服务器验证
	router.GET("/verify", controller.Verify)
	// 获取accessToken
	router.GET("/getAccessToken", controller.GetAccessToken)
}

func (controller *wxVerifyController) Verify(ctx *gin.Context) {
	signature := ctx.Query("signature")
	timestamp := ctx.Query("timestamp")
	nonce := ctx.Query("nonce")
	echostr := ctx.Query("echostr")
	valid := wxapi.CheckSignature(config.AppConfig.WxConfig.Token, timestamp, nonce, signature)
	if valid {
		log.Println("[/]:  the signature is valid")
		ctx.String(http.StatusOK, echostr)
		return
	}
	ctx.String(http.StatusOK, "Bad signature")
	log.Println("[/]: Bad signature")
}

func (controller *wxVerifyController) GetAccessToken(ctx *gin.Context) {
	acToken, err := wxapi.GetAccessToken(config.AppConfig.WxConfig.AppID, config.AppConfig.WxConfig.AppSecret)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "请求失败",
		})
		return
	}
	config.AppConfig.WxConfig.AccessToken = acToken.Value
	ctx.JSON(http.StatusOK, acToken)
}
