package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/api"
	"github.com/gin-gonic/gin"
)

type WXRouter struct {
}

func (s *WXRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router
	var SendTextMessage = api.ApiGroupApp.Api.SendTextMessage
	var SendMarkDownMessage = api.ApiGroupApp.Api.SendMarkDownMessage
	var SendTemplateCardMessage = api.ApiGroupApp.Api.SendTemplateCardMessage
	{
		router.POST("sendTextMessage", SendTextMessage)
		router.POST("sendMarkDownMessage", SendMarkDownMessage)
		router.POST("sendTemplateCardMessage", SendTemplateCardMessage)
	}
}
