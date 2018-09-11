package tasks

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ldap.v2"
	"strconv"
	"library/falcon"
	"library/falcon/api"
	"library/jumpserver"
	js_api "library/jumpserver/api"
	"library/common"
)
var dep_name []map[string]interface{}
var dep_user []map[string]interface{}
var LDAP_TEAMS map[string]interface{}
var LDAP_USERS map[string]interface{}
type Groups struct {
	Id 		string			`json:"id"`
	Users []string 	`json:"users"`
	Is_discard bool `json:"is_discard"`
	Discard_time string `json:"discard_time"`
	Name string `json:"name"`
	Comment string `json:"comment"`
	Date_created string `json:"date_created"`
	Created_by string 	`json:"created_by"`
}
func init() {
	orm.Debug = false
	LDAP_TEAMS = make(map[string]interface{})
	LDAP_USERS = make(map[string]interface{})
}
func UserGroupSync() error {

	ladpIp := beego.AppConfig.String("ldapIp")
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ladpIp, "389"))
	if err != nil {
		beego.Error(err)
		return err
	}

	o := orm.NewOrm()
	o.Using("ldap")
	defer l.Close()

	firm_lines := []string{
		"cn=[394]综合管理中心,dc=juanpi,dc=com",
		"cn=[652]卷皮事业部,dc=juanpi,dc=com",
		"cn=[143]客服中心,dc=juanpi,dc=com",
		"cn=[737]会过事业部,dc=juanpi,dc=com",
		"cn=[400]三当家事业部,dc=juanpi,dc=com",
	}

	//set, _ := falcon.SetFalconAdmin()
	//beego.Info("err---set admin===>", firm_lines)

	//删除 jumpserver 所有用户组
	//groupsList, _ := jumpserver.GetGroupsList()
	//	for _,j := range groupsList {
	//		err := jumpserver.DeleteGroup(common.GetString(j.Id))
	//		if err != nil {
	//			beego.Info("删除Jumpserver groups数据失败===>", err)
	//		} else {
	//			beego.Info("删除Jumpserver groups数据成功===>", common.GetString(j.Name))
	//		}
	//	}

	//return nil
	for _, firm_line := range firm_lines {
		visit_dep(l, o, firm_line, "", nil)
	}

	falcon_user_list, err := falcon.GetAllUserList()
	falcon_group_list, err := falcon.GetAllTeamList("/")
	jumpserver_group_list, _ := jumpserver.GetAllGroups()
	jumpserver_user_list, _ := jumpserver.GetAllUserList()

	//同步Jumpserver、Falcon 用户和用户组信息
	sync_falcon_user(falcon_user_list)
	sync_falcon_team(falcon_group_list, falcon_user_list)
	sync_jumpserver_group(jumpserver_group_list)
	sync_jumpserver_user(jumpserver_user_list, jumpserver_group_list)
	//反查ldap 删除falcon中无效数据
	go deleteFalconUser(falcon_user_list)
	go deleteFalconTeam(falcon_group_list)

	////反查ldap 删除jumpserver中无效数据
	go deleteJumpserverGroup(jumpserver_group_list)
	go deleteJumpserverUser(jumpserver_user_list)



	return nil
}

