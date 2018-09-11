package notice

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bytes"
	"github.com/logrus"
	"github.com/astaxie/beego"
)

func getToken(corpid, corpsect string) string {
	resp, err := http.Get(getTokenUrl(corpid, corpsect))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var tokenObject TokenObject
	err = json.Unmarshal(body, &tokenObject)
	if err != nil {
		return ""
	}
	return tokenObject.AccessToken
}

type TokenObject struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type TextMsgObject struct {
	Touser  string      `json:"touser"`
	Toparty string      `json:"toparty"`
	Totag   string      `json:"totag"`
	Msgtype string      `json:"msgtype"`
	Agentid int         `json:"agentid"`
	Text    *TextObject `json:"text"`
	Safe    int         `json:"safe"`
}

type TextObject struct {
	Content string `json:"content"`
}

func WechatSendMsg( users []string,message string) error {
	token:=getToken(Wechat_CorpID, Wechat_Secret)
	toUsers := strings.Join(users, "|")
	text := &TextObject{
		Content: message,
	}
	textMsg := &TextMsgObject{
		Touser:  toUsers,
		Toparty: "",
		Totag:   "",
		Msgtype: "text",
		Agentid: int(Wechat_AgentID),
		Safe:    0,
		Text:    text,
	}
	jsonStr, err := json.Marshal(&textMsg)
	if err != nil {
		beego.Error(err)
		return err
	}
	req, err := http.NewRequest("POST", getSendMsgUrl(token), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	logrus.Infof("response body : %v", string(body))
	return nil
}

func getSendMsgUrl(token string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%v", token)
}

func getTokenUrl(corpid, corpsect string) string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%v&corpsecret=%v", corpid, corpsect)
}