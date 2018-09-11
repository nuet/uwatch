package controllers

import (
	"encoding/json"
	"errors"
	"models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"crypto/tls"
	"library/common"
	"library/oa"
	"github.com/astaxie/beego/orm"
)

// SystemUserController operations for SystemUser
type SystemUserController struct {
	BaseRouter
}

type UserListResp struct {
	Status int      `json:"status"`
	Msg    string   `json:"Msg"`
	Data   []UserList `json:"Data"`
}

type UserList struct {
	Email    string           `json:"Email"`
	Id       string           `json:"Id"`
	Username string           `json:"Username"`
	Ischeck  string           `json:"ischeck"`
}

type DepListResp struct {
	Status int      `json:"status"`
	Msg    string   `json:"Msg"`
	Data   map[string]DepList `json:"Data"`
}

type DepList struct {
	Id       string            `json:"Id"`
	Persno   string		   `json:"Persno"`
	Name     string		   `json:"Name"`
	Father   string		   `json:"Father"`
	Status   string		   `json:"Status"`
	FullName string		   `json:"FullName"`
}

// URLMapping ...
func (c *SystemUserController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetList", c.GetList)
	c.Mapping("GetUserList", c.GetUserList)
	c.Mapping("GetUserRole", c.GetUserRole)
	c.Mapping("GetDepartmentList", c.GetDepartmentList)
	c.Mapping("GetTeamList", c.GetTeamList)
}

// Post ...
// @Title Post
// @Description create SystemUser
// @Param	body		body 	models.SystemUser	true		"body for SystemUser content"
// @Success 201 {int} models.SystemUser
// @Failure 403 body is empty
// @router / [post]
func (c *SystemUserController) Post() {
	var v models.SystemUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.SuName == "" || common.Empty(v.SuName) {
			c.SetJson(200, "用户名不能为空", "fail")
		} else {
			l, err := models.GetSystemUserByName(v.SuName)
			if err != nil {
				c.SetJson(200, err.Error(), "fail")
			} else if len(l) > 0 {
				c.SetJson(200, "用户已存在", "fail")
			} else {
				if v.SuIsAdmin == "" {
					v.SuIsAdmin = "否"
				}
				if v.SuStatus == "" {
					v.SuStatus = "enable"
				}
				if _, err := models.AddSystemUser(&v); err == nil {
					c.Ctx.Output.SetStatus(201)
					c.SetJson(200, "OK", "succ")
				} else {
					c.SetJson(200, err.Error(), "fail")
				}
			}
		}
	} else {
		c.SetJson(200, err.Error(), "fail")
	}
}

// GetOne ...
// @Title Get One
// @Description get SystemUser by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SystemUser
// @Failure 403 :id is empty
// @router /:id [get]
//func (c *SystemUserController) GetOne() {
//	idStr := c.Ctx.Input.Param(":id")
//	id, _ := strconv.Atoi(idStr)
//	v, err := models.GetSystemUserById(id)
//	if err != nil {
//		c.Data["json"] = err.Error()
//	} else {
//		c.Data["json"] = v
//	}
//	c.ServeJSON()
//}

// GetAll ...
// @Title Get All
// @Description get SystemUser
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SystemUser
// @Failure 403
// @router / [get]
func (c *SystemUserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
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
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSystemUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.SetJson(200, err.Error(), "fail")

	} else {
		c.SetJson(200, l, "succ")
	}
}

// Put ...
// @Title Put
// @Description update the SystemUser
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SystemUser	true		"body for SystemUser content"
// @Success 200 {object} models.SystemUser
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SystemUserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SystemUser{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSystemUserById(&v); err == nil {
			c.SetJson(200, "OK", "succ")
		} else {
			c.SetJson(200, err.Error(), "fail")
		}
	} else {
		c.SetJson(200, err.Error(), "fail")
	}
}

// Delete ...
// @Title Delete
// @Description delete the SystemUser
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SystemUserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSystemUser(id); err == nil {
		c.SetJson(200, "OK", "succ")
	} else {
		c.SetJson(200, err.Error(), "fail")
	}
}

// GetUserList ...
// @Title GetUserList
// @Description Get Juanpi user list from UIC
// @Param
// @Success 200 {string} UserListResp
// @Failure 403
// @router /getUserList [get]
func (c *SystemUserController) GetUserList() {

	var ret interface{}
	if !GetCache.IsExist("userList") || common.GetString(GetCache.Get("userList")) == "" {
		sql, _ := orm.NewQueryBuilder("mysql")
		sql_str := sql.Select("*").From("user").String()
		o := orm.NewOrm()
		o.Using("uic")
		var userlist []orm.Params
		o.Raw(sql_str).Values(&userlist)
		GetCache.Put("userList", userlist, 1800*time.Second)
		ret = userlist
	} else {
		ret = GetCache.Get("userList")
	}

	b, _ := json.Marshal(ret)
	c.SetJson(200, string(b), "succ")
}

