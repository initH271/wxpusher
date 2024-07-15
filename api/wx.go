package api

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewWxRouter(router *gin.Engine, rdb redis.Cmdable) {
	wxRouter := router.Group("wx")
	NewWxVerifyRouter(wxRouter, rdb)
	NewWxCallbackRouter(wxRouter, rdb)
	NewQrCodeRouter(wxRouter, rdb)
}
