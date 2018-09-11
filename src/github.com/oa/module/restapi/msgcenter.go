package restapi

type MsgCenterParam struct {
	ChannelID string   `json:"channelId"`
	UserInfo  UserInfo `json:"userInfo"`
	Ext       Ext      `json:"ext"`
	Params    Params   `json:"params"`
}

type UserInfo struct {
	JPID   string `json:"jpid"`
	Mobile string `json:"mobile"`
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type Ext map[string]string

type Params struct {
	Push   Push   `json:"push"`
	Mobile Mobile `json:"mobile"`
	Weixin Weixin `json:"weixin"`
}

type Mobile struct {
	LockTime string `json:"lockTime"`
	LockDay  string `json:"lockday"`
}

type Push struct {
	Appname  string `json:"appname"`
	ComeFrom string `json:"comefrom"`
	MsgType  string `json:"msgtype"`
	TaskID   string `json:"taskId"`
}

type Weixin struct {
	QueueType   string `json:"queue_type"`
	Suid        string `json:"suid"`
	Permissions string `json:"permissions"`
}

/**
 * 消息中心
 */
func NewMsgCenterApp() *App {
	return &App{
		"soa_service",
		"#5%78soa{#*&iad&^",
	}
}
