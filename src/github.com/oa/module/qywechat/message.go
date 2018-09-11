package qywechat

type RespHeader struct {
	ToUser  string `json:"touser,omitempty"`  // 非必须; 员工ID列表(消息接收者, 多个接收者用‘|’分隔, 最多支持1000个). 特殊情况: 指定为@all, 则向关注该企业应用的全部成员发送
	ToParty string `json:"toparty,omitempty"` // 非必须; 部门ID列表, 多个接收者用‘|’分隔, 最多支持100个. 当touser为@all时忽略本参数
	ToTag   string `json:"totag,omitempty"`   // 非必须; 标签ID列表, 多个接收者用‘|’分隔. 当touser为@all时忽略本参数

	MsgType string `json:"msgtype"` // 必须; 消息类型
	AgentId int64  `json:"agentid"` // 必须; 企业应用的id, 整型
	Safe    int    `json:"safe,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}

type RespText struct {
	RespHeader
	Text `json:"text"`
}

/**
 * 发送文本消息
 */
func (app *App) SendText(username, content string) error {
	header := RespHeader{
		username,
		"",
		"",
		"text",
		app.AgentId,
		0,
	}

	text := Text{
		content,
	}

	respText := RespText{
		header,
		text,
	}
	result := &Error{}
	err := app.Post("/message/send", respText, result)
	if err == nil {
		if result.ErrCode != ErrCodeOK {
			return result
		}
	}
	return err
}
