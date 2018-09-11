package notice

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/utils"
)

// 邮件通知配置
type emailConfig struct {
	UserName string `json:"username,omitempty"`
	PassWord string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
}

//subject邮件主题 content邮件内容  tittle模板标题
func Mail(toMailName,toMailCc []string, subject,content,tittle string) error{
	config := emailConfig{}
	config.UserName = beego.AppConfig.String("emailUsername")
	config.PassWord = beego.AppConfig.String("emailPwd")
	config.Host = beego.AppConfig.String("emailHost")
	config.Port, _ = beego.AppConfig.Int("emailPort")

	b, _ := json.Marshal(config)
	mail := utils.NewEMail(string(b))
	mail.To = toMailName
	mail.Cc = toMailCc
	mail.From = tittle+"<itil@juanpi.com>"
	mail.Subject = subject
	mail.Text = "中文"

	mail.HTML = `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
    <html xmlns="http://www.w3.org/1999/xhtml">
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta name="viewport" content="width=device-width"/>
        <style type="text/css">
            * {
              margin: 0;
              padding: 0;
              font-size: 100%;
              font-family: 'Avenir Next', "Helvetica Neue", "Helvetica", Helvetica, Arial, sans-serif;
              line-height: 1.65; }

            img {
              max-width: 100%;
              margin: 0 auto;
              display: block; }

            body,
            .body-wrap {
              width: 100% !important;
              height: 100%;
              background: #efefef;
              -webkit-font-smoothing: antialiased;
              -webkit-text-size-adjust: none; }

            a {
              color: #379dbc;
              text-decoration: none; }

            .text-center {
              text-align: center; }

            .text-right {
              text-align: right; }

            .text-left {
              text-align: left; }

            .button {
              display: inline-block;
              color: white;
              background: #379dbc;
              border: solid #379dbc;
              border-width: 10px 20px 8px;
              font-weight: bold;
              border-radius: 4px; }

            h1, h2, h3, h4, h5, h6 {
              margin-bottom: 10px;
              line-height: 1.25; }


            h1 {
              font-size: 24px; }

            h4 {
              font-size: 20px; }

            p, ul, ol {
              font-size: 16px;
              font-weight: normal;
              margin-bottom: 20px; }

            .container {
              display: block !important;
              clear: both !important;
              margin: 0 auto !important;
              max-width: 420px !important; }
              .container table {
                width: 100% !important;
                border-collapse: collapse; }
              .container .masthead {
                padding: 35px 0;
                background: #379dbc;
                color: white; }
                .container .masthead h1 {
                  margin: 0 auto !important;
                  max-width: 90%;
                  text-transform: uppercase; }
              .container .content {
                background: white;
                padding: 30px 15px; }
                .container .content.footer {
                  background: none; }
                  .container .content.footer p {
                    margin-bottom: 0;
                    color: #888;
                    text-align: center;
                    font-size: 5px; }
                  .container .content.footer a {
                    color: #888;
                    text-decoration: none;
                    font-weight: bold; }
        </style>
    </head>
    <body>
        <table class="body-wrap" align="center" >
            <tr>
                <td class="container">
                <!-- Message start -->
            <table>
                <tr >
                    <td align="center" class="masthead" style="height:10px">
                        <h1>`+ tittle +`</h1>
                    </td>
                </tr>
                <tr>
                   <td align="center" class="content">
					<div style="margin-left:auto;margin-right:auto">
                     <pre style="text-align:left;"> 
						`+ content +`
                     </pre>
					<div>
                        <table>
                            <tr>
                                <td align="center">
                                        <a href="http://uwatch.juanpi.org" class="button">点击前往查看详情</a><br/>
                                           <p style="font-size:13px">当前只支持办公网络与pc端查看</p>
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
            </table>

        </td>
    </tr>
    <tr>
        <td class="container">
            <table>
                <tr>
                    <td class="content footer" align="center"  bgcolor="#E1E1E1">
                 <div style="font-family:Helvetica,Arial,sans-serif;font-size:13px;margin-top:-15px;color:#828282;text-align:center;line-height:120%;">
                  <div style="color:#828282;">
                   此邮件为系统自动发送，请勿回复 <br /> Copyright &copy; 2017 - 2018 uwatch.juanpi.org. All Rights Reserved<br />
                   卷皮网 技术保障部 工具平台组 技术支持：连线运维
                  </div>
                 </div>
                </td>
                </tr>
            </table>

        </td>
    </tr>
</table>
</body>
</html>
	`
	err := mail.Send()
	if err != nil {
		beego.Error("邮件发送失败，错误详情：",err.Error())
		return  err
	}else{
		beego.Info("邮件发送成功")
	}
	return nil

}

