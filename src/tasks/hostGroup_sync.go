package tasks

import (
	"runtime"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"library/falcon"
	"library/cmdb"
	"strconv"
	"library/jumpserver"
	"library/jumpserver/api"
	"strings"
	"github.com/astaxie/beego/utils"
	"library/common"
)

var SYSTEM_USER map[string]string
var ADMIN_USER map[string]string
var JUMPSERVER_NODES map[string]jumpserver.Nodes
var JUMPSERVER_ASSETS map[string]interface{}

func init() {
	orm.Debug = false
	SYSTEM_USER = make(map[string]string)
	ADMIN_USER = make(map[string]string)
	JUMPSERVER_NODES = make(map[string]jumpserver.Nodes)
	JUMPSERVER_ASSETS = make(map[string]interface{})
}

//蓝鲸CMDB、Falcon+ 主机组/主机信息同步
func HostGroupSync() error {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("HostGroupSync 蓝鲸CMDB、Falcon+ 主机组/主机信息同步:", panic_err)
			beego.Error("HostGroupSync 蓝鲸CMDB、Falcon+ 主机组/主机信息同步主程序错误详细信息:", string(buf[0:c]))
		}
	}()
	beego.Debug("蓝鲸CMDB、Falcon+ 主机组/主机信息同步 Start." + time.Now().Format("2006-01-02 15:04:05"))

	go sync_falcon()
	go sync_jump_server()

	beego.Debug("蓝鲸CMDB、Falcon+ 主机组/主机信息同步 Stop." + time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func sync_falcon() {
	beego.Debug("Falcon主机组/主机数据同步 Start")
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Falcon主机组/主机数据同步主程序错误:", panic_err)
			beego.Error("Falcon主机组/主机数据同步主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	cmdb_hostgroups, err := cmdb.GetAllHostGroups()

	if err != nil {
		panic(err)
	}
	if len(cmdb_hostgroups) == 0 {
		beego.Debug("Falcon主机组/主机数据同步 Stop, CMDB未获取到主机组数据." + time.Now().Format("2006-01-02 15:04:05"))
		return
	}

	falcon_hostgroups, err := falcon.GetAllHostGroups()
	if err != nil {
		panic(err)
	}

	for _, cmdb := range cmdb_hostgroups {
		hostGroup := cmdb.Groupname
		if _, ok := falcon_hostgroups[hostGroup]; !ok {
			create_falcon_hostgroup(cmdb)
		} else {
			// 批量更新接口并发会造成死锁
			update_falcon_hostgroup(falcon_hostgroups[hostGroup].HGID, cmdb)
		}
	}

	for _, falcon := range falcon_hostgroups {
		hostGroup := falcon.HGName
		if _, ok := cmdb_hostgroups[hostGroup]; !ok {
			// 只删除root用户创建的hostgroup
			if falcon.CreateUser == "root" {
				go delete_falcon_hostgroup(falcon.HGID, falcon.HGName)
			}
		}
	}

	beego.Debug("Falcon主机组/主机数据同步 Stop")
}

func sync_jump_server() {
	beego.Debug("JumpServer主机组/主机数据同步 Start")
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("JumpServer主机组/主机数据同步主程序错误:", panic_err)
			beego.Error("JumpServer主机组/主机数据同步主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	cmdb_Nodes, err := cmdb.GetAllNodes()
	if err != nil {
		panic(err)
	}

	init_jumpserver()
	utils.Display("cmdb_Nodes", cmdb_Nodes)
	cNodes := make(map[string]string)
	for _, cmdbNode := range cmdb_Nodes {
		if len(cmdbNode) > 0 {
			for key, cmdb_node := range cmdbNode {
				if _, ok := JUMPSERVER_NODES[key]; !ok {
					value := key + "-" + cmdb_node.NodeName
					var node jumpserver.Nodes
					node.Id = create_jumpserver_node(value)
					JUMPSERVER_NODES[key] = node
				}
				update_jumpserver_node(cmdb_node, JUMPSERVER_NODES[key])
				if _, ok := cNodes[key]; !ok {
					cNodes[key] = cmdb_node.NodeName
				}
			}
		}
	}

	delete_jumpserver_node(cNodes)

	//同步 cmdb组 到jumpserver
	cmdb_hostgroups, err := cmdb.GetAllHostGroups()
	if err != nil {
		panic(err)
	}
	if len(cmdb_hostgroups) == 0 {
		beego.Debug("Jumpserver用户组数据同步 Stop, CMDB未获取到主机组数据." + time.Now().Format("2006-01-02 15:04:05"))
		return
	}

	jumpserver_usergroups, err := jumpserver.GetAllGroups()
	if err != nil {
		panic(err)
	}

	for _, cmdb := range cmdb_hostgroups {
		hostGroup := cmdb.Groupname
		//同步jumpserver cmdb用户组
		if _,ok := jumpserver_usergroups[hostGroup]; !ok {
			create_jumpserver_usergroup(cmdb)
		} else {
			update_jumpserver_usergroup(jumpserver_usergroups[hostGroup].Id, cmdb)
		}
	}

	for _, jump := range jumpserver_usergroups {
		userGroup := jump.Name
		if _, ok := cmdb_hostgroups[userGroup]; !ok {
			// 只删除CMDB 用户组
			if jump.Comment == "CMDB" {
				go delete_jumpserver_usergroup(jump.Id)
			}
		}
	}

	jumpserver_group_list, _ := jumpserver.GetAllGroups()
	//同步更新jumpserver cmdb用户组的主机组用户
	for _, cmdb := range cmdb_hostgroups {
		if common.GetString(cmdb.Operator) == "[]" {
			jumpserver_user, _:= jumpserver.GetUsersName("shuguang")
			id := jumpserver_user[0].Id
			groups_display := jumpserver_user[0].Groups_display + " " +cmdb.Groupname
			var groups []string
			for _, k := range jumpserver_user[0].Groups {
				groups = append(groups, k)
			}
			groups = append(groups, jumpserver_group_list[cmdb.Groupname].Id)
			name := jumpserver_user[0].Username
			wechat := jumpserver_user[0].Wechat
			cnname := jumpserver_user[0].Name
			email := jumpserver_user[0].Email
			comment := "LDAP"
			phone := jumpserver_user[0].Phone
			if groups_display != jumpserver_user[0].Groups_display {
				_, err := jumpserver.UpdateUser(id, cnname, name, email, phone,comment,wechat,groups_display,groups)
				if err != nil {
					beego.Info("Jumpserver更新ldap数据到CMDB组失败===>", err)
				} else {
					beego.Info("Jumpserver更新ldap数据到CMDB组成功===>", name)
				}
			} else {
				beego.Info("Jumpserver与ldap用户数据一致无需变更")
			}
		} else {
			for _, k := range cmdb.Operator {
				jumpserver_user, _:= jumpserver.GetUsersName(common.GetString(k))
				id := jumpserver_user[0].Id
				groups_display := jumpserver_user[0].Groups_display + " " +cmdb.Groupname
				var groups []string
				for _, k := range jumpserver_user[0].Groups {
					groups = append(groups, k)
				}
				groups = append(groups, jumpserver_group_list[cmdb.Groupname].Id)
				name := jumpserver_user[0].Username
				wechat := jumpserver_user[0].Wechat
				cnname := jumpserver_user[0].Name
				email := jumpserver_user[0].Email
				comment := "LDAP"
				phone := jumpserver_user[0].Phone
				if groups_display != jumpserver_user[0].Groups_display {
					_, err := jumpserver.UpdateUser(id, cnname, name, email, phone,comment,wechat,groups_display,groups)
					if err != nil {
						beego.Info("Jumpserver更新ldap数据到CMDB组失败===>", err)
					} else {
						beego.Info("Jumpserver更新ldap数据到CMDB组成功===>", name)
					}
				} else {
					beego.Info("Jumpserver与ldap用户数据一致无需变更")
				}
			}
		}
	}

	jumpserver_Nodes, err := jumpserver.GetAllNodes()

	for _,k := range jumpserver_Nodes {
		if k.Id == k.Parent {
			continue
		} else {
			var tree_temp []string
			tree_temp = append(tree_temp,common.GetString(k.Value))
			getTree(k.Id,k.Parent,jumpserver_Nodes,tree_temp,jumpserver_group_list)
		}
	}
	utils.Display("cmdb_hostgroups===cmdb_hostgroups==>", cmdb_hostgroups)
	//反查删除不存在组的授权规则
	delete_AssetPermissions(jumpserver_group_list)
	beego.Debug("JumpServer主机组/主机数据同步 Stop")
}


func getTree(id string, parent string,jumpserver_Nodes map[string]jumpserver.Nodes, tree_temp []string, jumpserver_group_list map[string]api.Groups){
	if len(tree_temp) < 4 {
		for _, v := range jumpserver_Nodes {
			if v.Id != id {
				if v.Id == parent {
					tree_temp = append(tree_temp,common.GetString(v.Value))
					getTree(v.Id, v.Parent,jumpserver_Nodes, tree_temp, jumpserver_group_list)
				}
			}
		}
	} else {
		if len(tree_temp) == 4 {
			var temp_group []string
			var temp_node []string
			var system_user []string
			var users []string
			var assets []string
			dep_one := str_Split(tree_temp[2])
			dep_two := str_Split(tree_temp[1])
			dep_three := str_Split(tree_temp[0])
			temp_group = append(temp_group, jumpserver_group_list[dep_one + "/" + dep_two + "/" + dep_three].Id)
			temp_node =  append(temp_node, JUMPSERVER_NODES[strings.Split(tree_temp[0],"-")[0]].Id)
			system_user = append(system_user, SYSTEM_USER["jumper"])
			name := dep_one + "/" + dep_two + "/" + dep_three
			is_active := true
			comment := "CMDB-GROUP"
			users = append(users, "")
			assets = append(assets, "")

			AssetPermissions, _:= jumpserver.GetAssetPermissions()
			if _, ok := AssetPermissions[name]; !ok {
				_, err := jumpserver.CreateAssetPermission(name, temp_group, temp_node, system_user, is_active, comment)
				if err != nil {
					beego.Info("Jumpserver创建授权规则失败===>", err)
				} else {
					beego.Info("Jumpserver创建授权规则成功===>", name)
					if _, ok := jumpserver_Nodes[strings.Split(tree_temp[0],"-")[0]]; ok {
						for k,v := range jumpserver_Nodes[strings.Split(tree_temp[0],"-")[0]].Assets {
							beego.Info("AssetPermissions---->", v.Hostname, AssetPermissions[common.GetString(v.Hostname)])
							if k != "" {
								if _, ok := AssetPermissions[common.GetString(v.Hostname)]; !ok {
									var assets_users []string
									var assets_assets []string
									var assets_system_user []string
									var assets_temp_group []string
									var assets_temp_node []string

									assets_temp_group = append(assets_temp_group, "")
									assets_temp_node =  append(assets_temp_node, "")
									assets_is_active := true
									assets_comment := "CMDB-NODE"
									assets_name := v.Hostname
									assets_jumpserver_user, _:= jumpserver.GetUsersName("shuguang")
									assets_id := assets_jumpserver_user[0].Id
									assets_users = append(assets_users, assets_id)
									assets_assets = append(assets_assets, v.Id)
									assets_system_user = append(assets_system_user, SYSTEM_USER["jumper"])
									_, err := jumpserver.CreateAssetPermissions(assets_name, assets_users, assets_assets, assets_system_user, assets_is_active, assets_comment)
									if err != nil {
										beego.Info("Jumpserver创建[资产]授权规则失败===>", err)
									} else {
										beego.Info("Jumpserver创建[资产]授权规则成功===>", name)
										JUMPSERVER_ASSETS[v.Hostname] = v.Hostname
									}
								} else {
									beego.Info("update==>【资产】授权")
									var assets_users []string
									var assets_assets []string
									var assets_system_user []string
									var assets_temp_group []string
									var assets_temp_node []string

									assets_temp_group = append(assets_temp_group, "")
									assets_temp_node =  append(assets_temp_node, "")
									assets_is_active := true
									assets_comment := "CMDB-NODE"
									assets_name := v.Hostname
									//assets_jumpserver_user, _:= jumpserver.GetUsersName("shuguang")
									//assets_id := assets_jumpserver_user[0].Id
									assets_id := v.AdminUser
									assets_users = append(assets_users, assets_id)
									assets_assets = append(assets_assets, v.Id)
									assets_system_user = append(assets_system_user, SYSTEM_USER["jumper"])
									_, err := jumpserver.UpdateAssetPermission(AssetPermissions[common.GetString(v.Hostname)].Id, assets_name, assets_users, assets_assets, assets_system_user, assets_is_active, assets_comment)
									if err != nil {
										beego.Info("Jumpserver更新[资产]授权规则失败===>", err)
									} else {
										beego.Info("Jumpserver更新[资产]授权规则成功===>", name)
										JUMPSERVER_ASSETS[v.Hostname] = v.Hostname
									}
								}

							}
						}
					}
				}
			} else {
				_, err := jumpserver.UpdateAssetPermissions(AssetPermissions[name].Id, name, temp_group, temp_node, system_user, is_active, comment)
				if err != nil {
					beego.Info("Jumpserver更新授权规则失败===>", err)
				} else {
					beego.Info("Jumpserver更新授权规则成功===>", name)
					if _, ok := jumpserver_Nodes[strings.Split(tree_temp[0],"-")[0]]; ok {
						for k,v := range jumpserver_Nodes[strings.Split(tree_temp[0],"-")[0]].Assets {
							beego.Info("AssetPermissions---->", v.Hostname, AssetPermissions[common.GetString(v.Hostname)])
							if k != "" {
								if _, ok := AssetPermissions[common.GetString(v.Hostname)]; !ok {
									var assets_users []string
									var assets_assets []string
									var assets_system_user []string
									var assets_temp_group []string
									var assets_temp_node []string

									assets_temp_group = append(assets_temp_group, "")
									assets_temp_node =  append(assets_temp_node, "")
									assets_is_active := true
									assets_comment := "CMDB-NODE"
									assets_name := v.Hostname
									assets_jumpserver_user, _:= jumpserver.GetUsersName("shuguang")
									assets_id := assets_jumpserver_user[0].Id
									assets_users = append(assets_users, assets_id)
									assets_assets = append(assets_assets, v.Id)
									assets_system_user = append(assets_system_user, SYSTEM_USER["jumper"])
									_, err := jumpserver.CreateAssetPermissions(assets_name, assets_users, assets_assets, assets_system_user, assets_is_active, assets_comment)
									if err != nil {
										beego.Info("Jumpserver创建[资产]授权规则失败===>", err)
									} else {
										beego.Info("Jumpserver创建[资产]授权规则成功===>", assets_name)
										JUMPSERVER_ASSETS[v.Hostname] = v.Hostname
									}
								} else {
									beego.Info("update==>【资产】授权")
									var assets_users []string
									var assets_assets []string
									var assets_system_user []string
									var assets_temp_group []string
									var assets_temp_node []string

									assets_temp_group = append(assets_temp_group, "")
									assets_temp_node =  append(assets_temp_node, "")
									assets_is_active := true
									assets_comment := "CMDB-NODE"
									assets_name := v.Hostname
									//assets_jumpserver_user, _:= jumpserver.GetUsersName("shuguang")
									//assets_id := assets_jumpserver_user[0].Id
									assets_id := v.AdminUser
									assets_users = append(assets_users, assets_id)
									assets_assets = append(assets_assets, v.Id)
									assets_system_user = append(assets_system_user, SYSTEM_USER["jumper"])
									_, err := jumpserver.UpdateAssetPermission(AssetPermissions[common.GetString(v.Hostname)].Id, assets_name, assets_users, assets_assets, assets_system_user, assets_is_active, assets_comment)
									if err != nil {
										beego.Info("Jumpserver更新[资产]授权规则失败===>", err)
									} else {
										beego.Info("Jumpserver更新[资产]授权规则成功===>", assets_name)
										JUMPSERVER_ASSETS[v.Hostname] = v.Hostname
									}
								}

							}
						}
					}
				}
			}
		} else {
			//utils.Display("含有资产的节点===》", tree_temp)
		}
	}
}

func delete_AssetPermissions(jumpserver_group_list map[string]api.Groups) {
	AssetPermissions, _:= jumpserver.GetAssetPermissions()
	for k, j := range AssetPermissions {
		if _, ok := jumpserver_group_list[k]; !ok {
			if _, ok := JUMPSERVER_ASSETS[k]; !ok {
				utils.Display("授权规则中不存在该组=》",k)
				err := jumpserver.DeleteAssetPermissions(j.Id)
				if err != nil {
					beego.Info("Jumpserver删除授权规则失败===>", err)
				} else {
					beego.Info("Jumpserver删除授权规则成功===>", j)
				}
			}
		}
	}
}

//节点名称处理
func str_Split(str string) (string) {
	str_arr := strings.Split(str,"-")
	res := ""
	if len(str_arr) > 2 {
		for i,k := range str_arr {
			if i != 0 {
				if res == "" {
					res = k
				} else {
					res = res + "-" + k
				}
			}
		}
	} else {
		res = str_arr[1]
	}
	return res
}
func create_falcon_hostgroup(cmdb cmdb.HostGroup) {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Falcon新增主机组主程序错误:", panic_err)
			beego.Error("Falcon新增主机组主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	gpId, err := falcon.CreateFalconHostGroup(cmdb.Groupname)
	if err != nil {
		panic(err)
	}
	if len(cmdb.Hosts) > 0 {
		_, err := falcon.AddHostsToFalconHostGroup(gpId, cmdb.Hosts)
		if err != nil {
			panic(err)
		}
	}

	// 创建hostgroup对应用户组
	users, err := falcon.GetFalconUsersId(cmdb.Operator)
	if err != nil {
		panic(err)
	}

	_, err = falcon.InsertTeam(cmdb.Groupname, "CMDB", users)
	if err != nil {
		panic(err)
	}

	// 创建hostgroup对应告警模板
	tplId, err := falcon.CreateTemplate(cmdb.Groupname, 0)
	if err != nil {
		panic(err)
	}

	_, err = falcon.CreateTplAction(tplId, cmdb.Groupname)
	if err != nil {
		panic(err)
	}

	procStrategy, _ := getProcStrategy(cmdb.Procs)
	baseStrategy := falcon.GetBaseTMP()
	baseStrategy = append(baseStrategy, procStrategy...)
	for _, strategy := range baseStrategy {
		go falcon.CreateStrategy(tplId, strategy)
	}

	// nodata 配置
	_, err = falcon.CreateNodata(cmdb.Groupname)
	if err != nil {
		panic(err)
	}

	// hostgroup绑定模板
	_, err = falcon.BindTemplateToHostGroup(tplId, gpId)
	if err != nil {
		panic(err)
	}

	beego.Debug("Falcon新增主机组成功 " + strconv.Itoa(gpId))
}

func update_falcon_hostgroup(hgId int, cmdb cmdb.HostGroup) {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Falcon更新主机组主程序错误:", panic_err)
			beego.Error("Falcon更新主机组主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	// 更新主机组主机
	if len(cmdb.Hosts) > 0 {
		_, err := falcon.AddHostsToFalconHostGroup(hgId, cmdb.Hosts)
		if err != nil {
			panic(err)
		}
	}

	// 更新team信息
	team, err := falcon.GetTeamList(cmdb.Groupname)
	if err != nil {
		panic(err)
	}

	users, err := falcon.GetFalconUsersId(cmdb.Operator)
	if err != nil {
		panic(err)
	}

	if len(team) == 0 {
		// 创建hostgroup对应用户组
		_, err = falcon.InsertTeam(cmdb.Groupname, "CMDB", users)
		if err != nil {
			panic(err)
		}
	} else {
		if len(users) > 0 {
			_, err = falcon.UpdateTeam(team[0].Team.Id, team[0].Team.Name, team[0].Team.Resume, users)
			if err != nil {
				panic(err)
			}
		}
	}

	// 更新进程及端口告警模板
	tpl, err := falcon.GetTemplateByName(cmdb.Groupname)
	if err != nil {
		if err.Error() == "record not found" {
			// 创建hostgroup对应告警模板
			tplId, err := falcon.CreateTemplate(cmdb.Groupname, 0)
			if err != nil {
				panic(err)
			}

			_, err = falcon.CreateTplAction(tplId, cmdb.Groupname)
			if err != nil {
				panic(err)
			}

			procStrategy, _ := getProcStrategy(cmdb.Procs)
			baseStrategy := falcon.GetBaseTMP()
			baseStrategy = append(baseStrategy, procStrategy...)
			for _, strategy := range baseStrategy {
				go falcon.CreateStrategy(tplId, strategy)
			}
		} else {
			panic(err)
		}
	} else {
		falconProcStr := make(map[string]int)
		for _, falconStr := range tpl.Stratges {
			mt := falconStr.Metric + "/" + falconStr.Tags
			if _, ok := falconProcStr[mt]; !ok {
				falconProcStr[mt] = falconStr.Id
			}
		}
		_, procStrategy := getProcStrategy(cmdb.Procs)
		for key, str := range procStrategy {
			if _, ok := falconProcStr[key]; !ok {
				go falcon.CreateStrategy(tpl.Template.TplId, str)
			}
		}
	}

	// nodata 配置
	nodata, err := falcon.GetNodataInfoByName(cmdb.Groupname)
	if err != nil {
		panic(err)
	}

	if nodata.Id == 0 {
		_, err = falcon.CreateNodata(cmdb.Groupname)
		if err != nil {
			panic(err)
		}
	}

	beego.Debug("Falcon更新主机组成功")
}

func delete_falcon_hostgroup(hgId int, hgName string) {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Falcon删除主机组主程序错误:", panic_err)
			beego.Error("Falcon删除主机组主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	// 删除用户组
	team, err := falcon.GetTeamList(hgName)
	if err != nil {
		panic(err)
	}
	_, err = falcon.DeleteTeam(team[0].Team.Id)
	if err != nil {
		panic(err)
	}

	// 删除模板
	tpl, err := falcon.GetTemplateByName(hgName)
	if err != nil {
		if err.Error() != "record not found" {
			panic(err)
		}
	} else {
		if _, err := falcon.DeleteTemplate(tpl.Template.TplId); err != nil {
			panic(err)
		}
	}

	// 删除hostgroup
	msg, err := falcon.DeleteFalconHostGroup(hgId)
	if err != nil {
		panic(err)
	}
	beego.Debug("Falcon删除主机组成功 " + msg)
}

func getProcStrategy(procs []cmdb.Proc) ([]falcon.Strategy, map[string]falcon.Strategy) {
	ret := []falcon.Strategy{}
	mps := make(map[string]falcon.Strategy)
	for _, proc := range procs {
		procStr := falcon.Strategy{"proc.num", "name=" + proc.BkProcessName, 3, 0, proc.BkProcessName + "进程异常告警", "all(#2)", "==", "0"}
		ret = append(ret, procStr)
		mt := "proc.num" + "/" + "name=" + proc.BkProcessName
		if _, ok := mps[mt]; !ok {
			mps[mt] = procStr
		}
		if proc.Port != "" {
			ports := strings.Split(proc.Port, ",")
			for _, port := range ports {
				portStr := falcon.Strategy{"net.port.listen", "port=" + port, 3, 0, port + "端口异常告警", "all(#3)", "==", "0"}
				ret = append(ret, portStr)
				pt := "net.port.listen" + "/" + "port=" + port
				if _, ok := mps[pt]; !ok {
					mps[pt] = portStr
				}
			}
		}
	}

	return ret, mps
}

func init_jumpserver() {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Jumpserver获取用户主程序错误:", panic_err)
			beego.Error("Jumpserver获取用户主程序错误:", string(buf[0:c]))
		}
	}()

	adminUsers, err := jumpserver.GetAdminUsers()
	if err != nil {
		panic(err)
	}

	for _, adminU := range adminUsers {
		if _, ok := ADMIN_USER[adminU.Username]; !ok {
			ADMIN_USER[adminU.Username] = adminU.Id
		}
	}
	if _, ok := ADMIN_USER["root"]; !ok {
		user, err := jumpserver.CreateAdminUser("root", "root", "root")
		if err != nil {
			panic(err)
		}
		ADMIN_USER["root"] = user.Id
	}

	sysUsers, err := jumpserver.GetSystemUsers()
	if err != nil {
		panic(err)
	}

	for _, sysU := range sysUsers {
		if _, ok := SYSTEM_USER[sysU.Username]; !ok {
			SYSTEM_USER[sysU.Username] = sysU.Id
		}
	}
	if _, ok := SYSTEM_USER["jumper"]; !ok {
		user, err := jumpserver.CreateSystemUser("jumper", "jumper")
		if err != nil {
			panic(err)
		}
		SYSTEM_USER["jumper"] = user.Id
	}

	jumpserver_Nodes, err := jumpserver.GetAllNodes()
	if err != nil {
		panic(err)
	}
	JUMPSERVER_NODES = jumpserver_Nodes
}

func create_jumpserver_node(value string) string {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Jumpserver新增节点主程序错误:", panic_err)
			beego.Error("Jumpserver新增节点主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	node, err := jumpserver.CreateNode(value)
	if err != nil {
		panic(err)
	}
	return node.Id
}

func update_jumpserver_node(cmdbNodes cmdb.Nodes, jumpserverNodes jumpserver.Nodes) {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Jumpserver更新节点主程序错误:", panic_err)
			beego.Error("Jumpserver更新节点主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	if len(cmdbNodes.Children) > 0 {
		var childNodes []string
		for _, node := range cmdbNodes.Children {
			if _, ok := JUMPSERVER_NODES[node.NodeKey]; !ok {
				value := node.NodeKey + "-" + node.NodeName
				var jnode jumpserver.Nodes
				jnode.Id = create_jumpserver_node(value)
				JUMPSERVER_NODES[node.NodeKey] = jnode
			} else {
				value := node.NodeKey + "-" + node.NodeName
				_, err := jumpserver.UpdateNode(value, JUMPSERVER_NODES[node.NodeKey].Id)
				if err != nil {
					panic(err)
				}
			}
			childNodes = append(childNodes, JUMPSERVER_NODES[node.NodeKey].Id)
		}

		err := jumpserver.AddNodeChildren(jumpserverNodes.Id, childNodes)
		if err != nil {
			panic(err)
		}
	}

	if len(cmdbNodes.Hosts) > 0 {
		sync_jumpserver_assets(cmdbNodes.Hosts, jumpserverNodes)
	}
}

func sync_jumpserver_assets(Hosts map[string]cmdb.Host, node jumpserver.Nodes) {
	var addHosts, removeHosts []string
	assets := node.Assets
	for ip, host := range Hosts {
		if _, ok := assets[ip]; !ok {
			assetInfo, err := jumpserver.GetAssetInfo(host.BkHostName, host.BkHostInnerip)
			if err != nil {
				beego.Error("Jumpserver新增节点主机主程序错误:", err)
				return
			}
			var hostId string
			if len(assetInfo) > 0 {
				hostId = assetInfo[0].Id
			} else {
				if _, ok := ADMIN_USER["root"]; ok {
					hostId, err = jumpserver.CreateAsset(host.BkHostName, host.BkHostInnerip, ADMIN_USER["root"])
					if err != nil {
						beego.Error("Jumpserver新增节点主机主程序错误:", err)
						return
					}
				} else {
					beego.Error("Jumpserver新增节点主机主程序错误: 没有获取到root管理用户", )
					return
				}
			}
			addHosts = append(addHosts, hostId)
		}
	}

	for jip, asset := range assets {
		if _, ok := Hosts[jip]; !ok {
			removeHosts = append(removeHosts, asset.Id)
		}
	}

	if len(addHosts) > 0 {
		jumpserver.NodesAssetsAdd(node.Id, addHosts)
	}

	if len(removeHosts) > 0 {
		jumpserver.NodesAssetsRemove(node.Id, removeHosts)
	}
}

func delete_jumpserver_node(cNodes map[string]string) {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			c := runtime.Stack(buf, false)
			beego.Error("Jumpserver删除节点主程序错误:", panic_err)
			beego.Error("Jumpserver删除节点主程序错误详细信息:", string(buf[0:c]))
		}
	}()

	for key, jnode := range JUMPSERVER_NODES {
		if _, ok := cNodes[key]; !ok {
			go jumpserver.DeleteNode(jnode.Id)
		}
	}
}

func create_jumpserver_usergroup(cmdb cmdb.HostGroup) {
	var users []string
	_, err := jumpserver.CreateGroup(cmdb.Groupname, "CMDB", users)
	if err != nil {
		beego.Info("插入Jumpserver CMDB-groups数据失败===>", err)
	} else {
		beego.Info("插入Jumpserver CMDB-groups数据成功===>", cmdb.Groupname)
	}
}

func update_jumpserver_usergroup(jump_id string, cmdb cmdb.HostGroup) {
	var users []string
	_, err := jumpserver.UpdateGroup(jump_id,cmdb.Groupname, "CMDB", users)
	if err != nil {
		beego.Info("更新Jumpserver CMDB-groups数据失败===>", err)
	} else {
		beego.Info("更新Jumpserver CMDB-groups数据成功===>", cmdb.Groupname,cmdb.Operator, cmdb)
	}
}

func delete_jumpserver_usergroup(jump_id string) {
	err := jumpserver.DeleteGroup(jump_id)
	if err != nil {
		beego.Info("删除Jumpserver CMDB-groups数据失败===>", err)
	} else {
		beego.Info("删除Jumpserver CMDB-groups数据成功===>", jump_id)
	}
}