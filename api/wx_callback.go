package api

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"wxpusher/config"
	"wxpusher/pkg/wxapi"

	"github.com/gin-gonic/gin"
)

type wxCallbackController struct {
}

func NewWxCallbackRouter(router *gin.RouterGroup) {
	controller := &wxCallbackController{}
	// 微信事件推送回调
	router.POST("/verify", controller.Callback)
}

func (controller *wxCallbackController) Callback(ctx *gin.Context) {
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
		// 模板消息发送成功事件
		log.Printf("[Event] 发送模板消息给 user %s.", wxEvent.FromUserName)
	default:
		// 其他事件触发
		log.Printf("[Event] user %s 触发了其他事件: %+v", wxEvent.FromUserName, wxEvent)
	}
	ctx.JSON(http.StatusOK, nil)
}
