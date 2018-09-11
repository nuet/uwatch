//卷皮新后台相关接口
package new

import (
	"errors"
	"library/juanpiapi"
)

var api = juanpiapi.Api{
	Url:    "http://b9603b9ece8b.juanpi.org",
	Appkey: "oa",
	Secret: "!@oaapis&*#",
}

//后台岗位
type Station struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Depath int    `json:"depath"`
	Pid    int    `json:"pid"`
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

//获取后台岗位
func GetStations() (stations []Station, err error) {
	result := &struct {
		respData
		Data []Station
	}{}

	err = api.Post("/OaApi/getStations", map[string]string{}, result)
	parseRespIsOk(result.respData, &err)
	if err == nil {
		stations = result.Data
	}
	return
}

//更新用户信息
func SetUserRole(data map[string]string) error {
	result := &respData{}
	if data["nice_name"] == "" || data["user_name"] == "" || data["depart_id"] == "" || data["depart_name"] == "" || data["station_id"] == "" {
		result.Msg = "必传参数为空!"
		return result
	}
	err := api.Post("/OaApi/setUserRole", data, result)
	parseRespIsOk(*result, &err)
	if err != nil && err.Error() == "EOF" {
		err = errors.New("更新用户失败")
	}
	return err
}

// SetUsersRoles 批量更新用户信息
func SetUsersRoles(data map[string]string) error {
	result := &respData{}
	err := api.Post("/OaApi/setUsersRoles", data, result)
	parseRespIsOk(*result, &err)
	if err != nil && err.Error() == "EOF" {
		err = errors.New("更新用户失败")
	}
	return err
}

// UpdateDepartmentName OA部门更名同步新后台角色信息
func UpdateDepartmentName(data map[string]string) error {
	result := &respData{}
	if data["depart_id"] == "" || data["depart_name"] == "" {
		result.Msg = "比传参数为空"
		return result
	}
	err := api.Post("/OaApi/updateDepartmentName", data, result)
	parseRespIsOk(*result, &err)
	if err != nil && err.Error() == "EOF" {
		err = errors.New("更新用户失败")
	}
	return err
}

//新增后台用户
func AddUserMes(data map[string]string) error {
	result := &respData{}
	var paramsKey = []string{"username","password","roleid","nice_name","user_name","depart_id","depart_name","station_id"};
	for _,v := range paramsKey{
		if (data[v] == "") {
			result.Msg = "必传参数为空!"
			return result
		}
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
