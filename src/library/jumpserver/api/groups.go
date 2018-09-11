package api

import (
	"encoding/json"
	"fmt"
)

const BASE_GROUPS_API_URL  = "/api/users/v1"

type Groups struct {
	Id 		string			`json:"id"`
	Users []string 	`json:"users"`
	Is_discard bool `json:"is_discard"`
	Discard_time string `json:"discard_time"`
	Name string `json:"name"`
	Comment string `json:"comment"`
	Date_created string `json:"date_created"`
	Created_by string 	`json:"created_by"`
}


func (c *Api) GetGroupsList() ([]Groups, error) {
	c.SetRequestType(BASE_GROUPS_API_URL + "/groups/")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Groups{}

	if nil != rstErr {
		fmt.Println("failed to search the GroupsList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) SearchGroups(GroupName string) ([]Groups, error) {
	c.SetRequestType(BASE_GROUPS_API_URL + fmt.Sprintf("/groups/?name=%s", GroupName))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Groups{}

	if nil != rstErr {
		fmt.Println("failed to search the GroupName, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateGroups(data map[string]interface{}) (Groups, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_GROUPS_API_URL + "/groups/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := Groups{}

	if nil != rstErr {
		fmt.Println("failed to create the CreateGroups, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) UpdateGroups(groupId string, data map[string]interface{}) (Groups, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_GROUPS_API_URL + fmt.Sprintf("/groups/%s/", groupId))
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := Groups{}

	if nil != rstErr {
		fmt.Println("failed to update the UpdateGroups, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteGroups(groupId string) error {
	c.SetRequestType(BASE_GROUPS_API_URL + fmt.Sprintf("/groups/%s/", groupId))
	_, rstErr := c.Send("DELETE", []byte{})
	if nil != rstErr {
		fmt.Println("failed to delete the DeleteGroups, error info is ", rstErr.Error())
		return rstErr
	}
	return nil
}