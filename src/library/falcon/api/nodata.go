package api

import (
	"encoding/json"
	"fmt"
)


type NoData struct {
	Id 		int		`json:"id"`
	Name 		string		`json:"name"`
	Obj 		string		`json:"obj"`
	ObjType 	string		`json:"obj_type"`
	Metric 		string		`json:"metric"`
	Tags 		string		`json:"tags"`
	Dstype 		string		`json:"dstype"`
	Step 		int		`json:"step"`
	Mock 		int		`json:"mock"`
	Creator 	string		`json:"creator"`
}

type NoDataReq struct {
	Name 		string 		`json:"name"`
	Obj  		string 		`json:"obj"`
	ObjType 	string 		`json:"obj_type"`
	Metric  	string  	`json:"metric"`
	Tags    	string 		`json:"tags"`
	DsType  	string 		`json:"dstype"`
	Step    	int     	`json:"step"`
	Mock    	float64 	`json:"mock"`
}


func (c *Api) CreateNodata(data NoDataReq) (NoData, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType("/nodata/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := NoData{}

	if nil != rstErr {
		fmt.Println("failed to create the NoData, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetNodataList() ([]NoData, error) {
	c.SetRequestType("/nodata")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []NoData{}

	if nil != rstErr {
		fmt.Println("failed to search the NodataList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}