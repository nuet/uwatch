package dashboard

import (
	"encoding/json"
	"errors"
	"models"
	"strconv"
	"strings"

	//"github.com/astaxie/beego"
	"controllers"
	"github.com/astaxie/beego"
	"time"
)

// UwNavListController operations for UwNavList
type UwNavListController struct {
	controllers.BaseRouter
}

// URLMapping ...
func (c *UwNavListController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create UwNavList
// @Param	body		body 	models.UwNavList	true		"body for UwNavList content"
// @Success 201 {int} models.UwNavList
// @Failure 403 body is empty
// @router / [post]
func (c *UwNavListController) Post() {
	var v models.UwNavList
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Createtime = time.Now()
		v.Updatetime = time.Now()
		if _, err := models.AddUwNavList(&v); err == nil {
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
// @Description get UwNavList by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UwNavList
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UwNavListController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUwNavListById(id)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		c.SetJson(200,v,"succ")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get UwNavList
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UwNavList
// @Failure 403
// @router / [get]
func (c *UwNavListController) GetAll() {
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

	// offset: 0 (default is 0)
	if v, err := c.GetInt64("page"); err == nil {
		if v < 1 {
			v = 1
		}
		start := (v - 1) * limit
		offset = start
	}

	_, err, l := models.GetAllUwNavList(query, fields, sortby, offset, limit)
	num := models.CountNavList(query)
	beego.Info("GetAllUwNav====>",l)
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
// @Description update the UwNavList
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UwNavList	true		"body for UwNavList content"
// @Success 200 {object} models.UwNavList
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UwNavListController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.UwNavList{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Updatetime = time.Now()
		if err := models.UpdateUwNavListById(&v); err == nil {
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
// @Description delete the UwNavList
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UwNavListController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUwNavList(id); err == nil {
		c.SetJson(200,"OK","fail")
	} else {
		c.SetJson(201,err.Error(),"fail")
	}
}
