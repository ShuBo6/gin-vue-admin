package response

// 具体配置看文档 此处为只是供参考的文本和模板消息，在Api中使用的是template的模板

//企业微信文档地址:https://developer.work.weixin.qq.com/document/path/91770

func NewTextAndImageMessage(msgType string, content interface{}) TextAndImageMessage {
	return TextAndImageMessage{
		Msgtype: msgType,
		Content: content,
	}
}

// TextContent 文本类型
type TextContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

// MarkDownContent markdown类型
type MarkDownContent struct {
	Content string `json:"content"`
}

// TextAndImageMessage 发送 文本，图片，markdown类型使用的结构体
type TextAndImageMessage struct {
	Msgtype string      `json:"msgtype"`
	Content interface{} `json:"content"` //基础类型
}

// ImageContent 图片类型
type ImageContent struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

// NewsContent 图文类型
type NewsContent struct {
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Picurl      string `json:"picurl"`
	} `json:"articles"`
}

// TemplateCardContent 发送模板卡片类型消息使用的结构体
type TemplateCardContent struct {
	Msgtype      string `json:"msgtype"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl   string `json:"icon_url"`
			Desc      string `json:"desc"`
			DescColor int    `json:"desc_color"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		EmphasisContent struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"emphasis_content"`
		QuoteArea struct {
			Type      int    `json:"type"`
			Url       string `json:"url"`
			Appid     string `json:"appid"`
			Pagepath  string `json:"pagepath"`
			Title     string `json:"title"`
			QuoteText string `json:"quote_text"`
		} `json:"quote_area"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int    `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int    `json:"type"`
			Url      string `json:"url,omitempty"`
			Title    string `json:"title"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int    `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
}
