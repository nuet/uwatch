package api

import (
	"encoding/json"
	"fmt"
	"library/falcon/service"
)


type HostGroupResponse struct {
	HGID 	int			`json:"id"`
	HGName 	string			`json:"grp_name"`
	CreateUser string		`json:"create_user"`
}

type BindTHResponse struct {
	HGID 	int			`json:"grp_id"`
	TplId 	int			`json:"tpl_id"`
	BindUser int64			`json:"bind_user"`
}


func (c *Api) SearchHostGroup() ([]HostGroupResponse, error) {
	c.SetRequestType("/hostgroup")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []HostGroupResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the HostGroup, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateHostGroup(data map[string]interface{}) (HostGroupResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/hostgroup")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := HostGroupResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the HostGroup, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) AddHostsToHostGroup(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/hostgroup/host")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to add hosts to the HostGroup, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteHostGroup(HGID int) (service.CommonSuccResponse, error) {
	c.SetRequestType(fmt.Sprintf("/hostgroup/%d", HGID))
	rst, rstErr := c.Send("DELETE", []byte{})
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to delete HostGroup, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) BindTemplateToHostGroup(data map[string]interface{}) (BindTHResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/hostgroup/template")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := BindTHResponse{}

	if nil != rstErr {
		fmt.Println("failed to add bind template to the HostGroup, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}