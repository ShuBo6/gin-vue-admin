package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/feishu/api"
	"github.com/gin-gonic/gin"
)

type FeishuRouter struct {
}

func (s *FeishuRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router
	var SendTextMessage = api.ApiGroupApp.Api.SendTextMessage
	var SendPostMessage = api.ApiGroupApp.Api.SendPostMessage
	{
		router.POST("sendTextMessage", SendTextMessage)
		router.POST("sendPostMessage", SendPostMessage)
	}
}
