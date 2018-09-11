package dashboard

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/httplib"
	"library/common"
	"strings"
	"encoding/json"
	"time"
)

// UwCounterController operations for UwCounter
type UwMonitorController struct {
	beego.Controller
}

type Alarm struct {
	Time 	string
	Status 	string
	Level 	string
	Name 	string
	Flag 	string
}

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

// URLMapping ...
func (c *UwMonitorController) URLMapping() {
	c.Mapping("Alarm", c.Alarm)
	c.Mapping("HgGraph", c.HgGraph)
}

// Usage ...
// @Title Usage
// @Description CPU MEM IO Usage
// @Param
// @Success 200 :OK
// @Failure 403 empty
// @router /alarm [get]
func (c *UwMonitorController) Alarm() {
	o := orm.NewOrm()
	o.Using("duty")
	var list []orm.Params
	var alarms []Alarm
	num, err := o.Raw(`SELECT REPLACE(REPLACE(content, CHAR(10), '|'), CHAR(13), '|') as content, type FROM alarm_events ORDER BY id DESC LIMIT 10`).Values(&list)
	if num > 0 && err == nil {
		for _, item := range list {
			var alarm Alarm
			contents := strings.Split(common.GetString(item["content"]), "|")
			if common.GetString(item["type"]) == "zabbix" {
				if len(contents) < 9 {
					continue
				}
				alarm.Name = strings.Split(contents[0], ":")[1]
				alarm.Flag = strings.Split(contents[5], ":")[1]
				alarm.Level = strings.Split(contents[4], ":")[1]
				alarm.Status = strings.Split(contents[8], ":")[1]
				alarm.Time = common.SubString(strings.Split(contents[3], " ")[1], 0, 5)
			} else {
				if contents[0] == "Open-Falcon监控报警" {
					alarm.Name = strings.Split(contents[1], ":")[1]
					alarm.Flag = strings.Split(contents[5], ":")[1]
					alarm.Level = strings.Split(contents[3], ":")[1]
					alarm.Status = strings.Split(contents[4], ":")[1]
					alarm.Time = common.SubString(strings.Split(contents[11], " ")[1], 0, 5)
				} else {
					continue
				}
			}
			alarms = append(alarms, alarm)
		}
	}
	c.Data["json"] = map[string]interface{}{"code":200, "msg":"succ", "data":alarms}
	c.ServeJSON()
	c.StopRun()
}


// HgGraph ...
// @Title HgGraph
// @Description huiguo graph data
// @Param
// @Success 200 :OK
// @Failure 403 empty
// @router /hgGraph [get]
func (c *UwMonitorController) HgGraph() {
	requesturl := beego.AppConfig.String("falconHost") + "/graph/history"
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	theTime, _ := time.ParseInLocation(timeLayout, time.Now().Format("2006-01-02 15:04:05"), loc)
	end := int(theTime.Unix())
	start := int(end - 3600)
	hgParams := []EndpointPram{
		{"api.huiguo.net", "intraffic"},
		{"api.huiguo.net", "outtraffic"},
		{"api.huiguo.net", "connum"},
	}
	prams := &Prams{start, end, "AVERAGE", hgParams}
	bi, err := json.Marshal(prams)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":200, "msg":"fial", "data":err.Error()}
		c.ServeJSON()
		c.StopRun()
	}
	endpointRes := []EndpointRes{}
	var req *httplib.BeegoHTTPRequest
	req = httplib.Post(requesturl)
	req.SetTimeout(5 * time.Second, 5 * time.Second)
	req.Body(bi)
	err = req.ToJSON(&endpointRes)
	beego.Info(req.String())
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":200, "msg":"fial", "data":err.Error()}
		c.ServeJSON()
		c.StopRun()
	} else {
		ret_data := map[string][]EndpointValues{}
		var inData, outData []EndpointValues
		for _, item := range endpointRes {
			switch item.Counter {
			case "intraffic":
				if len(item.Values) <= 12 {
					inData = item.Values
				} else {
					for i := 0; i < len(item.Values); i += 5 {
						inData = append(inData, item.Values[i])
					}
				}
			case "outtraffic":
				outData = item.Values
			}
		}
		newOutData := []EndpointValues{}
		for _, in := range inData {
			for _, out := range outData {
				if in.Timestamp == out.Timestamp {
					newOutData = append(newOutData, out)
				}
			}
		}

		_, ok := ret_data["intraffic"]
		if !ok {
			ret_data["intraffic"] = inData
		}
		_, ok1 := ret_data["outtraffic"]
		if !ok1 {
			ret_data["outtraffic"] = newOutData

		}

		c.Data["json"] = map[string]interface{}{"code":200, "msg":"succ", "data":ret_data}
		c.ServeJSON()
		c.StopRun()
	}
}