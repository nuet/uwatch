package api

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"net/url"
)

type EndpointList struct {
	ID int `json:"id"`
	Endpoint string `json:"endpoint"`
}

type CounterList struct {
	Counter string `json:"counter"`
	Endpoint_id int `json:"endpoint_id"`
}

type GraphHistory struct {
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
	dstype   string `json:"dstype"`
	Values   []HistoryList `json:"values"`
}

type HistoryList struct {
	Timestamp int `json:"timestamp"`
	Value float64 `json:"value"`
}

type Prams struct {
	Step                	 int         `json:"step"`
	Start_time             int            `json:"start_time"`
	Hostnames 					[]string 			`json:"hostnames"`
	End_time               int            `json:"end_time"`
	Counters 						[]string			`json:"counters"`
	Consol_fun					string 				`json:"consol_fun"`
}

func (c *Api) SearchEndpointList(endpoint string) ([]EndpointList, error) {
	c.SetRequestType("/graph/endpoint?q=" + endpoint)
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []EndpointList{}

	if nil != rstErr {
		fmt.Println("failed to search the EndpointList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) SearchCounterList(endpoint string) ([]CounterList, error) {
	c.SetRequestType("/graph/endpoint_counter?eid=" + endpoint)
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []CounterList{}
	if nil != rstErr {
		fmt.Println("failed to search the EndpointList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}
	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}
	return rstObj, nil
}

func (c *Api) SearchHistoryList(data map[string]interface{}) ([]GraphHistory, error) {
	dataStr, _ := json.Marshal(data)
	beego.Info("dataStory===>", dataStr)
	c.SetRequestType("/graph/history")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := []GraphHistory{}

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