package qywechat

import (
	"fmt"
	"time"

	"github.com/oa/module/request"
)

const cfgPath = "conf/wechat.ini"

type AccessToken struct {
	Error
	Token   string `json:"access_token" ini:"token"`
	Expires int64  `json:"expires_in" ini:"expires"`
}

/**
 * 获取access_token
 */
func (app *App) GetToken() string {
	// 如果获取token失败或者token已失效则重新获取token
	if app.AccessToken.Validate() == false {
		fmt.Println("Token已失效，将重新获取Token")
		err := app.RequestToken()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	}

	return app.AccessToken.Token
}

func (app *App) RequestToken() error {
	err := request.GetJSON(baseUrl+"/gettoken?corpid="+app.CorpId+"&corpsecret="+app.CorpSecret, app.AccessToken)
	if err != nil {
		return err
	} else {
		if app.AccessToken.ErrCode != 0 {
			fmt.Printf("%s\n", app.AccessToken.Error.Error())
		} else {
			app.AccessToken.Expires = time.Now().Unix() + 7200
		}
		return nil
	}
}

/**
 * 验证token是否有效
 */
func (accessToken *AccessToken) Validate() bool {
	return (accessToken.Token != "") && (accessToken.Expires > time.Now().Unix())
}
