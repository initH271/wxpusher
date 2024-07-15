package api

import "github.com/gin-gonic/gin"

func NewWxRouter(router *gin.Engine) {
	wxRouter := router.Group("wx")
	NewWxVerifyRouter(wxRouter)
	NewWxCallbackRouter(wxRouter)
	NewQrCodeRouter(wxRouter)
}
