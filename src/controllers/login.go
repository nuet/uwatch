package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"library/oa"
	"encoding/json"
	"net/url"
	"library/common"
	"time"
)

type LoginController struct {
	beego.Controller
}

const (
	// oa 登录
	LOGIN_URL = "api/login"
)

var Cache cache.Cache

func init() {
	Cache, _ = cache.NewCache("memory", `{"interval":60}`)
}

func (this *LoginController) Get()  {
	oaAuthWeb := beego.AppConfig.String("oa_auth_web")
	tokenValue := this.GetString("token")
	webUrl := this.GetString("webUrl")
	if !common.Empty(webUrl) {
		Cache.Put("webUrl",webUrl,120*time.Second)
	}
	requestHost := beego.AppConfig.String("requestHost")
	if tokenValue == "" {
		loginUrl := oaAuthWeb + LOGIN_URL + "?jump_url=" +  requestHost + this.Ctx.Input.URI()
		beego.Info("PC OA login url: ", loginUrl)
		this.Redirect(loginUrl, 302)
		this.StopRun()
	}
	// 用token向OA获取用户信息
	userInfo, err := oa.GetUserInfo(tokenValue)
	if err != nil {
		beego.Error("Get user info err: ", err)
		loginUrl := oaAuthWeb + LOGIN_URL + "?jump_url=" + requestHost + this.Ctx.Input.URI()
		this.Redirect(loginUrl, 302)
		this.StopRun()
	}

	resUserInfo:=map[string] interface{}{"user":userInfo,"login":true}
	userInfoJson,err:=json.Marshal(resUserInfo)
	this.Ctx.SetCookie("juanpi_uwatch_user_info", url.QueryEscape(string(userInfoJson)),3600*24*2 , "/")
	loginUrl := requestHost
	if Cache.IsExist("webUrl") {
		loginUrl = Cache.Get("webUrl").(string)
	}

	this.Redirect(loginUrl, 302)
	this.StopRun()
}