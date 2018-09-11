package falcon

import (
	"github.com/astaxie/beego"
	"controllers"
	"strings"
	"library/notice"
)


type AlertCallBackController struct {
	controllers.BaseRouter
}

//钉钉通知组 uwork通知
func (c *AlertCallBackController) ImCallback(){
	content := c.GetString("content")
	tos := c.GetString("tos")
	beego.Info(content,tos)
	users := strings.Split(tos,",")
	//uwork通知
	go func (){
		if tos != ""{
			notice.WechatSendMsg(users,content)
		}
	}()

	//钉钉通知
	go func(){
		if tos != ""{
			for _,user := range users {
				notice.SendDingMsg(user, "Open-Falcon监控报警", content, "http://uwatch.juanpi.org")
			}
		}
	}()

	c.SetJson(200,nil,"微信通知、钉钉通知 接口调用成功")
}

//邮件组通知
func (c *AlertCallBackController) MailCallback(){
	content := c.GetString("content")
	tos := c.GetString("tos")
	subject := c.GetString("subject")
	users := strings.Split(tos,",")
	mailCc := strings.Split(beego.AppConfig.String("jsbzb"),";")
	//邮件通知
	go func (){
		if tos != "" {
			notice.Mail(users, mailCc, subject, content, "Uwatch-监控告警")
		}
	}()
	c.SetJson(200,nil,"邮件通知接口调用成功")
}

//短信通知  -》语音
func (c *AlertCallBackController) VoiceCallback(){
	content := c.GetString("content")
	tos := c.GetString("tos")
	phones := strings.Split(tos,",")
	go func(){
		if tos != "" {
			for _, mobile := range phones {
				_, err := notice.SendAlertReport(content, mobile, "20011", "", "OPS", 1)
				if err != nil {
					beego.Error("语音通知失败，失败详情：", err.Error())
				}
			}
		}
	}()

	c.SetJson(200,nil,"语音通知成功")
}

func (c *AlertCallBackController) Test(){
	users := []string{"dian"}
	mailCc := []string{}
	notice.Mail(users,mailCc,"测试邮件","\nOpen-Falcon监控报警\n机器标识:10.13.44.146\n主机名称:unknown\n级别:P3\n状态:PROBLEM\n标签:/分区空间剩余20%，告警 fstype=ext4,mount=/\n告警条件:all(#3)||df.bytes.free.percent<20\n监控值:19.99618\n告警次数:1/3\n时间:2018-02-07 20:00:00\n链接:http://portal.juanpi.org/template/view/39\n","Uwatch-监控告警")


	c.SetJson(200,nil,"语音通知成功")
}
