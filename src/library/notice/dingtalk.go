package notice

import (
	"crypto/tls"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

type QxRet struct {
	Data    interface{} `json:"data"`
	ErrCode string      `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
}

/*
 * username: 发送人, haosheng|dian|siji
 * title: 标题
 * content: 内容
 */
func SendDingMsg(username, title, content, url string) {
	beego.Info(username, title)
	msgtype := "text"
	if url != "" {
		msgtype = "link"
	}

	req1 := httplib.Get("https://oa.juanpi.org:11433/v1/token")
	req1.Param("appid", "2017072101")
	req1.Param("appsecret", "a4133b7700b8a49fa6c112f582b6797b")
	req1.SetTimeout(10*time.Second, 10*time.Second)
	req1.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req1.Header("content-type", "application/json; charset=UTF-8")
	ret, err := req1.String()
	if err != nil {
		beego.Error(err.Error())
		return
	}
	var token_arr Token
	err = json.Unmarshal([]byte(ret), &token_arr)
	if err != nil {
		beego.Error("json格式不正确", err.Error())
	}

	req2 := httplib.Get("https://oa.juanpi.org:11433/v1/wechatqy/send")
	req2.Param("username", username)
	req2.Param("msgtype", msgtype)
	req2.Param("title", title)
	req2.Param("url", url)
	req2.Param("content", content)
	req2.SetTimeout(10*time.Second, 10*time.Second)
	req2.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req2.Header("content-type", "application/json; charset=UTF-8")
	req2.Header("Authorization", token_arr.AccessToken)
	ret2, err := req2.String()
	if err != nil {
		beego.Error(err.Error())
		return
	}
	var wx_ret QxRet
	err = json.Unmarshal([]byte(ret2), &wx_ret)
	beego.Info(wx_ret.ErrMsg)
	if wx_ret.ErrMsg != "OK" {
		beego.Error("钉钉消息通知发送失败")
		//处理逻辑
	}
}