// GetTeamList ...
// @Title GetTeamList
// @Description Get Uic team list from UIC
// @Param
// @Success 200 {string} TeamListResp
// @Failure 403
// @router /getTeamList [get]
func (c *SystemUserController) GetTeamList() {

	var ret interface{}
	sql, _ := orm.NewQueryBuilder("mysql")
	sql_str := sql.Select("*").From("team").String()
	o := orm.NewOrm()
	o.Using("uic")
	var teamlist []orm.Params
	o.Raw(sql_str).Values(&teamlist)
	ret = teamlist

	b, _ := json.Marshal(ret)
	c.SetJson(200, string(b), "succ")
}

// GetDepartmentList ...
// @Title GetDepartmentList
// @Description Get Juanpi Department list from OA
// @Param
// @Success 200 {string} DepartmentListResp
// @Failure 403
// @router /getDepartmentList [get]
func (c *SystemUserController) GetDepartmentList() {

	params := map[string]string{}
	params["appKey"] = "OA"
	serect := "OA@#$"
	sign := oa.MakeSign(params, serect)

	oaAuthIp := beego.AppConfig.String("oa_auth_web")
	getDepartmentListUrl := oaAuthIp + DEPARTMENT_LIST_URL + "?appKey=OA&sign=" + sign

	DepListResp := DepListResp{}

	req := httplib.Get(getDepartmentListUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(8 * time.Second, 8 * time.Second)
	beego.Info("getDepartmentListUrl===>",getDepartmentListUrl)
	//reqStr, _ := req.String()
	beego.Info("reqreqreqreqreq===>",req)
	err := req.ToJSON(&DepListResp)
	beego.Info("DepListResp_data===>",DepListResp.Data)
	if err != nil || DepListResp.Status != 200 {
		beego.Error("GET userList err: ", err)
		c.SetJson(200, "ERROR", "fail")

	}

	b, _ := json.Marshal(DepListResp.Data)
	c.SetJson(200, string(b), "succ")
}

// GetList ...
// @Title Get All
// @Description get SystemUser
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SystemUser
// @Failure 403
// @router /getList [get]
func (c *SystemUserController) GetList() {
	var fields []string
	var sortby []string
	var order []string
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
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.SetJson(200, errors.New("Error: invalid query key/value pair"), "fail")
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllSystemUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.SetJson(200, err.Error(), "fail")
	} else {
		for i, item := range l {
			roleNames := ""
			if (!common.Empty(item["SuRoles"])) {
				roles := strings.Split(item["SuRoles"].(string), ",")
				for _, idstr := range roles {
					id, _ := strconv.Atoi(idstr)
					role, _ := models.GetSystemRoleById(id)
					roleNames = roleNames + "," + role.SrName
				}
			}
			l[i]["SuRoleName"] = common.SubString(roleNames,1,len(roleNames))
		}
		c.Data["json"] = l
	}
	c.SetJson(200, l, "succ")
}

// GetUserRole ...
// @Title GetUserRole
// @Description Get Login User Role Perms
// @Param
// @Success 200
// @Failure 403
// @router /getUserRole [get]
func (c *SystemUserController) GetUserRole() {

	user, err := c.GetUserName();
	if user == "" || err != nil {
		c.SetJson(50008, "", "非法请求")
	} else {
		type RoleData struct {
			Role      []string
			RoleName  []string
		}
		var data RoleData
		user, err := models.GetSystemUserByName(user)
		if len(user) == 0 || err != nil {
			data.Role = append(data.Role, "visitor")
			data.RoleName = append(data.RoleName, "visitor")
		} else {
			if common.GetString(user[0].SuIsAdmin) == "是" {
				data.Role = append(data.Role, "admin")
				data.RoleName = append(data.RoleName, "admin")
			} else {
				if common.Empty(common.GetString(user[0].SuRoles)) {
					data.Role = append(data.Role, "visitor")
					data.RoleName = append(data.RoleName, "visitor")
				} else {
					data.Role,data.RoleName = getPerms(common.GetString(user[0].SuRoles))
				}
			}
		}

		c.SetJson(200, data, "succ")
	}

}

func getPerms(roleIds string) (per []string,names []string) {
	list := strings.Split(roleIds, ",")
	for _, idstr := range list {
		id, _ := strconv.Atoi(idstr)
		role, _ := models.GetSystemRoleById(id)
		names = append(names,role.SrName)
		perms := []string{}
		err := json.Unmarshal([]byte(role.SrPerms), &perms)
		beego.Info(err)
		for _, perm := range perms {
			per = append(per, perm)
		}
	}

	return
}