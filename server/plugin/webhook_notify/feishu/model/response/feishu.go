package response

// 具体配置看文档 此处为只封装普通和富文本，其余更加简单 参考 https://open.feishu.cn/document/ukTMukTMukTM/uMDMxEjLzATMx4yMwETM
// https://www.feishu.cn/hc/zh-CN/articles/360024984973#lineguid-CdVcCt

type TextFeishu struct { // 信息结构体
	Content string `json:"content"` // 发送的内容
}

type PostFeishu struct { // 信息结构体
	Content struct {
		Post struct {
			ZhCn struct {
				Title   string        `json:"title"`
				Content []interface{} `json:"content"`
			} `json:"zh_cn"`
		} `json:"post"`
	} `json:"content"` // 发送的内容

}
