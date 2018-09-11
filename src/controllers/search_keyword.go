package controllers

import (
	"encoding/json"
	"errors"
	"models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// SearchKeywordController operations for SearchKeyword
type SearchKeywordController struct {
	BaseRouter
}

// URLMapping ...
func (c *SearchKeywordController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SearchKeyword
// @Param	body		body 	models.SearchKeyword	true		"body for SearchKeyword content"
// @Success 201 {int} models.SearchKeyword
// @Failure 403 body is empty
// @router / [post]
func (c *SearchKeywordController) Post() {
	var v models.SearchKeyword
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSearchKeyword(&v); err == nil {
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
// @Description get SearchKeyword by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SearchKeyword
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SearchKeywordController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSearchKeywordById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SearchKeyword
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SearchKeyword
// @Failure 403
// @router / [get]
func (c *SearchKeywordController) GetAll() {
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

	_, err, l := models.GetAllSearchKeyword(query, fields, sortby, offset, limit)
	beego.Info("script====>",l)
	if err != nil {
		c.SetJson(400,err.Error(),"fail")
	} else {
		req := make(map[string]interface{})
		req["list"] = l
		c.SetJson(200,req,"succ")
	}
}

// Put ...
// @Title Put
// @Description update the SearchKeyword
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SearchKeyword	true		"body for SearchKeyword content"
// @Success 200 {object} models.SearchKeyword
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SearchKeywordController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SearchKeyword{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSearchKeywordById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SearchKeyword
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SearchKeywordController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSearchKeyword(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
