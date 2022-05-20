package feishu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/router"
	"github.com/gin-gonic/gin"
)

// WXPlugin 企业微信文档地址:https://developer.work.weixin.qq.com/document/path/91770
type WXPlugin struct {
	key string
	Url string
}

func CreateWXPlug(url, key string) *WXPlugin {
	global.GlobalConfig_.Url = url
	global.GlobalConfig_.Key = key
	return &WXPlugin{}
}

func (*WXPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRouter(group)
}

func (*WXPlugin) RouterPath() string {
	return "wx"
}
