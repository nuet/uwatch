	package dashboard

	import (
		"github.com/astaxie/beego"
		"github.com/astaxie/beego/httplib"
		"time"
		"encoding/json"
		"library/common"
		"library/falcon"
		"strings"
		"errors"
		"models"
		"controllers"

		"github.com/astaxie/beego/utils"
	)

	type GetGraphDataController struct {
		controllers.BaseRouter
	}
	type EndpointValues struct {
		Timestamp int     `json:"timestamp"`
		Value     float64 `json:"value"`
	}

	type EndpointRes struct {
		Endpoint  string           `json:"endpoint"`
		Counter   string           `json:"counter"`
		Step      int              `json:"step"`
		Dstype    string           `json:"dstype"`
		Values    []EndpointValues `json:"values"`
		Avg       string          `json:"avg"`
		Min       string          `json:"min"`
		Max       string          `json:"max"`
		Axislabel string      `json:"axislabel"`
		Dividend  string      `json:"dividend"`
		Title     string      `json:"title"`
	}

	type GetCounterList struct {
		Msg  string `json:"msg"`
		Data []string `json:"data"`
		Ok   bool    `json:"ok"`
	}

	func (c *GetGraphDataController) FindOneGraphData() {
		//拉取数据begin
		type EndpointPram struct {
			Endpoint string `json:"endpoint"`
			Counter  string `json:"counter"`
		}

		type Hosts struct {
			Query string  `json:"query"`
		}

		type Prams struct {
			Start             int            `json:"start"`
			End               int            `json:"end"`
			Cf                string         `json:"cf"`
			Endpoint_counters []EndpointPram `json:"endpoint_counters"`
		}

		type Counter struct {
			Id         int `json:"id"`
			Counter    string `json:counter`
			Counter_cn string `json:counter_cn`
			Status     string  `json:status`
		}

		requesturl := beego.AppConfig.String("falconHost") + "/graph/history"
		loc, _ := time.LoadLocation("Local")
		timeLayout := "2006-01-02 15:04:05"
		theTime, _ := time.ParseInLocation(timeLayout, time.Now().Format("2006-01-02 15:04:05"), loc)

		jsonMap := map[string]interface{}{}
		json.Unmarshal([]byte(c.Ctx.Input.RequestBody), &jsonMap)

		var fields []string
		var sortby []string
		var query = make(map[string]string)
		var limit int64 = 50
		var offset int64
		var times int64
		var status int64
		// fields: col1,col2,entity.col3
		if v := c.GetString("fields"); v != "" {
			fields = strings.Split(v, ",")
		}
		// limit: 10 (default is 10)
		if v, err := c.GetInt64("limit"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := c.GetInt64("offset"); err == nil {
			offset = v
		}
		// sortby: col1,col2
		if v := c.GetString("sortby"); v != "" {
			sortby = strings.Split(v, ",")
		}

		// query: {k:v,k:v}
		if v := c.GetString("query"); v != "" {
			if err := json.Unmarshal([]byte(v), &query); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}

		if v := c.GetString("times"); v != "" {
			if err := json.Unmarshal([]byte(v), &times); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}

		if v := c.GetString("status"); v != "" {
			if err := json.Unmarshal([]byte(v), &status); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}

		end := int(theTime.Unix())
		start := int(end - 3600 * common.GetInt(times))

		var co_fields []string
		var co_sortby []string
		var co_query = make(map[string]string)
		var co_limit int64 = 50
		var co_offset int64
		_, _, counter := models.GetAllUwCounterData(status, co_query, co_fields, co_sortby, co_offset, co_limit)
		_, _, l := models.ListAllGraphData(status, co_query, fields, sortby, offset, limit)
		beego.Info(counter)
		beego.Info("Query====>", query["Query"])

		q := strings.Split(query["Query"], "^")
		beego.Info("q===>", q[0], q[1])
		if common.GetString(q[0]) != "" {
			var ret_data []map[string]interface{}
			beego.Info(q[0])
			param := EndpointPram{common.GetString(q[0]), common.GetString(q[1])}
			cpuPrams := make([]EndpointPram, 0)
			cpuPrams = append(cpuPrams, param)
			prams := &Prams{start, end, "AVERAGE", cpuPrams}
			beego.Info("prams===>", prams)
			bi, err := json.Marshal(prams)

			if err != nil {
				beego.Info(err)
				return
			}
			endpointRes := []EndpointRes{}
			var req *httplib.BeegoHTTPRequest
			req = httplib.Post(requesturl)
			req.SetTimeout(5 * time.Second, 5 * time.Second)
			req.Body(bi)
			err = req.ToJSON(&endpointRes)
			if err != nil {
				beego.Info("拉取GRAPH失败", err)
				c.Ctx.Output.SetStatus(201)
				c.SetJson(200, "false", "false")
			} else {
				beego.Info("data===>", endpointRes)
				if endpointRes[0].Values != nil {
					for t, _ := range endpointRes {
						temp := map[string]interface{}{}
						for _, v := range l {
							if strings.Contains(common.GetString(v["Counter"]), common.GetString(q[1])){
								temp["counter"] = common.GetString(v["Counter_cn"])
								temp["avg"] = common.GetString(v["Avg"])
								temp["min"] = common.GetString(v["Min"])
								temp["max"] = common.GetString(v["Max"])
								temp["axislabel"] = common.GetString(v["Axislabel"])
								temp["dividend"] = common.GetString(v["Dividend"])
								temp["title"] = common.GetString(v["Counter_cn"])
								temp["values"] = endpointRes[t].Values
								ret_data = append(ret_data, temp)
							}
						}
					}
				}
				beego.Info("ret_data===>", ret_data)
				c.Ctx.Output.SetStatus(201)
				c.SetJson(200, ret_data, "succ")
			}
		} else {
			c.Ctx.Output.SetStatus(201)
			c.SetJson(200, "false", "false")
		}
	}

	//拉取0.2版本falcon数据
	func (c *GetGraphDataController) FindOneFalconData() {
		var query = make(map[string]string)
		var times int64
		// query: {k:v,k:v}
		if v := c.GetString("query"); v != "" {
			if err := json.Unmarshal([]byte(v), &query); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}

		if v := c.GetString("times"); v != "" {
			if err := json.Unmarshal([]byte(v), &times); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}
		beego.Info("Query====>", query["Query"])

		q := strings.Split(query["Query"], "^")
		beego.Info("q===>", q[0])
		var ret_data []map[string]interface{}
		if common.GetString(q[0]) != "" {
			loc, _ := time.LoadLocation("Local")
			timeLayout := "2006-01-02 15:04:05"
			theTime, _ := time.ParseInLocation(timeLayout, time.Now().Format("2006-01-02 15:04:05"), loc)

			end := int(theTime.Unix())
			start := int(end - 3600 * common.GetInt(times))
			var hostnames []string
			var counters []string
			h := strings.Split(q[0], ",")
			c := strings.Split(q[1], ",")
			hostnames = h
			counters = c
			counterList, _ := falcon.GetGraphHistory(start, hostnames, end, counters)

			type HistoryList struct {
				Timestamp int `json:"timestamp"`
				Value int `json:"value"`
			}
			type GraphHistory struct {
				Endpoint string `json:"endpoint"`
				Counter  string `json:"counter"`
				dstype   string `json:"dstype"`
				Values   []HistoryList `json:"values"`
			}

			var fields []string
			var sortby []string
			var f_query = make(map[string]string)
			var limit int64 = 200
			var offset int64
			_, _, l := models.GetAllUwCounter(f_query, fields, sortby, offset, limit)
			beego.Info("GetGraphHistory====>",counterList)


			for _, v := range counterList {
				for _, u := range l {
					if strings.Contains(common.GetString(v.Counter), common.GetString(u["Counter"])) {
						temp := map[string]interface{}{}
						temp["counter"] = common.GetString(u["Counter"])
						temp["title"] = common.GetString(u["Counter_cn"])
						temp["axislabel"] = common.GetString(u["Axislabel"])
						temp["dividend"] = common.GetString(u["Dividend"])
						temp["avg"] = common.GetString(u["Avg"])
						temp["values"] = v.Values
						ret_data = append(ret_data, temp)
					}
				}
			}

		} else {
			c.Ctx.Output.SetStatus(201)
			c.SetJson(200, "false", "false")
		}
		c.Ctx.Output.SetStatus(201)
		c.SetJson(200, ret_data, "succ")
	}
	func (c *GetGraphDataController) GetOneFalconData() {
		var query = make(map[string]string)
		var times int64
		// query: {k:v,k:v}
		if v := c.GetString("query"); v != "" {
			if err := json.Unmarshal([]byte(v), &query); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}

		if v := c.GetString("times"); v != "" {
			if err := json.Unmarshal([]byte(v), &times); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}
		beego.Info("Query====>", query["Query"])

		q := strings.Split(query["Query"], "^")
		beego.Info("oneoneone===>", q[0])
		var ret_data []map[string]interface{}
		if common.GetString(q[0]) != "" {
			loc, _ := time.LoadLocation("Local")
			timeLayout := "2006-01-02 15:04:05"
			theTime, _ := time.ParseInLocation(timeLayout, time.Now().Format("2006-01-02 15:04:05"), loc)

			end := int(theTime.Unix())
			start := int(end - 3600 * common.GetInt(times))
			var hostnames []string
			var counters []string
			h := strings.Split(q[0], ",")
			c := strings.Split(q[1], "^")
			hostnames = h
			counters = c
			counterList, _ := falcon.GetGraphHistory(start, hostnames, end, counters)
			beego.Info("counterList===>", counterList[0])

			type HistoryList struct {
				Timestamp int `json:"timestamp"`
				Value int `json:"value"`
			}
			type GraphHistory struct {
				Endpoint string `json:"endpoint"`
				Counter  string `json:"counter"`
				dstype   string `json:"dstype"`
				Values   []HistoryList `json:"values"`
			}


			for _, v := range counterList {
						temp := map[string]interface{}{}
						beego.Info("counterList==>", v.Endpoint)
						temp["counter"] = v.Counter
						temp["title"] = v.Endpoint
						temp["axislabel"] = ""
						temp["dividend"] = ""
						temp["avg"] = 1
						temp["values"] = v.Values
						ret_data = append(ret_data, temp)
				}
		} else {
			c.Ctx.Output.SetStatus(201)
			c.SetJson(200, "false", "false")
		}
		utils.Display("oneGraph===>", ret_data)
		c.Ctx.Output.SetStatus(201)
		c.SetJson(200, ret_data, "succ")
	}
	//拉取0.1版本的falcon
	func (c *GetGraphDataController) GetCounter() {
		var query = make(map[string]string)
		// query: {k:v,k:v}
		if v := c.GetString("query"); v != "" {
			if err := json.Unmarshal([]byte(v), &query); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}
		var req *httplib.BeegoHTTPRequest
		requesturl := beego.AppConfig.String("dashboard") + "/api/endpoints?q=" + common.GetString(query["Query"]) + "&tags=&limit=100"
		req = httplib.Get(requesturl)
		req.SetTimeout(5 * time.Second, 5 * time.Second)
		ret, err1 := req.String()
		if err1 != nil {
			beego.Info("拉取GRAPH失败", err1)
		}

		ProsserSite := beego.AppConfig.String("dashboard") + "/api/counters"
		var req_post *httplib.BeegoHTTPRequest
		req_post = httplib.Post(ProsserSite)
		req_post.SetTimeout(5*time.Second, 5*time.Second)
		req_post.Param("endpoints", common.GetString(query["Query"]))
		req_post.Param("q", common.GetString(query["Query"]))
		req_post.Param("limit","20")
		reqstr_post, _ := req_post.String()
		beego.Info("拉取Counter数据===>",reqstr_post)

		var dat GetCounterList
		err2 := json.Unmarshal([]byte(ret), &dat)
		if err2 != nil {
			beego.Info(err2)
		}

		var list map[string]interface{}
		err3 := req.ToJSON(&list)
		if err3 != nil {
			beego.Info("拉取GRAPH失败", err3)
		}
		var c_fields []string
		var c_sortby []string
		var c_query = make(map[string]string)
		var c_offset int64

		_, err, l := models.GetAllSearchGraphData(c_query, c_fields, c_sortby, c_offset, 200)
		beego.Info(err)
		var ret_data []map[string]interface{}
		for _, k := range dat.Data {
			//utils.Display("oo==>", o)
			for _, v := range l {
				temp := map[string]interface{}{}
				temp["label"] = k + "^" + common.GetString(v["Counter_cn"])
				temp["value"] = k + "^" + common.GetString(v["Counter"])
				ret_data = append(ret_data, temp)
			}
		}
		c.Ctx.Output.SetStatus(201)
		c.SetJson(200, dat.Data, "succ")
	}
	func (c *GetGraphDataController) GetGraphCounter() {
		var fields []string
		var sortby []string
		var f_query = make(map[string]string)
		var limit int64 = 200
		var offset int64
		var query []string
		if v := c.GetString("query"); v != "" {
			if err := json.Unmarshal([]byte(v), &query); err != nil {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
		}
		beego.Info("query+++++++===>", query)
		_, err, l := models.GetAllUwCounter(f_query, fields, sortby, offset, limit)
		beego.Info("script====>", l)
		var ret_data []map[string]interface{}
		for _, v := range l {
			temp := map[string]interface{}{}
			temp["name"] = "0.2版本->>" + common.GetString(v["Counter_cn"]) + "—可模糊检索"
			ret_data = append(ret_data, temp)
		}
		if err != nil {
			c.SetJson(400, err.Error(), "fail")
		} else {
			c.SetJson(200, ret_data, "succ")
		}
	}