func visit_dep(conn *ldap.Conn, o orm.Ormer, filter, dn string, ret []map[string]interface{}) {
	searchRequest := ldap.NewSearchRequest(
		filter, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=*))",                                                        // The filter to apply
		[]string{"dn", "cn", "sn", "mail", "mobile","ou", "uniqueMember", "objectClass"}, // A list attributes to retrieve
		nil,
	)

	var ret_data []map[string]interface{}
	if ret != nil {
		ret_data = ret
	}

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return
	}
	for _, entry := range sr.Entries {
		if entry.GetAttributeValue("objectClass") == "groupOfUniqueNames" {
			idStr := getDepId(entry.DN)
			name := getDepName(entry.DN)
			id, _ := strconv.Atoi(idStr)
			temp := map[string]interface{}{}
			if common.GetString(dn) == "" {
				temp["Parent"] = "1"
			} else {
				temp["Parent"] = dn
			}
			temp["Name"] = name
			temp["Id"] = id
			ret_data = append(ret_data, temp)
			t := ""
			dep_id := ""
			for j,x := range ret_data {
				if j == 0 {
					t = common.GetString(x["Name"])
					dep_id = common.GetString(x["Id"])
				} else {
					t = t + "/" + common.GetString(x["Name"])
					dep_id = dep_id + "/" + common.GetString(x["Id"])
				}
			}
			dep_temp := map[string]interface{}{}
			dep_temp["dep_name"] = t
			dep_temp["dep_id"] = dep_id
			dep_temp["dep_users"] = map[string]interface{}{}
			dep_name = append(dep_name, dep_temp)
			LDAP_TEAMS[t] = common.GetString(t) + "-" + common.GetString(dep_id)
			for _, member := range entry.GetAttributeValues("uniqueMember") {
				if member != "" {
					visit_dep(conn, o, member, idStr, ret_data)
				}
			}
		}else {
			t := ""
			dep_id := ""
			for j,x := range ret_data {
				if j == 0 {
					t = common.GetString(x["Name"])
					dep_id = common.GetString(x["Id"])
				} else {
					t = t + "/" + common.GetString(x["Name"])
					dep_id = dep_id + "/" + common.GetString(x["Id"])
				}
			}
			nikiname := strings.Split(entry.GetAttributeValue("mail"), "@")[0]
			user_temp := map[string]interface{}{}
			for _, k := range dep_name {
				if common.GetString(k["dep_name"]) == common.GetString(t) {
					user_temp["dep"] = t
					user_temp["dep_user"] = common.GetString(nikiname)
					user_temp["password"] = common.GetString(nikiname)
					user_temp["cnname"] = entry.GetAttributeValue("cn")
					user_temp["email"] = entry.GetAttributeValue("mail")
					user_temp["im"] = common.GetString(nikiname)
					user_temp["phone"] = entry.GetAttributeValue("mobile")
					user_temp["qq"] = entry.GetAttributeValue("mobile")
					user_temp["im"] = common.GetString(nikiname)
					dep_user = append(dep_user,user_temp)
					LDAP_USERS[common.GetString(nikiname)] = user_temp["dep_user"]
				} else {
					continue
				}
			}
		}
	}
}

func getDepId(dn string) string {
	re, _ := regexp.Compile("[[0-9]+]")
	idStr := re.FindString(dn)
	id := strings.Trim(strings.Trim(idStr, "["), "]")
	return id
}

func getDepName(dn string) string {
	cn := strings.Split(dn, ",")[0]
	name := strings.Split(cn, "]")[1]
	return name
}

func getUsername(dn string) string {
	cn := strings.Split(dn, ",")[0]
	name := strings.Split(cn, "=")[0]
	return name
}

func sync_falcon_user(falcon_users map[string]api.GetUserByName){
	for _, user := range dep_user {
		if _, ok := falcon_users[common.GetString(user["dep_user"])]; !ok {
			//插入
			name := common.GetString(user["dep_user"])
			password := common.GetString(user["password"])
			cnname := common.GetString(user["cnname"])
			email := common.GetString(user["email"])
			im := common.GetString(user["im"])
			phone := common.GetString(user["phone"])
			qq := common.GetString(user["qq"])
			_, err := falcon.InsertUser(name, password, cnname, email, im, phone, qq)
			if err != nil {
				beego.Info("falcon插入ldap用户数据失败===>", err)
			} else {
				beego.Info("falcon插入ldap用户数据成功===>", name)
			}
		} else {
			beego.Info("匹配falcon中是否存在ldap用户==>Yes->更新")
			id := falcon_users[common.GetString(user["dep_user"])].Id
			name := common.GetString(user["dep_user"])
			cnname := common.GetString(user["cnname"])
			email := common.GetString(user["email"])
			im := common.GetString(user["dep_user"])
			phone := common.GetString(user["phone"])
			qq := common.GetString(user["qq"])
			if email != falcon_users[common.GetString(user["dep_user"])].Email || im != falcon_users[common.GetString(user["dep_user"])].Im || phone != falcon_users[common.GetString(user["dep_user"])].Phone ||qq != falcon_users[common.GetString(user["dep_user"])].Qq {
				_, err := falcon.UpdateUser(id, name, cnname, email, im, phone, qq, name)
				if err != nil {
					beego.Info("falcon更新ldap数据失败===>", err)
				} else {
					beego.Info("falcon更新ldap数据成功===>", name)
				}
			} else {
				beego.Info("falcon与ldap数据一致无需变更")
			}
		}
	}

}

