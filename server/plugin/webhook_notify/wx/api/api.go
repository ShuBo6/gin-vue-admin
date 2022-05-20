package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Api struct {
}

func (s *Api) SendTextMessage(c *gin.Context) {
	var req request.TextContentReq
	_ = c.ShouldBindJSON(&req)
	if err := service.ServiceGroupApp.SendTextMessage(req.Content); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

func (s *Api) SendMarkDownMessage(c *gin.Context) {
	var req request.TextContentReq
	_ = c.ShouldBindJSON(&req)
	if err := service.ServiceGroupApp.SendMarkDownMessage(req.Content); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

func (s *Api) SendTemplateCardMessage(c *gin.Context) {
	if err := service.ServiceGroupApp.SendTemplateCardMessage(); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
