package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	feishu_response "github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/feishu/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/feishu/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Api struct {
}

// @Tags Feishu
// @Summary 发送文字消息接口
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body feishu_response.TextFeishu true "发送文字消息的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /feishu/sendTextMessage [post]
func (s *Api) SendTextMessage(c *gin.Context) {
	var textFeishu feishu_response.TextFeishu
	_ = c.ShouldBindJSON(&textFeishu)
	if err := service.ServiceGroupApp.SendTextMessage(textFeishu.Content); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

// @Tags Feishu
// @Summary 发送富文本消息接口
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body feishu_response.PostFeishu true "发送富文本消息的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /feishu/sendPostMessage [post]
func (s *Api) SendPostMessage(c *gin.Context) {
	var postFeishu feishu_response.PostFeishu
	_ = c.ShouldBindJSON(&postFeishu)
	if err := service.ServiceGroupApp.SendPostMessage(postFeishu.Content); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
