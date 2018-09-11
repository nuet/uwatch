package api

import (
	"encoding/json"
	"fmt"
	"library/cmdb/service"
)


type DescribeSearchProcResponse struct{
	service.CommonResponse
	Data SearchProcResponse	`json:"data"`
}

type SearchProcResponse struct {
	Count int 			`json:"count"`
	Info  []ProcResponse		`json:"info"`
}

type ProcResponse struct {
	BkBizId int			`json:"bk_biz_id"`
	BkProcessId int			`json:"bk_process_id"`
	BkProcessName string 		`json:"bk_process_name"`
	Port string			`json:"port"`
}


func (c *Api) SearchProc(ownerID string, bizID int, query map[string]interface{}) (*DescribeSearchProcResponse, error) {
	queryStr, _ := json.Marshal(query)
	c.SetRequestType(fmt.Sprintf("/proc/search/%s/%d", ownerID, bizID))
	rst, rstErr := c.Send("POST", queryStr)
	rstObj := &DescribeSearchProcResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the proc, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

type DescribeProcessBindModuleResponse struct{
	service.CommonResponse
	Data []ProcessBindModuleResponse	`json:"data"`
}

type ProcessBindModuleResponse struct {
	BkModuleName string		`json:"bk_module_name"`
	IsBind int			`json:"is_bind"`
	SetNum int 			`json:"set_num"`
}

func (c *Api) GetProcessBindModule(ownerID string, bizID, procID int) (*DescribeProcessBindModuleResponse, error) {
	c.SetRequestType(fmt.Sprintf("/proc/module/%s/%d/%d", ownerID, bizID, procID))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := &DescribeProcessBindModuleResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the ProcessBindModule, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil

}