package restapi

import "testing"

/**
 * 测试消息中心的消息推送,测试商家微信模版消息推送
 */
func TestMsgCenter(t *testing.T) {
	app := NewWechatApp()

	weixin := Weixin{
		QueueType:   "4",
		Suid:        "",
		Permissions: "a,b,c,d",
	}

	msgCenterParam := MsgCenterParam{
		// ChannelID: "229",
		ChannelID: "136",
		UserInfo: UserInfo{
			Uid: "3201056",
		},
		Params: Params{
			Weixin: weixin,
		},
		Ext: Ext{
			// "first":    "测试一下啊",
			// "keyword1": "测试2",
			// "keyword2": "测试3",
			// "remark":   "测试4",
			"first":    "违规商家处罚提醒",
			"keyword1": "亲爱的reyew2",
			"keyword2": "2016年09月17日",
			"keyword3": "扣除20.00元",
			"keyword4": "保证金余额49215.90元",
			"keyword5": "您有1笔订单被判定为延时发货处罚，系统已自动扣除20.00元作为处罚。",
			"remark":   "测试",
		},
	}

	app.PostJSON("/MsgCenter/send", msgCenterParam)
}
