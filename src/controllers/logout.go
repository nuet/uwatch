package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

const (
	// oa 登录
	LOGOUT_URL = "api/logout"
)


func (this *LogoutController) Get()  {
	webUrl := this.GetString("webUrl")
	oaAuthWeb := beego.AppConfig.String("oa_auth_web")
	logoutUrl := oaAuthWeb + LOGOUT_URL + "?jump_url=" + webUrl
	beego.Info(logoutUrl)
	this.Redirect(logoutUrl, 302)
	this.StopRun()
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "sucess"}
	this.ServeJSON()
	return
}