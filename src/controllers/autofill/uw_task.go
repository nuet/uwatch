package autofill

import (
	"controllers"
	"encoding/json"
	"errors"
	"strings"
	"models"
	"strconv"
	"github.com/astaxie/beego/orm"
	"library/common"
	"github.com/astaxie/beego"
)

type UwTaskController struct {
	controllers.BaseRouter
}

// URLMapping ...
func (c *UwTaskController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetRecords", c.GetRecords)
}

// GetAll ...
// @Title Get All
// @Description get UwTask
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UwTask
// @Failure 403
// @router / [get]
func (c *UwTaskController) GetAll() {
	var fields []string
	var sortby []string
	var query = make(map[string]string)
	var limit int64 = 20
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
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
	// sortby: col1,-col2
	if v := c.GetString("sort"); v != "" {
		sortby = strings.Split(v, ",")
	}

	// query: {k:v,k:v}
	if v := c.GetString("query"); v != "" {
		if err := json.Unmarshal([]byte(v), &query); err != nil {
			c.SetJson(200,errors.New("Error: invalid query key/value pair"),"fail")
		}
	}

	_, err, l := models.ListAllUwTask(query, fields, sortby, offset, limit)
	num := models.CountUwTask(query)
	if err != nil {
		beego.Error(err)
		c.SetJson(200,err.Error(),"fail")
	} else {
		for i, item := range l {
			v, err := models.GetUwAutofillById(common.GetInt(item["AutofillId"]))
			if err != nil {
				l[i]["AutoFill"] = nil
			} else {
				l[i]["AutoFill"] = v
			}
		}

		req := make(map[string]interface{})
		req["list"] = l
		req["total"] = num
		c.SetJson(200,req,"succ")
	}
}


// GetOne ...
// @Title Get One
// @Description get ConsoleScript by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UwTask
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UwTaskController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUwTaskById(id)

	var records []orm.Params
	if id <= 0 {
		timeNow := c.GetString("time")
		o := orm.NewOrm()
		o.Raw("SELECT * FROM `uw_record` where task_id =?  and created_at > ? ORDER BY `id` ASC ", idStr, timeNow).Values(&records)
	} else {
		o := orm.NewOrm()
		o.Raw("SELECT * FROM `uw_record` where task_id = ? ORDER BY `id` ASC ", idStr).Values(&records)
	}

	req := make(map[string]interface{})
	req["task"] = v
	req["record"] = records
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		c.SetJson(200,req,"succ")
	}
	c.ServeJSON()
}

// GetRecords ...
// @Title GetRecords
// @Description get Records
// @Param	taskId	query	string	true	"The Key of uw_task"
// @Param	time	query	string	false	"timestamp"
// @Success 200 {object} models.UwRecord
// @Failure 403
// @router /getRecords [get]
func (c *UwTaskController) GetRecords() {
	taskId := c.GetString("taskId")
	var records []orm.Params
	if common.GetInt(taskId) <= 0 {
		timeNow := c.GetString("time")
		o := orm.NewOrm()
		o.Raw("SELECT a.*, b.title FROM `uw_record` a LEFT JOIN `uw_operation` b ON a.opt_id = b.`id` WHERE a.task_id =?  and a.created_at > ? ORDER BY a.`id` ASC ", taskId, timeNow).Values(&records)
	} else {
		o := orm.NewOrm()
		o.Raw("SELECT a.*, b.title FROM `uw_record` a LEFT JOIN `uw_operation` b ON a.opt_id = b.`id` where a.task_id = ? ORDER BY a.`id` ASC ", taskId).Values(&records)
	}
	c.SetJson(200, records, "succ")
	return
}
