package feishu

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/feishu/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/feishu/router"
	"github.com/gin-gonic/gin"
)

type fsPlugin struct {
	Token    string
	UseToken bool
	Url      string // https://open.feishu.cn/open-apis/bot/v2/hook/84599758-0489-4e4f-a1f7-d2ee0c218586
}

func CreateFSPlug(url, Token string, UseToken bool) *fsPlugin {
	global.GlobalConfig_.Url = url
	global.GlobalConfig_.Token = Token
	return &fsPlugin{}
}

func (*fsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRouter(group)
}

func (*fsPlugin) RouterPath() string {
	return "feishu"
}