func sync_jumpserver_user(jumpserver_users map[string]js_api.Users, jumpserver_groups map[string]js_api.Groups){
	for _, user := range dep_user {
		if _, ok := jumpserver_users[common.GetString(user["dep_user"])]; !ok {
			//插入
			groups_display := common.GetString(user["dep"])
			var groups []string
			groups = append(groups, jumpserver_groups[groups_display].Id)
			name := common.GetString(user["dep_user"])
			wechat := common.GetString(user["dep_user"])
			cnname := common.GetString(user["cnname"])
			email := common.GetString(user["email"])
			comment := "LDAP"
			phone := common.GetString(user["phone"])
			_, err := jumpserver.CreateUser(cnname, name, email, phone,comment,wechat,groups_display,groups)
			if err != nil {
				beego.Info("插入Jumpserver user数据失败===>", err)
			} else {
				beego.Info("插入Jumpserver user数据成功===>", cnname)
			}
		} else {
			id := jumpserver_users[common.GetString(user["dep_user"])].Id
			groups_display := common.GetString(user["dep"])
			var groups []string
			groups = append(groups, jumpserver_groups[groups_display].Id)
			name := common.GetString(user["dep_user"])
			wechat := common.GetString(user["dep_user"])
			cnname := common.GetString(user["cnname"])
			email := common.GetString(user["email"])
			comment := "LDAP"
			phone := common.GetString(user["phone"])
			if groups_display != jumpserver_users[common.GetString(user["dep_user"])].Groups_display || email != jumpserver_users[common.GetString(user["dep_user"])].Email || phone != jumpserver_users[common.GetString(user["dep_user"])].Phone || wechat != jumpserver_users[common.GetString(user["dep_user"])].Wechat || comment != jumpserver_users[common.GetString(user["dep_user"])].Comment{
				_, err := jumpserver.UpdateUser(id, cnname, name, email, phone,comment,wechat,groups_display,groups)
				if err != nil {
					beego.Info("Jumpserver更新ldap数据失败===>", err)
				} else {
					beego.Info("Jumpserver更新ldap数据成功===>", name)
				}
			} else {
				beego.Info("Jumpserver与ldap用户数据一致无需变更")
			}
		}
	}

}

func sync_falcon_team(falcon_team map[string]api.GetTeamListData, falcon_users map[string]api.GetUserByName){
	for _, dep := range dep_name {
		var users []int64
		for _,j := range dep_user {
			if common.GetString(dep["dep_name"]) == common.GetString(j["dep"]) {
				for _,i := range falcon_users {
					if common.GetString(i.Name) == common.GetString(j["dep_user"]) {
						users = append(users, i.Id)
					}
				}
			}
		}
		dep["dep_users"] = users
		if _, ok := falcon_team[common.GetString(dep["dep_name"])]; !ok {
			q := strings.Split(common.GetString(dep["dep_name"]), "/")
			if strings.Contains(common.GetString(q[len(q)-1]), "部") ||  strings.Contains(common.GetString(q[len(q)-1]), "心") {

			} else {
				//插入
				team_name := common.GetString(dep["dep_name"])
				resume := "LDAP"
				_, err1 := falcon.InsertTeam(team_name, resume, users)
				if err1 != nil {
					beego.Info("falcon插入ldap用户组数据失败===>", err1)
				} else {
					beego.Info("falcon插入ldap用户组数据成功===>", team_name)
				}
			}
		} else {
			// 批量更新接口并发会造成死锁
			team_name := common.GetString(dep["dep_name"])
			resume := "LDAP"
			team_id := falcon_team[common.GetString(dep["dep_name"])].Team.Id
			if common.GetInt(users) != common.GetInt(falcon_team[common.GetString(dep["dep_name"])].Users) {
				_, err2 := falcon.UpdateTeam(team_id, team_name, resume, users)
				if err2 != nil {
					beego.Info("falcon更新ldap用户组数据失败===>", err2)
				} else {
					beego.Info("falcon更新ldap用户组数据成功===>", team_name)
				}
			} else {
				beego.Info("Falcon与ldap用户组数据一致无需变更")
			}

		}
	}
}


