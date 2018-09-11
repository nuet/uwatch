package restapi

import (
	"net/url"
	"testing"
)

/**
 * 测试签名是否正确
 * @param {[type]} t *testing.T [description]
 */
func TestGenerateSign(t *testing.T) {
	app := NewWechatApp()
	postData := url.Values{
		"username": {"yahu"},
		"content":  {"240%2C"},
		"agent_id": {"34"},
	}
	t.Log(generateSign(app.Secret, postData))
}

/**
 * push 企业微信消息
 * @return {[type]} [description]
 */
func TestPushWechat(t *testing.T) {
	app := NewWechatApp()

	postData := url.Values{
		"username": {"huali,gangtie"},
		"content":  {"测试一下啊"},
		"agent_id": {"34"},
	}

	app.Post("/WechatMsg/send", postData)
}

/**
 * 短信发送测试
 */
func TestPushSMS(t *testing.T) {
	app := NewSmsApp()

	postData := url.Values{
		"number":    {"13986100639"}, //18611886113
		"position":  {"6"},
		"t_content": {"再次进行测试啊"},
		"t_time":    {"2016-08-23"},
		"t_ip":      {"192.168.1.168"},
	}

	app.Post("/SmsGate/send", postData)
}
