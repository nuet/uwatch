package dashboard

import (
	"encoding/json"
	"errors"
	"models"
	"strconv"
	"strings"

	//"github.com/astaxie/beego"
	"controllers"
	"time"
	"github.com/astaxie/beego"
	"library/common"
)

// UwNavController operations for UwNav
type UwNavController struct {
	controllers.BaseRouter
}

// URLMapping ...
func (c *UwNavController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create UwNav
// @Param	body		body 	models.UwNav	true		"body for UwNav content"
// @Success 201 {int} models.UwNav
// @Failure 403 body is empty
// @router / [post]
func (c *UwNavController) Post() {
	var v models.UwNav
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Createtime = time.Now()
		v.Updatetime = time.Now()
		if _, err := models.AddUwNav(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.SetJson(200,"OK","succ")
		} else {
			c.SetJson(200,err.Error(),"fail")
		}
	} else {
		c.SetJson(200,err.Error(),"fail")
	}
}

func (c *UwNavController) GetNavAll() {
	var fields []string
	var sortby []string
	var query = make(map[string]string)
	var limit int64 = 100
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

	_, err, l := models.GetAllUwNav(query, fields, sortby, offset, limit)
	_, err, list := models.GetAllUwNavList(query, fields, sortby, offset, limit)
	beego.Info("GetAllUwNav====>",l)
	beego.Info("GetAllUwNavList===>", list)
	var lData [][]map[string]interface{}
	for _,v := range l {
		var listData []map[string]interface{}
		for _,k := range list {
			if v["Id"] == k["NavId"] {
				temp := map[string]interface{}{}
				temp["name"] = common.GetString(k["Name"])
				temp["url"] = common.GetString(k["Url"])
				temp["pname"] = common.GetString(v["Name"])
				temp["purl"] = common.GetString(v["Url"])
				listData = append(listData, temp)
			}
		}
		lData = append(lData, listData)
	}
	beego.Info("listData==>", lData)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		req := make(map[string]interface{})
		req["list"] = lData
		c.SetJson(200,req,"succ")
	}
}

// GetOne ...
// @Title Get One
// @Description get UwNav by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UwNav
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UwNavController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUwNavById(id)
	if err != nil {
		c.SetJson(200,err.Error(),"fail")
	} else {
		c.SetJson(200,v,"succ")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get UwNav
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UwNav
// @Failure 403
// @router / [get]
func (c *UwNavController) GetAll() {
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

	_, err, l := models.GetAllUwNav(query, fields, sortby, offset, limit)
	num := models.CountNav(query)
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
// @Description update the UwNav
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UwNav	true		"body for UwNav content"
// @Success 200 {object} models.UwNav
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UwNavController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.UwNav{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Updatetime = time.Now()
		if err := models.UpdateUwNavById(&v); err == nil {
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
// @Description delete the UwNav
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UwNavController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUwNav(id); err == nil {
		c.SetJson(200,"OK","fail")
	} else {
		c.SetJson(201,err.Error(),"fail")
	}
}
