package falcon

import (
	"library/falcon/api"
)

func GetAllHostGroups() (map[string]api.HostGroupResponse, error){
	ret := make(map[string]api.HostGroupResponse)
	falconApi := api.New()
	hostGroups, err := falconApi.SearchHostGroup()
	if err != nil {
		return ret, err
	}

	for _, hostGroup := range hostGroups {
		_, ok := ret[hostGroup.HGName]
		if !ok {
			ret[hostGroup.HGName] = hostGroup
		}
	}

	return ret, nil
}

func CreateFalconHostGroup(name string) (int, error) {
	falconApi := api.New()
	param := map[string]interface{}{
		"name" : name,
	}
	hostGroup, err := falconApi.CreateHostGroup(param)
	if err != nil {
		return 0, err
	}
	return hostGroup.HGID, nil
}

func AddHostsToFalconHostGroup(hgId int, hosts []string) (string, error) {
	falconApi := api.New()
	param := map[string]interface{}{
		"hostgroup_id" : hgId,
		"hosts" : hosts,
	}
	rsp, err := falconApi.AddHostsToHostGroup(param)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func DeleteFalconHostGroup(hgId int) (string, error) {
	falconApi := api.New()
	rsp, err := falconApi.DeleteHostGroup(hgId)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func GetSearchEndpoint(endpoint string) ([]api.EndpointList, error){
	falconApi := api.New()
	EndpointList, err := falconApi.SearchEndpointList(endpoint)
	if err != nil {
		return EndpointList, err
	}

	return EndpointList, nil
}

func GetAllEndpointCounter(counter string) ([]api.CounterList, error){
	falconApi := api.New()
	CounterList, err := falconApi.SearchCounterList(counter)
	if err != nil {
		return CounterList, err
	}
	return CounterList, nil
}

func GetGraphHistory(start int, hostname []string, end int, counters []string) ([]api.GraphHistory, error){
	param := map[string]interface{}{
			"step" : 60,
			"start_time" : start,
			"hostnames" : hostname,
			"end_time" : end,
			"counters" : counters,
			"consol_fun" : "AVERAGE",
	}
	falconApi := api.New()
	GraphHistory, err := falconApi.SearchHistoryList(param)
	if err != nil {
		return GraphHistory, err
	}

	return GraphHistory, nil
}

func GetUserList() ([]api.GetUserByName, error){
	falconApi := api.New()
	res, err := falconApi.GetUserList()
	if err != nil {
		return res, err
	}

	return res, nil
}

func GetAllUserList() (map[string]api.GetUserByName, error){
	ret := make(map[string]api.GetUserByName)
	UserList := api.New()
	users, err := UserList.GetUserList()
	if err != nil {
		return ret, err
	}

	for _, user := range users {
		_, ok := ret[user.Name]
		if !ok {
			ret[user.Name] = user
		}
	}

	return ret, nil
}

func GetUserName(nikiname string) (api.GetUserByName, error){
	falconApi := api.New()
	res, err := falconApi.SearchFalconUserName(nikiname)
	if err != nil {
		return res, err
	}

	return res, nil
}

func InsertUser(name string, password string, cnname string, email string, im string, phone string, qq string) (api.GetUserByName, error){
	param := map[string]interface{}{
		"name" : name,
		"password" : password,
		"cnname" : cnname,
		"email" : email,
		"im" : im,
		"phone" : phone,
		"qq" : qq,
	}
	falconApi := api.New()
	PostUser, err := falconApi.CreateFalconUser(param)
	if err != nil {
		return PostUser, err
	}

	return PostUser, nil
}

func UpdateUser(id int64, name string, cnname string, email string, im string, phone string, qq string, login string) (api.GetUserByName, error){
	param := map[string]interface{}{
		"id" : id,
		"name" : name,
		"cnname" : cnname,
		"email" : email,
		"im" : im,
		"phone" : phone,
		"qq" : qq,
	}
	falconApi := api.New()
	PostUser, err := falconApi.UpdateFalconUser(param, login)
	if err != nil {
		return PostUser, err
	}

	return PostUser, nil
}

func DeleteUser(user_id int64) (string, error) {
	falconApi := api.New()
	param := map[string]interface{}{
		"user_id" : user_id,
	}
	rsp, err := falconApi.DeleteUserId(param)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func GetTeamName(name string) (api.GetTeamByName, error){
	falconApi := api.New()
	res, err := falconApi.SearchFalconTeamName(name)
	if err != nil {
		return res, err
	}

	return res, nil
}

func InsertTeam(team_name string, resume string, users []int64) (string, error){
	falconApi := api.New()
	param := map[string]interface{}{
		"team_name" : team_name,
		"resume" : resume,
		"users" : users,
	}
	PostTeam, err := falconApi.CreateFalconTeam(param)
	if err != nil {
		return "", err
	}

	return PostTeam.SuccMsg, nil
}

func UpdateTeam(team_id int, name string, resume string, users []int64) (string, error){
	falconApi := api.New()
	param := map[string]interface{}{
		"team_id" : team_id,
		"name" : name,
		"resume" : resume,
		"users" : users,
	}
	PutTeam, err := falconApi.UpdateFalconTeam(param)
	if err != nil {
		return "", err
	}

	return PutTeam.SuccMsg, nil
}

func DeleteTeam(team_id int) (string, error) {
	falconApi := api.New()
	rsp, err := falconApi.DeleteTeamId(team_id)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func GetTeamList(name string) ([]api.GetTeamListData, error){
	falconApi := api.New()
	TeamList, err := falconApi.SearchTeamList(name)
	if err != nil {
		return TeamList, err
	}
	return TeamList, nil
}

func GetAllTeamList(name string) (map[string]api.GetTeamListData, error){
	ret := make(map[string]api.GetTeamListData)
	TeamList := api.New()
	userTeam, err := TeamList.SearchTeamList(name)
	if err != nil {
		return ret, err
	}
	for _, userGroup := range userTeam {
		_, ok := ret[userGroup.Team.Name]
		if !ok {
			ret[userGroup.Team.Name] = userGroup
		}
	}
	return ret, nil
}
func SetFalconAdmin() (string, error){
	falconApi := api.New()
	param := map[string]interface{}{
		"user_id" : 498,
		"admin" : "yes",
	}
	PutTeam, err := falconApi.SetFalconAdmin(param)
	if err != nil {
		return "", err
	}
	return PutTeam.SuccMsg, nil
}

func GetFalconUsersId(users []string) ([]int64, error) {
	falconApi := api.New()
	var ret []int64
	for _, user := range users {
		res, err := falconApi.SearchFalconUserName(user)
		if err != nil {
			return ret, err
		} else {
			ret = append(ret, res.Id)
		}
	}
	return ret, nil
}

func CreateTemplate(name string, parentId int) (int, error) {
	ret := 0
	falconApi := api.New()
	param := map[string]interface{}{
		"parent_id" : parentId,
		"name" : name,
	}
	_, err := falconApi.CreateTemplate(param)
	if err != nil {
		return ret, err
	}

	list, err := falconApi.GetTemplateList()

	for _, tpl := range list.Templates {
		if tpl.Template.CreateUser == "root" && tpl.Template.TplName == name {
			ret = tpl.Template.TplId
			break
		}
	}

	return ret, nil
}

func DeleteTemplate(tplId int) (string, error) {
	falconApi := api.New()
	rsp, err := falconApi.DeleteTemplate(tplId)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func GetTemplateByName(name string) (api.TplInfo, error) {
	ret := api.TplInfo{}
	falconApi := api.New()
	list, err := falconApi.GetTemplateList()
	if err != nil {
		return ret, err
	}

	tplId := 0
	for _, tpl := range list.Templates {
		if tpl.Template.CreateUser == "root" && tpl.Template.TplName == name {
			tplId = tpl.Template.TplId
			break
		}
	}

	ret, err = falconApi.GetTplInfoById(tplId)
	return ret, err
}

func CreateTplAction(tplId int, uic string) (string, error) {
	falconApi := api.New()
	param := map[string]interface{}{
		"uic" : uic,
		"tpl_id" : tplId,
	}
	rsp, err := falconApi.CreateTplActoin(param)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func CreateNodata(groupName string) (api.NoData, error){
	falconApi := api.New()
	var ret api.NoData
	param := api.NoDataReq{
		Name: groupName + "-" + "agent.alive",
		Obj: groupName,
		ObjType: "group",
		Metric: "agent.alive",
		Tags: "",
		DsType: "GAUGE",
		Step: 60,
		Mock: -1,
	}
	ret, err := falconApi.CreateNodata(param)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetNodataInfoByName(groupName string) (api.NoData, error){
	falconApi := api.New()
	ret := api.NoData{}
	noDatas, err := falconApi.GetNodataList()
	if err != nil {
		return ret, err
	}
	name := groupName + "-" + "agent.alive"
	for _, noData := range noDatas {
		if noData.Name == name {
			ret = noData
			break
		}
	}
	return ret, nil
}

func CreateStrategy(tplId int, s Strategy) (string, error) {
	falconApi := api.New()
	param := map[string]interface{}{
		"tpl_id" : tplId,
		"tags" : s.Tags,
		"right_value": s.Value,
		"priority": s.Priority,
		"op" : s.Op,
		"note" : s.Note,
		"metric" : s.Metric,
		"max_step" : s.MaxStep,
		"func" : s.Func,
		"run_begin" : "",
		"run_end" : "",
	}
	rsp, err := falconApi.CreateStrategy(param)
	if err != nil {
		return "", err
	}
	return rsp.SuccMsg, nil
}

func BindTemplateToHostGroup(tplId, hgId int) (api.BindTHResponse, error) {
	falconApi := api.New()
	ret := api.BindTHResponse{}
	param := map[string]interface{}{
		"tpl_id" : tplId,
		"grp_id" : hgId,
	}
	ret, err := falconApi.BindTemplateToHostGroup(param)
	if err != nil {
		return ret, err
	}
	return ret, nil
}