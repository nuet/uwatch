package notice


import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/heroicyang/wechat-qy/api"
)

var Wechat_CorpID = beego.AppConfig.String("Wechat_CorpID")
var Wechat_Secret = beego.AppConfig.String("Wechat_Secret")
var Wechat_Token = beego.AppConfig.String("Wechat_Token")
var Wechat_EncodingAESKey = beego.AppConfig.String("Wechat_EncodingAESKey")
var Wechat_AgentID, _ = beego.AppConfig.Int64("Wechat_AgentID")

func Wechat_send(msg map[string]string) {
	title := msg["title"]
	content := msg["content"]
	username := msg["username"]
	depid := msg["depid"]
	wechatAPI := api.New(Wechat_CorpID, Wechat_Secret, Wechat_Token, Wechat_EncodingAESKey)
	//主动发送消息
	article := api.Article{
		Title:       title,
		Description: content,
		URL:         "",
		PicURL:      "",
	}
	articles := api.Articles{
		Articles: []api.Article{},
	}
	articles.Articles = append(articles.Articles, article)
	sentext := api.NewsMessage{
		ToUser:  username,
		ToParty: depid,
		ToTag:   "",
		MsgType: api.NewsMsg,
		AgentID: Wechat_AgentID,
		News:    articles,
	}
	//发送消息
	err := wechatAPI.SendMessage(sentext)
	if err != nil {
		beego.Error("微信企业号发送消息错误:", err.Error())
	} else {
		data, _ := json.Marshal(sentext)
		beego.Info("微信企业号发送消息成功:", string(data))
	}
}

