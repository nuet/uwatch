package api

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"library/falcon/service"
)

type GetUserByName struct {
	Cnname string `json:"cnname"`
	Email string `json:"email"`
	Id int64 `json:"id"`
	Im string `json:"im"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Qq string `json:"qq"`
	Role int `json:"role"`
}

//用户组
type GetTeamByName struct {
	Creater int `json:"creater"`
	Creater_name string `json:"creater_name"`
	Id int 	`json:"id"`
	Name string `json:"name"`
	Resume string `json:"resume"`
	Users []GetUserByName `json:"users"`
}

type CreateTeamCode struct {
	Message string `json:"message"`
}

type GetTeamListData struct {
	Team TeamList `json:"team"`
	Create_name string `json:"create_name"`
	Users []GetUserByName `json:"users"`
}

type TeamList struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Resume string `json:"resume"`
	Creator int `json:"creator"`
}


func (c *Api) GetUserList() ([]GetUserByName, error) {
	c.SetRequestType("/user/users")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []GetUserByName{}
	if nil != rstErr {
		fmt.Println("failed to create the UserList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) SearchFalconUserName(nikiname string) (GetUserByName, error) {
	c.SetRequestType("/user/name/" + nikiname)
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := GetUserByName{}
	if nil != rstErr {
		fmt.Println("failed to create the GetUserByName, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateFalconUser(data map[string]interface{}) (GetUserByName, error) {
	dataStr, _ := json.Marshal(data)
	beego.Info("dataStory===>", dataStr)
	c.SetRequestType("/user/create")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := GetUserByName{}

	if nil != rstErr {
		fmt.Println("failed to create the CreateUser, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) UpdateFalconUser(data map[string]interface{},login string) (GetUserByName, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/user/update")
	c.SetRequestLogin(login)
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := GetUserByName{}

	if nil != rstErr {
		fmt.Println("failed to create the UpdateUser, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) SearchFalconTeamName(name string) (GetTeamByName, error) {
	c.SetRequestType("/team/name/" +name)
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := GetTeamByName{}
	if nil != rstErr {
		fmt.Println("failed to create the SearchFalconTeamName, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteUserId(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/admin/delete_user")
	rst, rstErr := c.Send("DELETE", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to delete DeleteUserId, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateFalconTeam(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/team")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the CreateFalconTeam, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) UpdateFalconTeam(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/team")
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the UpdateFalconTeam, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteTeamId(team_id int) (service.CommonSuccResponse, error) {
	c.SetRequestType(fmt.Sprintf("/team/%d", team_id))
	rst, rstErr := c.Send("DELETE", []byte{})
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to delete DeleteTeamId, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) SearchTeamList(team string) ([]GetTeamListData, error) {
	c.SetRequestType("/team?q=" + team)
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []GetTeamListData{}
	if nil != rstErr {
		fmt.Println("failed to search the SearchTeamList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}
	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}
	return rstObj, nil
}

func (c *Api) SetFalconAdmin(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/admin/change_user_role")
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the UpdateUser, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}