func sync_jumpserver_group(jumpserver_groups map[string]js_api.Groups){
	for _, dep := range dep_name {
		if _, ok := jumpserver_groups[common.GetString(dep["dep_name"])]; !ok {
			q := strings.Split(common.GetString(dep["dep_name"]), "/")
			if strings.Contains(common.GetString(q[len(q) - 1]), "部") || strings.Contains(common.GetString(q[len(q) - 1]), "心") {
				continue
			} else {
				var users []string
				for _, j := range dep_user {
					users = append(users, common.GetString(j["dep_user"]))
				}
				_, err := jumpserver.CreateGroup(common.GetString(dep["dep_name"]), "LDAP", users)
				if err != nil {
					beego.Info("插入Jumpserver groups数据失败===>", err)
				} else {
					beego.Info("插入Jumpserver groups数据成功===>", common.GetString(dep["dep_name"]))
				}
			}
		} else {
			// 批量更新接口并发会造成死锁
			var users []string
			for _, j := range dep_user {
				users = append(users, common.GetString(j["dep_user"]))
			}
			if common.GetString(dep["dep_name"]) != common.GetString(jumpserver_groups[common.GetString(dep["dep_name"])].Name) ||
				common.GetString(users) != common.GetString(jumpserver_groups[common.GetString(dep["dep_name"])].Users) ||
				 common.GetString(jumpserver_groups[common.GetString(dep["dep_name"])].Comment) != "LDAP" {
				_, err := jumpserver.UpdateGroup(common.GetString(jumpserver_groups[common.GetString(dep["dep_name"])].Id),common.GetString(dep["dep_name"]), "LDAP", users)
				if err != nil {
					beego.Info("更新Jumpserver groups数据失败===>", err)
				} else {
					beego.Info("更新Jumpserver groups数据成功===>", common.GetString(dep["dep_name"]))
				}
			} else {

			}
		}
	}
}


func deleteFalconUser(falcon_users map[string]api.GetUserByName) {
	for _, user := range falcon_users {
		////如果falcon有数据，反查ldap是否有数据，如果有就更新，否则删除falcon的数据
		if _, ok := LDAP_USERS[common.GetString(user.Name)]; !ok {
			if user.Name != "root" {
				_, res2 := falcon.DeleteUser(user.Id)
				beego.Info("DeleteUser===>", res2)
				if res2 != nil {
					beego.Error(res2.Error())
				} else {
					beego.Info("已删除Falcon无效LDAP用户:", user.Name)
				}
			}
		}
	}
}

func deleteJumpserverUser(jumpserver_users map[string]js_api.Users) {
	for _, user := range jumpserver_users {
		////如果falcon有数据，反查ldap是否有数据，如果有就更新，否则删除falcon的数据
		if _, ok := LDAP_USERS[common.GetString(user.Username)]; !ok {
			if user.Role != "admin" && user.Comment == "LDAP" {
				res2 := jumpserver.DeleteUser(user.Id)
				if res2 != nil {
					beego.Error(res2.Error())
				} else {
					beego.Info("已删除jumpserver无效LDAP用户:", user.Username)
				}
			}
		}
	}
}

func deleteFalconTeam(falcon_team map[string]api.GetTeamListData) {
	//FalconDepList, _ := falcon.GetTeamList("/")
	for _, dep := range falcon_team {
		////如果falcon有数据，反查ldap是否有数据，如果有就更新，否则删除falcon的数据
		if _, ok := LDAP_TEAMS[common.GetString(dep.Team.Name)]; !ok {
			//删除在ldap中不存在的falcon 用户组信息 同时只能删除falcon 用户组中字段creator为1
			if dep.Team.Creator == 1 && dep.Team.Resume == "LDAP" {
				_, err2 := falcon.DeleteTeam(common.GetInt(strconv.Itoa(dep.Team.Id)))
				if err2 != nil {
					beego.Error(err2.Error())
				} else {
					beego.Info("已删除falcon无效用户组:", dep.Team.Name)
				}
			}
		}
	}
}

func deleteJumpserverGroup(jumpserver_groups map[string]js_api.Groups) {
	for _, dep := range jumpserver_groups {
		////如果falcon有数据，反查ldap是否有数据，如果有就更新，否则删除falcon的数据
		if _, ok := LDAP_TEAMS[common.GetString(dep.Name)]; !ok {
			//删除在ldap中不存在的jumpserver用户组信息
			if dep.Comment == "LDAP" {
						err := jumpserver.DeleteGroup(common.GetString(dep.Id))
						if err != nil {
							beego.Info("删除Jumpserver groups数据失败===>", err)
						} else {
							beego.Info("删除Jumpserver groups数据成功===>", dep.Name)
						}
			}
		}
	}
}
