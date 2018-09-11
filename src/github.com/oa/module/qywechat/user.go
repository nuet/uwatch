//企业号成员管理
package qywechat

import (
	"fmt"
)

type User struct {
	Userid		string	`json:"userid"`
	Name		string	`json:"name,omitempty"`
	Department 	[]int	`json:"department,omitempty"`
	Position 	string	`json:"position,omitempty"`
	Mobile		string	`json:"mobile,omitempty"`
	Gender		string	`json:"gender,omitempty"`
	Email		string	`json:"email,omitempty"`
	Weixinid	string	`json:"weixinid,omitempty"`
	Enable		int 	`json:"enable"`
	AvatarMediaid	string	`json:"avatar_mediaid,omitempty"`
	Status		int	`json:"-"`
	Extattr		interface{} `json:"extattr"`
}

//企业号成员添加
func (app *App) CreateUser(data User) (err error){
	result := &Error{}
	err = app.Post("/user/create",data, result)
	if err!=nil{
		fmt.Println("添加微信企业用户失败，"+err.Error())
		return
	}
	if result.ErrCode != ErrCodeOK{
		err = result
		fmt.Println("添加微信企业用户失败，"+err.Error())
		return
	}
	return nil
}

//企业号成员删除
func (app *App) DeleteUser(userid string) (err error){
	result := &Error{}
	query := map[string]string{"userid":userid}
	err = app.Get("/user/delete",query, result)

	if err!=nil{
		fmt.Println("删除微信企业用户失败，"+err.Error())
		return
	}
	if result.ErrCode != ErrCodeOK{
		err = result
		fmt.Println("删除微信企业用户失败，"+err.Error())
		return
	}
	return nil
}

//企业号成员更新
func (app *App) UpdateUser(data User) (err error){
	result := &Error{}
	err = app.Post("/user/update",data, result)
	if err!=nil{
		fmt.Println("更新微信企业用户失败，"+err.Error())
		return
	}
	if result.ErrCode != ErrCodeOK{
		err = result
		fmt.Println("更新微信企业用户失败，"+err.Error())
		return
	}
	return nil
}

//获取企业号成员
func (app *App) GetUser(userid string) (user User,err error){
	var result struct {
		Error
		User
	}
	query := map[string]string{"userid":userid}
	err = app.Get("/user/get",query,&result)
	if err!=nil{
		fmt.Println("获取微信企业用户失败，"+err.Error())
		return
	}
	if result.ErrCode != ErrCodeOK{
		err = &result.Error
		fmt.Println("获取微信企业用户失败，"+err.Error())
		return
	}
	user = result.User
	return
}