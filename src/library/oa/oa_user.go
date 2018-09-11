package oa

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
)
type UserQuery struct {
	Users []*FeUserInfo `json:"users"`
}

type FeUserInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	CnName  string `json:"cnname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Im      string `json:"im"`
	QQ      string `json:"qq"`
	Role    int    `json:"role"`
	Created string `json:"created"`
	TeamID  string `json:"teamid"`
}


func GetUserGroupByPhone(phones []string) []string {
	users := []string{}
	for _,phone := range phones{
		var userquery UserQuery
		feIp := beego.AppConfig.String("feIp")
		url := feIp + "/user/query/?query=" + phone
		req := httplib.Get(url)
		resp, _ := req.String()
		_ = json.Unmarshal([]byte(resp), &userquery)
		for _, s := range userquery.Users {
			if s.Phone == phone {
				users = append(users,s.Name)
			}
		}
	}
	return users
}



func GetMobile(user string) string {
	var userquery UserQuery
	feIp := beego.AppConfig.String("feIp")
	url := feIp + "/user/query/?query=" + user
	req := httplib.Get(url)

	resp, _ := req.String()
	_ = json.Unmarshal([]byte(resp), &userquery)
	for _, s := range userquery.Users {
		if s.Name == user {
			return s.Phone
		}
	}
	return ""
}


func GetUser(user string) (UserQuery, error) {
	var userquery UserQuery
	feIp := beego.AppConfig.String("feIp")
	url := feIp + "user/query/?query=" + user
	req := httplib.Get(url)
	beego.Debug(url)
	beego.Debug(user)

	resp, err := req.String()
	if err != nil {
		beego.Info(err)
		return userquery, err
	}
	beego.Debug(resp)
	err = json.Unmarshal([]byte(resp), &userquery)
	return userquery, err
}