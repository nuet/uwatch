package dashboard

import (
	"encoding/json"
	"errors"
	"controllers"
	"library/cmdb"
	"library/falcon"
	"strings"
	"library/common"
	"github.com/astaxie/beego/utils"
)

type GetCmdbHostController struct {
	controllers.BaseRouter
}

func (c *GetCmdbHostController) GetCmdbHostList() {
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	// query: {k:v,k:v}
	if v := c.GetString("query"); v != "" {
		if err := json.Unmarshal([]byte(v), &query); err != nil {
			c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
		}
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("page"); err == nil {
		if v < 1 {
			v = 1
		}
		start := (v - 1) * limit
		offset = start
	}
	q := strings.Split(query["Query"], "^")

	cmdb_temp, _, _ := cmdb.GetCmdbHosts(common.GetString(q[0]), 1000, 0)
	utils.Display("cmdb_temp==>", offset)
	//先拉取falcon endpoint，再通过名称拉取cmdb
	res, err := falcon.GetSearchEndpoint(q[0])
	var list []map[string]interface{}
  total := len(res)
	ptotal := 0
	for k, _ := range res {
		if common.GetInt(k) >= common.GetInt(offset) && common.GetInt(ptotal) < common.GetInt(limit) {
			temp := map[string]interface{}{}
			if _, ok := cmdb_temp[res[k].Endpoint]; ok {
				temp["title"] = cmdb_temp[res[k].Endpoint]["title"]
				temp["BkHostName"] = cmdb_temp[res[k].Endpoint]["BkHostName"]
				temp["BkHostInnerip"] = cmdb_temp[res[k].Endpoint]["BkHostInnerip"]
				temp["BkHostOuterip"] = cmdb_temp[res[k].Endpoint]["BkHostOuterip"]
				temp["BkHostType"] = cmdb_temp[res[k].Endpoint]["BkHostType"]
				temp["Operator"] = cmdb_temp[res[k].Endpoint]["Operator"]
				temp["BkBakOperator"] = cmdb_temp[res[k].Endpoint]["BkBakOperator"]
				temp["BkManufacturer"] = cmdb_temp[res[k].Endpoint]["BkManufacturer"]
				temp["BkOsName"] = cmdb_temp[res[k].Endpoint]["BkOsName"]
				temp["BkOsBit"] = cmdb_temp[res[k].Endpoint]["BkOsBit"]
				temp["BkCpuModule"] = cmdb_temp[res[k].Endpoint]["BkCpuModule"]
				temp["BkIspName"] = cmdb_temp[res[k].Endpoint]["BkIspName"]
				temp["BkStatus"] = cmdb_temp[res[k].Endpoint]["BkStatus"]
				temp["BkCurrentStatus"] = cmdb_temp[res[k].Endpoint]["BkCurrentStatus"]
				temp["BkProductName"] = cmdb_temp[res[k].Endpoint]["BkProductName"]
				temp["Module"] = cmdb_temp[res[k].Endpoint]["Module"]
				ptotal++
			} else {
				temp["title"] = res[k].Endpoint
				temp["BkHostName"] = res[k].Endpoint
				temp["BkHostInnerip"] = "暂无"
				temp["BkHostOuterip"] = "暂无"
				temp["BkHostType"] = "暂无"
				temp["Operator"] = "暂无"
				temp["BkBakOperator"] = "暂无"
				temp["BkManufacturer"] = "暂无"
				temp["BkOsName"] = "暂无"
				temp["BkOsBit"] = "暂无"
				temp["BkCpuModule"] = "暂无"
				temp["BkIspName"] = "暂无"
				temp["BkStatus"] = "暂无"
				temp["BkCurrentStatus"] = "暂无"
				temp["BkProductName"] = "暂无"
				temp["Module"] = "暂无"
				ptotal++
			}
			counterList, _ := falcon.GetAllEndpointCounter(common.GetString(res[k].ID))
			var counter []map[string]string
			for k, c := range counterList {
				t := map[string]string{}
				t["name"] = c.Counter
				t["title"] = common.GetString(temp["title"])
				t["id"] = common.GetString(k + 1)
				counter = append(counter,t)
			}
			temp["Counter"] = counter
			temp["length"] = len(counter)
			list = append(list, temp)
		}
	}

	//cmdb_hostgroups, total, err  := cmdb.GetCmdbHosts(common.GetString(q[0]), limit, offset)

	if err != nil {
		c.SetJson(200, err.Error(), "fail")
	} else {
		c.Ctx.Output.SetStatus(201)
		req := make(map[string]interface{})
		req["list"] = list
		req["total"] = total
		c.SetJson(200, req, "succ")
	}
}

