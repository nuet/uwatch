package api

import (
	"encoding/json"
	"fmt"
)

const BASE_USERS_API_URL  = "/api/users/v1"

type Users struct {
	Id 		string			`json:"id"`
	Groups_display string `json:"groups_display"`
	Groups []string `json:"groups"`
	Last_login string `json:"last_login"`
	Is_active bool `json:"is_active"`
	Date_joined string `json:"date_joined"`
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Avatar string `json:"avatar"`
	Wechat string `json:"wechat"`
	Phone string `json:"phone"`
	Otp_level int `json:"otp_level"`
	Comment string `json:"comment"`
	Is_first_login bool `json:"is_first_login"`
	Date_expired string `json:"date_expired"`
	Created_by string `json:"created_by"`
	Get_role_display string `json:"get_role_display"`
	Is_valid bool `json:"is_valid"`
}


func (c *Api) GetUsersList() ([]Users, error) {
	c.SetRequestType(BASE_USERS_API_URL + "/users/")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Users{}

	if nil != rstErr {
		fmt.Println("failed to search the GetUsersList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) SearchUsers(UserName string) ([]Users, error) {
	c.SetRequestType(BASE_USERS_API_URL + fmt.Sprintf("/users/?username=%s", UserName))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Users{}

	if nil != rstErr {
		fmt.Println("failed to search the SearchUsers, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateUsers(data map[string]interface{}) (Users, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_USERS_API_URL + "/users/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := Users{}

	if nil != rstErr {
		fmt.Println("failed to create the CreateUsers, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) UpdateUsers(userId string, data map[string]interface{}) (Users, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_USERS_API_URL + fmt.Sprintf("/users/%s/", userId))
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := Users{}

	if nil != rstErr {
		fmt.Println("failed to update the UpdateUsers, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteUsers(userId string) error {
	c.SetRequestType(BASE_USERS_API_URL + fmt.Sprintf("/users/%s/", userId))
	_, rstErr := c.Send("DELETE", []byte{})
	if nil != rstErr {
		fmt.Println("failed to delete the DeleteUsers , error info is ", rstErr.Error())
		return rstErr
	}
	return nil
}