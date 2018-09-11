package oa

import (
	"crypto/tls"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"models"
	"encoding/json"
)
var GetCache cache.Cache
const (
	// oa 检查token并返回用户信息
	CHECK_TOKEN_URL = "api/check"
)
func init() {
	GetCache, _ = cache.NewCache("memory", `{"interval":60}`)
}
// 返回用户信息
type UserResp struct {
	Code       int        `json:"code"`
	UserData UserData 	`json:"data"`
	Message    string     `json:"msg"`
}
// OA 的用户信息, 用户名为汉字
type UserData struct {
	Username        string `json:"username"`
	Departmentid    string `json:"departmentid"`
	Email           string `json:"email"`
	Rtx             string `json:"rtx"`
}

type UserDetail struct {
	UserName string `json:"name"`
}


type UserInfo struct {
	UserName     string  `json:"name"`
	Departmentid string  `json:"departmentid"`
	Email        string  `json:"email"`
	Rtx          string  `json:"rtx"`
	Token string  `json:"token"`
}

// 用token向 OA请求获取用户信息
func GetUserInfo(tokenValue string) (*UserInfo, error) {
	var err error = nil
	oaAuthIp := beego.AppConfig.String("oa_auth_web")
	getUserInfoUrl := oaAuthIp + CHECK_TOKEN_URL + "?token=" + tokenValue
	beego.Info("getUserInfoUrl: ", getUserInfoUrl)
	userResp := UserResp{}
	for i := 0; i < 3; i++ {
		req := httplib.Get(getUserInfoUrl)
		req.SetTimeout(8*time.Second, 8*time.Second)
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		err = req.ToJSON(&userResp)
		if err == nil {
			reqStr, _ := req.String()
			beego.Info("reqStr: ", reqStr)
			break
		}
	}
	if err != nil {
		beego.Error("Send OA check token err: ", err)
		return nil, err
	}
	beego.Info("userResp: ", userResp)

	userInfo := UserInfo{}
	userInfo.UserName = userResp.UserData.Username
	userInfo.Departmentid = userResp.UserData.Departmentid
	userInfo.Email = userResp.UserData.Email
	userInfo.Rtx = userResp.UserData.Rtx
	userInfo.Token=tokenValue
	beego.Info("userInfo: ", userInfo)

	token := models.UserToken{}
	token.UserName = userInfo.UserName
	token.Token =tokenValue
	token.Expires = time.Now().Add(1 * time.Hour)
	userInfoJson,_:=json.Marshal(userInfo)
	token.Info=string(userInfoJson)
	models.AddUserToken(&token)

	go ExpiresTokenDEL()
	return &userInfo, nil
}


func ExpiresTokenDEL() {
	o := orm.NewOrm()
	sql := "DELETE FROM `user_token` WHERE `expires` < NOW()"
	_, err := o.Raw(sql).Exec()
	if err != nil {
		beego.Error(" 删除过期token错误:", err.Error())
	}

}
func GetUserInfoByToken(tokenValue string) (*UserInfo, error) {
	userInfo := UserInfo{}
	o := orm.NewOrm()
	var UesrToken []models.UserToken
	ss, err := o.Raw("SELECT * FROM `user_token`  WHERE `expires` > NOW() AND token=? ", tokenValue).QueryRows(&UesrToken)
	if err == nil && ss > 0 {
		beego.Info("11111111",tokenValue)
		err:=json.Unmarshal([]byte(UesrToken[0].Info),&userInfo)
		if(err != nil){
			return nil, err
		}else{
			return &userInfo, nil
		}
	}
	return nil, fmt.Errorf("no token")
}