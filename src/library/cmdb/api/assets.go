package api

import (
	"encoding/json"
	"fmt"
	"library/cmdb/service"
)

type DescribesearchBusinessResponse struct{
	service.CommonResponse
	Data SearchBusinessResponse	`json:"data"`
}

type SearchBusinessResponse struct {
	Count int 			`json:"count"`
	Info  []BusinessResponse	`json:"info"`
}

type BusinessResponse struct {
	BkBizId int			`json:"bk_biz_id"`
	BkBizName string		`json:"bk_biz_name"`
	BkBizMaintainer string		`json:"bk_biz_maintainer"`
	BkBizDeveloper string		`json:"bk_biz_developer"`
}

func (c *Api) SearchBusiness(ownerID string, query map[string]interface{}) (*DescribesearchBusinessResponse, error) {
	queryStr, _ := json.Marshal(query)
	c.SetRequestType(fmt.Sprintf("/biz/search/%s", ownerID))
	rst, rstErr := c.Send("POST", queryStr)
	rstObj := &DescribesearchBusinessResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the biz, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

type DescribeSearchSetResponse struct{
	service.CommonResponse
	Data SearchSetResponse		`json:"data"`
}

type SearchSetResponse struct {
	Count int 			`json:"count"`
	Info  []SetResponse		`json:"info"`
}

type SetResponse struct {
	BkBizId int			`json:"bk_biz_id"`
	BkSetId int			`json:"bk_set_id"`
	BkSetName string		`json:"bk_set_name"`
}


func (c *Api) SearchSet(ownerID string, bizID int, query map[string]interface{}) (*DescribeSearchSetResponse, error) {
	queryStr, _ := json.Marshal(query)
	c.SetRequestType(fmt.Sprintf("/set/search/%s/%d", ownerID, bizID))
	rst, rstErr := c.Send("POST", queryStr)
	rstObj := &DescribeSearchSetResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the set, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

type DescribeSearchModuleResponse struct{
	service.CommonResponse
	Data SearchModuleResponse	`json:"data"`
}

type SearchModuleResponse struct {
	Count int 			`json:"count"`
	Info  []ModuleResponse		`json:"info"`
}

type ModuleResponse struct {
	BkBizId int			`json:"bk_biz_id"`
	BkSetId int			`json:"bk_set_id"`
	BkModuleId int 			`json:"bk_module_id"`
	BkModuleName string		`json:"bk_module_name"`
	Operator string			`json:"operator"`
	BkBakOperator string		`json:"bk_bak_operator"`
}


func (c *Api) SearchModule(ownerID string, bizID, setID int, query map[string]interface{}) (*DescribeSearchModuleResponse, error) {
	queryStr, _ := json.Marshal(query)
	c.SetRequestType(fmt.Sprintf("/module/search/%s/%d/%d", ownerID, bizID, setID))
	rst, rstErr := c.Send("POST", queryStr)
	rstObj := &DescribeSearchModuleResponse{}

	if nil != rstErr {
		fmt.Println("failed to search the module, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

type DescribeSearchHostsResponse struct{
	service.CommonResponse
	Data SearchHostsResponse	`json:"data"`
}

type SearchHostsResponse struct {
	Count int 			`json:"count"`
	Info  []HostsResponse		`json:"info"`
}

type HostsResponse struct {
	Module []Modules `json:"module"`
	Host Hosts			`json:"host"`
}

type Modules struct {
	TopModuleName string `json:"TopModuleName"`
	Bk_module_id int `json:"bk_module_id"`
	Bk_module_name string `json:"bk_module_name"`
}

type Hosts struct {
	BkHostInnerip string		`json:"bk_host_innerip"`
	BkHostName string		`json:"bk_host_name"`
	BkAgentStatus string `json:"bk_agent_status"`
	BkAgentUpdateTime string `json:"bk_agent_update_time"`
	BkAgentVersion string `json:"bk_agent_version"`
	BkAssetId string `json:"bk_asset_id"`
	BkBakOperator string `json:"bk_bak_operator"`
	BkComment string `json:"bk_comment"`
	//BkCpu int `json:"bk_cpu"`
	//BkCpuMhz int `json:"bk_cpu_mhz"`
	BkCpuModule string `json:"bk_cpu_module"`
	BkCurrentStatus string `json:"bk_current_status"`
	//BkDisk int `json:"bk_disk"`
	//BkHostId int `json:"bk_host_id"`
	BkHostManageip string `json:"bk_host_manageip"`
	BkHostOuterip string `json:"bk_host_outerip"`
	BkHostType string `json:"bk_host_type"`
	BkLanGateway string `json:"bk_lan_gateway"`
	BkLanMask string `json:"bk_lan_mask"`
	BkLevel string `json:"bk_level"`
	BkMac string `json:"bk_mac"`
	BkManageGateway string `json:"bk_manage_gateway"`
	BkManageMask string `json:"bk_manage_mask"`
	BkManufacturer string `json:"bk_manufacturer"`
	//BkMem int `json:"bk_mem"`
	BkOsBit string `json:"bk_os_bit"`
	BkOsName string `json:"bk_os_name"`
	BkOsType string `json:"bk_os_type"`
	BkOsVersion string `json:"bk_os_version"`
	BkOuterGateway string `json:"bk_outer_gateway"`
	BkOuterMac string `json:"bk_outer_mac"`
	BkOuterMask string `json:"bk_outer_mask"`
	BkProductName string `bk_productName`
	BkServiceTerm string `json:"bk_service_term"`
	BkSla string `json:"bk_sla"`
	BkSn string `json:"bk_sn"`
	BkStatus string `json:"bk_status"`
	BkUuid string `json:"bk_uuid"`
	Buytime string `json:"buytime"`
	CreateTime string `json:"create_time"`
	Expiretime string `json:"expiretime"`
	Firstusetime string `json:"firstusetime"`
	ImportFrom string `json:"import_from"`
	Operator string `json:"operator"`
	Inputtime string `json:"inputtime"`
	Outputtime string `json:"outputtime"`
	BkIspName string `json:"bk_isp_name"`
}

func (c *Api) SearchHosts(query map[string]interface{}) (*DescribeSearchHostsResponse, error) {
	queryStr, _ := json.Marshal(query)
	c.SetRequestType("/hosts/search")
	rst, rstErr := c.Send("POST", queryStr)
	rstObj := &DescribeSearchHostsResponse{}
	if nil != rstErr {
		fmt.Println("failed to search the hosts, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}