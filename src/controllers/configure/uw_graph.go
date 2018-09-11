package configure

import (
	"encoding/json"
	"errors"
	"models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"controllers"
	"time"
)

// UwGraphController operations for UwGraph
type UwGraphController struct {
	controllers.BaseRouter
}

// URLMapping ...
func (c *UwGraphController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create UwGraph
// @Param	body		body 	models.UwGraph	true		"body for UwGraph content"
// @Success 201 {int} models.UwGraph
// @Failure 403 body is empty
// @router / [post]
func (c *UwGraphController) Post() {
	var v models.UwGraph
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.InsertTime = time.Now()
		v.UpdateTime = time.Now()
		if _, err := models.AddUwGraph(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.SetJson(200,"OK","succ")
		} else {
			c.SetJson(200,err.Error(),"fail")
		}
	} else {
		c.SetJson(200,err.Error(),"fail")
	}
}

// GetOne ...
// @Title Get One
// @Description get UwGraph by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UwGraph
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UwGraphController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUwGraphById(id)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		c.SetJson(200,v,"succ")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get UwGraph
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UwGraph
// @Failure 403
// @router / [get]
func (c *UwGraphController) GetAll() {
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
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("page"); err == nil {
		if v < 1 {
			v = 1
		}
		start := (v - 1) * limit
		offset = start
	}
	// query: {k:v,k:v}
	if v := c.GetString("query"); v != "" {
		if err := json.Unmarshal([]byte(v), &query); err != nil {
			c.SetJson(200,errors.New("Error: invalid query key/value pair"),"fail")
		}
	}

	beego.Info("===>", query)
	_, err, l := models.ListAllGraph(query, fields, sortby, offset, limit)
	num := models.CountGraph(query)
	beego.Info("GetAllUwAutofill====>",l)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		req := make(map[string]interface{})
		req["list"] = l
		req["total"] = num
		c.SetJson(200,req,"succ")
	}
}

// Put ...
// @Title Put
// @Description update the UwGraph
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UwGraph	true		"body for UwGraph content"
// @Success 200 {object} models.UwGraph
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UwGraphController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.UwGraph{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUwGraphById(&v); err == nil {
			c.SetJson(200,"OK","succ")
		} else {
			c.SetJson(201,err.Error(),"fail")
		}
	} else {
		c.SetJson(201,err.Error(),"fail")
	}
}

// Delete ...
// @Title Delete
// @Description delete the UwGraph
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UwGraphController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUwGraph(id); err == nil {
		c.SetJson(200,"OK","succ")
	} else {
		c.SetJson(201,err.Error(),"fail")
	}
}
