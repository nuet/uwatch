package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"strings"
	"library/oa"
	"errors"
	"runtime"
)

const (
	// oa 获取用户列表
	USER_LIST_URL = "api/user/list"
	DEPARTMENT_LIST_URL = "api/department/list"
)

var GetCache cache.Cache

func init() {
	GetCache, _ = cache.NewCache("memory", `{"interval":60}`)
}

type BaseRouter struct {
	beego.Controller
	User    *oa.UserInfo
}

func (this *BaseRouter) Get() {
	this.TplName = "index.tpl"
}

// Prepare implemented Prepare method for baseRouter.
func (c *BaseRouter) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                                                      //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")                               //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,UserToken,Content-Type,Authorization") //header的类型
	if c.Ctx.Input.IsOptions() {
		return
	}

	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			runtimec := runtime.Stack(buf, false)
			beego.Error("控制器错误:", panic_err, string(buf[0:runtimec]))
		}
	}()
	token := ""
	if ah := c.Ctx.Input.Header("Authorization"); ah != "" {
		if len(ah) > 5 && strings.ToUpper(ah[0:5]) == "TOKEN" {
			token = ah[6:]
			if (token != "") {
				userInfo, err := oa.GetUserInfoByToken(token)
				if(err == nil) {
					c.User = userInfo
				}
			}
		}
	}
}

func (this *BaseRouter) AllowCross() {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                                                      //允许访问源
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")                               //允许post访问
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,UserToken,Content-Type,Authorization") //header的类型
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	this.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}
func (c *BaseRouter) Options() {
	c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

func (this *BaseRouter) SetJson(code int, data interface{}, Msg string) {
	if (code == 0) {
		if (Msg == "") {
			Msg = "sucess"
		}
		this.Data["json"] = map[string]interface{}{"code":code, "msg":Msg, "data":data}
		this.ServeJSON()
		this.StopRun()
	} else {
		this.Data["json"] = map[string]interface{}{"code":code, "msg":Msg, "data":data}
		this.ServeJSON()
		this.StopRun()
	}
}


// 将用户名从请求头中取出
func (c *BaseRouter) GetUserName() (string, error) {
	token := ""
	if ah := c.Ctx.Input.Header("Authorization"); ah != "" {
		if len(ah) > 5 && strings.ToUpper(ah[0:5]) == "TOKEN" {
			token = ah[6:]
			if (token != "") {
				userInfo, err := oa.GetUserInfoByToken(token)
				if(err==nil){
					beego.Info(userInfo)
					return userInfo.UserName,nil
				}else{
					c.SetJson(50014,"","Token不存在或已失效")
					beego.Error(err)
					return "",err
				}
			}
		}
	}

	return "", errors.New("no token")
}