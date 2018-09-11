package dashboard

import (
	"encoding/json"
	"errors"
	"strings"
	"controllers"
	"library/falcon"
	"github.com/astaxie/beego"
	"library/common"
	"models"
)

type GetFalconCounterController struct {
	controllers.BaseRouter
}

type CounterList struct {
	Counter string `json:"counter"`
	Endpoint_id string `json:"endpoint_id"`
}
func (c *GetFalconCounterController) GetFalconCounter() {
	var query = make(map[string]string)
	// query: {k:v,k:v}
	if v := c.GetString("query"); v != "" {
		if err := json.Unmarshal([]byte(v), &query); err != nil {
			c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
		}
	}
	q := strings.Split(query["Query"], "^")
	var fields []string
	var sortby []string
	var f_query = make(map[string]string)
	var limit int64 = 200
	var offset int64
	_, _, l := models.GetAllUwCounter(f_query, fields, sortby, offset, limit)
	var ret_data []map[string]interface{}
		//拉取falcon_0.2 endpoint
		res, _ := falcon.GetSearchEndpoint(q[0])
		for _, v := range res {
			counterList, _ := falcon.GetAllEndpointCounter(common.GetString(v.ID))

			temp := map[string]interface{}{}
			Counter := ""
			counter := ""
			for _, c := range counterList {
				for j, u := range l {
					if strings.Contains(common.GetString(c.Counter), common.GetString(u["Counter"])) {
						beego.Info("counterlistCom===>",common.GetString(c.Counter), common.GetString(u["Counter"]))
						if j == 0 {
							Counter = common.GetString(c.Counter)
						} else {
							Counter = Counter + "," + common.GetString(c.Counter)
						}
						//if common.GetString(c.Counter) == "cpu.user" || common.GetString(c.Counter) == "cpu.user" || common.GetString(c.Counter) == "df.statistics.used.percent" {
						//	if j == 0 {
						//		counter = common.GetString(c.Counter)
						//	} else {
						//		counter = counter + "," + common.GetString(c.Counter)
						//	}
						//}
					}
				}
			}
			//counter = "cpu.user,mem.memfree.percent,df.statistics.used.percent"
			counter = "load.1min,load.5min,load.15min"
			temp["value"] = v.Endpoint + "^" + Counter
			temp["counter"] = v.Endpoint + "^" + counter
			ret_data = append(ret_data, temp)
		}
	beego.Info("ret_data==>", ret_data)

	c.Ctx.Output.SetStatus(201)
	c.SetJson(200, ret_data, "succ")
}
