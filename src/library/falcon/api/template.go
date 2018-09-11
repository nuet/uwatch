package api

import (
	"encoding/json"
	"fmt"
	"library/falcon/service"
)

type TemplateListResponse struct {
	Templates	[]TemplateList `json:"templates"`
}

type TemplateList struct {
	Template	Template	`json:"template"`
	ParentName	string		`json:"parent_name"`
}

type Template struct {
	TplId 		int		`json:"id"`
	TplName 	string		`json:"tpl_name"`
	ParentId 	int		`json:"parent_id"`
	ActionId 	int		`json:"action_id"`
	CreateUser 	string		`json:"create_user"`
}

type Action struct {
	Id 		int		`json:"id"`
	Uic 		string		`json:"uic"`
}

type Strategy struct {
	Id 		int		`json:"id"`
	TplId 		int		`json:"tpl_id"`
	Metric		string		`json:"metric"`
	Tags		string		`json:"tags"`
	MaxStep		int		`json:"max_step"`
	Priority 	int		`json:"priority"`
	Note    	string		`json:"note"`
	Func 		string		`json:"func"`
	Op      	string		`json:"op"`
	RightValue   	string		`json:"right_value"`
}

type TplInfo struct {
	Action 		Action		`json:"action"`
	Template 	Template	`json:"template"`
	Stratges 	[]Strategy	`json:"stratges"`
	ParentName 	string		`json:"parent_name"`
}


func (c *Api) CreateTemplate(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/template")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the Template, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetTemplateList() (TemplateListResponse, error) {
	c.SetRequestType("/template")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := TemplateListResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the TemplateList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetTplInfoById(tplId int) (TplInfo, error) {
	c.SetRequestType(fmt.Sprintf("/template/%d", tplId))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := TplInfo{}

	if nil != rstErr {
		fmt.Println("failed to get the TemplateInfo, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateTplActoin(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/template/action")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the TemplateAction, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateStrategy(data map[string]interface{}) (service.CommonSuccResponse, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/strategy")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to create the Strategy, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteTemplate(tplId int) (service.CommonSuccResponse, error) {
	c.SetRequestType(fmt.Sprintf("/template/%d", tplId))
	rst, rstErr := c.Send("DELETE", []byte{})
	rstObj := service.CommonSuccResponse{}

	if nil != rstErr {
		fmt.Println("failed to delete Template, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}