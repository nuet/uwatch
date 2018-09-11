package autofill

import (
	"controllers"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"models"
	"time"
	"github.com/astaxie/beego"
	"library/components"
)

type Oper struct {
	Operation    []models.UwOperation
	OperationDel []int
}

// UwAutofillController operations for UwAutofill
type UwAutofillController struct {
	controllers.BaseRouter
}

// URLMapping ...
func (c *UwAutofillController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Detection", c.Detection)
}

// Post ...
// @Title Post
// @Description create UwAutofill
// @Param	body		body 	models.UwAutofill	true		"body for UwAutofill content"
// @Success 201 {int} models.UwAutofill
// @Failure 403 body is empty
// @router / [post]
func (c *UwAutofillController) Post() {
	var v models.UwAutofill
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.InsertTime = time.Now()
		v.UpdateTime = time.Now()
		if fid, err := models.AddUwAutofill(&v); err == nil {
			var l Oper
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &l); err == nil {
				for i,oper := range l.Operation {
					oper.Step = int(i) + 1
					oper.Fid = int(fid)
					oper.UpdateTime = time.Now()
					if _, err := models.AddUwOperation(&oper); err != nil {
						c.SetJson(200,err.Error(),"fail")
					}
				}
				c.SetJson(200,"OK","succ")
			} else {
				c.SetJson(200,err.Error(),"fail")
			}
		} else {
			c.SetJson(200,err.Error(),"fail")
		}
	} else {
		c.SetJson(200,err.Error(),"fail")
	}
}

// GetOne ...
// @Title Get One
// @Description get UwAutofill by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UwAutofill
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UwAutofillController) GetOne() {
	type Rsp struct {
		*models.UwAutofill
		Operation []interface{}
	}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUwAutofillById(id)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		query := make(map[string]string)
		query["Fid"] = strconv.Itoa(v.Id)
		ml, err := models.GetAllUwOperation(query, []string{}, []string{"Step"}, []string{"asc"}, 0, 99)
		if err != nil {
			c.SetJson(200,err.Error(),"fail")
		}
		ret := &Rsp{v,ml}
		c.SetJson(200, ret, "succ")
	}
}

// GetAll ...
// @Title Get All
// @Description get UwAutofill
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UwAutofill
// @Failure 403
// @router / [get]
func (c *UwAutofillController) GetAll() {
	var fields []string
	var sortby []string
	var query = make(map[string]string)
	var limit int64 = 10
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
			c.SetJson(200,errors.New("Error: invalid query key/value pair"),"fail")
		}
	}

	_, err, l := models.ListAllAutoFill(query, fields, sortby, offset, limit)
	num := models.CountAutofill(query)
	beego.Info("GetAllUwAutofill====>",l)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		for i, _ := range l {
			l[i]["InsertTime"] = l[i]["InsertTime"].(time.Time).Format("2006-01-02 15:04:05")
			l[i]["UpdateTime"] = l[i]["UpdateTime"].(time.Time).Format("2006-01-02 15:04:05")
		}

		req := make(map[string]interface{})
		req["list"] = l
		req["total"] = num
		c.SetJson(200,req,"succ")
	}
}

// Put ...
// @Title Put
// @Description update the UwAutofill
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UwAutofill	true		"body for UwAutofill content"
// @Success 200 {object} models.UwAutofill
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UwAutofillController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.UwAutofill{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.UpdateTime = time.Now()
		if err := models.UpdateUwAutofillById(&v); err == nil {
			var l Oper
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &l); err == nil {
				for i,oper := range l.Operation {
					oper.Step = int(i) + 1
					oper.UpdateTime = time.Now()
					if oper.Id > 0 {
						if err := models.UpdateUwOperationById(&oper); err != nil {
							c.SetJson(200,err.Error(),"fail")
						}
					} else {
						oper.Fid = v.Id
						if _, err := models.AddUwOperation(&oper); err != nil {
							c.SetJson(200,err.Error(),"fail")
						}
					}
				}

				for _,oid := range l.OperationDel {
					models.DeleteUwOperation(oid)
				}

				c.SetJson(200,"OK","succ")

			} else {
				c.SetJson(200,err.Error(),"fail")
			}
		} else {
			c.SetJson(200,err.Error(),"fail")
		}
	} else {
		c.SetJson(200,err.Error(),"fail")
	}
}

// Delete ...
// @Title Delete
// @Description delete the UwAutofill
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UwAutofillController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUwAutofill(id); err == nil {
		if err1 := models.DeleteUwOperationByFid(id); err1 == nil {
			c.SetJson(200,"OK","succ")
		} else {
			c.SetJson(200,err1.Error(),"fail")
		}
	} else {
		c.SetJson(200,err.Error(),"fail")
	}
}

// Detection ...
// @Title Detection
// @Description Detection autoFill
// @Param	id		path 	int	true		"The id you want to detection"
// @Success 200 :OK
// @Failure 403 :id is empty
// @router /detection/:id [get]
func (c *UwAutofillController) Detection() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	UwAutofill, _ := models.GetUwAutofillById(id)
	s := components.BaseComponents{}
	s.SetAutofill(UwAutofill)
	s.SetTask(&models.UwTask{Id: -1})
	s.SetOperator(c.User.Rtx)
	s.SetDuration(0)
	s.SetFinal(0)

	ids, err := models.GetOperationsByFid(id)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	}

	for i, oid := range ids {
		oper, _ := models.GetUwOperationById(oid)
		s.SetOperation(oper)

		if (i == (len(ids) -1)) {
			s.SetFinal(1)
		}

		// 权限与免密码登录检测
		err := s.TestSsh()
		if err != nil {
			c.SetJson(200, "ssh目标机器错误" + err.Error(), "fial")
		}
	}

	c.SetJson(200,"OK","succ")
}