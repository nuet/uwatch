//卷皮后台相关接口
package admin

import (
	"library/juanpiapi"
)

var api = juanpiapi.Api{
	Url:    "http://c20c39d15cad.juanpi.org",
	Appkey: "oa",
	Secret: "!@oaapis&*#",
}

type Role struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
	GroupId  string `json:"group_id"`
	RoleType string `json:"role_type"`
	Ordid    string `json:"ordid"`
	Pid      int    `json:"pid"`
}

type respData juanpiapi.RespData

func (e *respData) Error() string {
	return e.Msg
}

//解析返回的结果的状态，200表示正确，否则有错
func parseRespIsOk(e respData, err *error) {
	if *err == nil {
		if e.Code != 200 {
			*err = &e
		}
	}
}

//根据角色ID获取角色信息
func GetRoleByRoleid(roleid string) (role Role) {
	if roleid != "" {
		roles, err := GetRoleList()
		if err != nil {
			return
		}
		for _, v := range roles {
			if roleid == v.Id {
				role = v
			}
		}
	}
	return
}

// 获取后台角色列表
func GetRoleList() (roles []Role, err error) {
	result := &struct {
		respData
		Data []Role
	}{}
	err = api.Post("/OaApi/getRoles", map[string]string{}, result)

	parseRespIsOk(result.respData, &err)
	if err == nil {
		roles = result.Data
	}
	return
}

//新增后台用户
func AddUserMes(data map[string]string) error {
	result := &respData{}
	if data["username"] == "" || data["password"] == "" || data["roleid"] == "" {
		result.Msg = "必传参数为空!"
		return result
	}
	err := api.Post("/OaApi/addUserMes", data, result)
	parseRespIsOk(*result, &err)
	return err
}

//更新后台用户
func UpdateUserMes(data map[string]string) error {
	result := &respData{}
	if data["username"] == "" {
		result.Msg = "必传参数为空!"
		return result
	}
	err := api.Post("/OaApi/updateUserMes", data, result)
	parseRespIsOk(*result, &err)
	return err
}
