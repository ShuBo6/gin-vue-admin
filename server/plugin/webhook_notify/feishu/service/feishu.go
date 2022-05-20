package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/webhook_notify/feishu/global"
	"io/ioutil"
	"net/http"
	"time"
)

type FeishuService struct {
}

//@function: SendTextMessage
//@description: 发送飞书文字信息
//@params content string发送的文字内容
//@return: err error

func (e *FeishuService) SendTextMessage(content string) (err error) {
	timestamp := time.Now().UnixNano()
	sign, _ := GenSign(global.GlobalConfig_.Token, timestamp)
	msg := map[string]interface{}{
		"timestamp": timestamp,
		"sign":      sign,
		"msg_type":  "text",
		"content": map[string]string{
			"text": content,
		},
	}
	return SendMessage(msg)
}

//@function: SendPostMessage
//@description: 发送飞书文字信息
//@params content string发送的文字内容
//@return: err error

func (e *FeishuService) SendPostMessage(content interface{}) (err error) {
	timestamp := time.Now().UnixNano()
	sign, _ := GenSign(global.GlobalConfig_.Token, timestamp)
	msg := map[string]interface{}{
		"timestamp": timestamp,
		"sign":      sign,
		"msg_type":  "post",
		"content":   content,
	}
	return SendMessage(msg)
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
		return fmt.Errorf("send dingTalk message failed, %s", httpError(req, res, result, "http code is not 200"))
	}
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(req, res, result, err.Error()))
	}

	type response struct {
		ErrCode int `json:"errcode"`
	}
	var ret response

	if err := json.Unmarshal(result, &ret); err != nil {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(req, res, result, err.Error()))
	}

	if ret.ErrCode != 0 {
		return fmt.Errorf("send dingTalk message failed, %s", httpError(req, res, result, "errcode is not 0"))
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

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
