package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"wxpusher/config"
	"wxpusher/pkg/wxapi"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/skip2/go-qrcode"
)

func main() {
	// 加载环境配置
	_ = config.LoadEnv()
	router := gin.Default()

	// 服务器验证
	router.GET("/verify", func(ctx *gin.Context) {
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
	})
	// 微信事件推送回调
	router.POST("/verify", func(ctx *gin.Context) {
		body := ctx.Request.Body
		b, _ := io.ReadAll(body)
		var wxEvent wxapi.WxEvent
		err := xml.Unmarshal(b, &wxEvent)
		if err != nil {
			log.Println("Error on xml.Unmarshal: ", err)
			ctx.JSON(http.StatusOK, nil)
			return
		}
		log.Printf("POST /verify [event]: %+v", wxEvent)
		switch wxEvent.Event {
		case "SCAN":
			// 扫码事件, 已关注用户触发
			log.Printf("[Event] user %s 扫码成功.", wxEvent.FromUserName)
			wxapi.SendLoginTemplateMsg("1W-NG2s_7Ka42FnzCzEQaug5cPDaCEebNdjJYAc9sxU", wxEvent.FromUserName, config.AppConfig.WxConfig.AccessToken)
		case "subscribe":
			// 用户点击关注
			log.Printf("[Event] user %s 关注公众号.", wxEvent.FromUserName)
			wxapi.SendLoginTemplateMsg("ss8JStEE8H1TOe-uQAi-YKfJAH6FxgTlAeoynZuYa0s", wxEvent.FromUserName, config.AppConfig.WxConfig.AccessToken)
		case "unsubscribe":
			// 用户取消关注
			log.Printf("[Event] user %s 取消了关注.", wxEvent.FromUserName)
		case "TEMPLATESENDJOBFINISH":
			// 模板消息发送成功

		default:
			// 其他事件触发
			log.Printf("[Event] user %s 触发了其他事件: %+v", wxEvent.FromUserName, wxEvent)
		}
		ctx.JSON(http.StatusOK, nil)
	})

	// 获取accessToken
	router.GET("/getAccessToken", func(ctx *gin.Context) {
		acToken, err := wxapi.GetAccessToken()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"err": "请求失败",
			})
			return
		}
		config.AppConfig.WxConfig.AccessToken = acToken.Value
		ctx.JSON(http.StatusOK, acToken)
	})

	// 获取临时二维码
	router.GET("/qrCode/create", func(ctx *gin.Context) {
		ticket, err := wxapi.GetTicket()
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

	})

	log.Fatal(router.Run(":8080"))
}
