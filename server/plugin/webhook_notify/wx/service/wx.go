package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/wx/model/response"
	"go.uber.org/zap"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type WXService struct {
}

//@function: SendTextMessage
//@description: 发送企业微信文字信息
//@params content string发送的文字内容
//@return: err error

func (e *WXService) SendTextMessage(content string) (err error) {
	msg, err := renderTemplate(response.TextMessage, map[string]string{"Content": content})
	if err != nil {
		return err
	}
	return SendMessage(msg)
}

//@function: SendMarkDownMessage
//@description: 发送企业微信MarkDown信息
//@params content string发送的文字内容
//@return: err error

func (e *WXService) SendMarkDownMessage(content string) (err error) {
	msg, err := renderTemplate(response.MarkDownMessage, map[string]string{"Content": content})
	if err != nil {
		return err
	}
	return SendMessage(msg)
}

//@function: SendPostMessage
//@description: 发送企业微信卡片模板信息
//@params content string发送的文字内容
//@return: err error

func (e *WXService) SendTemplateCardMessage() (err error) {
	//模板卡片类型信息可以自助定制，此处就直接发送模板了
	return SendMessage(response.TemplateCardMessage)
}

func SendMessage(msg interface{}) error {
	client := &http.Client{}
	body := bytes.NewBuffer(nil)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, global.GlobalConfig_.Url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return fmt.Errorf("send wx message failed, %s", httpError(req, res, result, "http code is not 200"))
	}
	if err != nil {
		return fmt.Errorf("send wx message failed, %s", httpError(req, res, result, err.Error()))
	}

	type response struct {
		ErrCode int `json:"errcode"`
	}
	var ret response

	if err := json.Unmarshal(result, &ret); err != nil {
		return fmt.Errorf("send wx message failed, %s", httpError(req, res, result, err.Error()))
	}

	if ret.ErrCode != 0 {
		return fmt.Errorf("send wx message failed, %s", httpError(req, res, result, "errcode is not 0"))
	}

	return nil
}

func httpError(request *http.Request, response *http.Response, body []byte, error string) string {
	return fmt.Sprintf(
		"http request failure, error: %s, status code: %d, %s %s, body:\n%s",
		error,
		response.StatusCode,
		request.Method,
		request.URL.String(),
		string(body),
	)
}
func renderTemplate(tpl string, obj interface{}) (ret string, err error) {
	var t *template.Template
	t, err = template.New("renderTemplate").Parse(tpl)
	if err != nil {
		zap.L().Error("renderTemplate err", zap.Error(err))
		return
	}
	b := strings.Builder{}
	err = t.Execute(&b, obj)
	if err != nil {
		zap.L().Error("renderTemplate err", zap.Error(err))
		return
	}
	ret = b.String()
	return
}
