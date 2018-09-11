package cmdb

import (
)
import (
	"github.com/pkg/errors"
	"library/common"
	"library/cmdb/api"
	"strings"
	"regexp"
	"github.com/astaxie/beego"
)

type HostGroup struct {
	Groupname string
	Operator  []string
	Hosts     []string
	Procs     []Proc
}

type Proc struct {
	BkProcessName string
	Port 	      string
}

func GetAllHostGroups() (map[string]HostGroup, error){
	cmdbApi := api.New()
	ret := make(map[string]HostGroup)
	bizs, err := cmdbApi.SearchBusiness("0", map[string]interface{}{})
	if err != nil {
		return ret, err
	}
	if !bizs.Result {
		return ret, errors.New(bizs.Message)
	}

	for _, biz := range bizs.Data.Info {
		procs, err := GetProcsByBiz("0", biz.BkBizId)
		if err != nil {
			return ret, err
		}
		sets, err := cmdbApi.SearchSet("0", biz.BkBizId, map[string]interface{}{})
		if err != nil {
			return ret, err
		}
		if !sets.Result {
			return ret, errors.New(sets.Message)
		}
		for _, set := range  sets.Data.Info {
			modules, err := cmdbApi.SearchModule("0", biz.BkBizId, set.BkSetId, map[string]interface{}{})
			if err != nil {
				return ret, err
			}
			if !modules.Result {
				return ret, errors.New(modules.Message)
			}
			for _, module := range modules.Data.Info {
				hostGroupName := biz.BkBizName + "/" + set.BkSetName + "/" + module.BkModuleName
				_, ok := ret[hostGroupName]
				if !ok {
					var hostGroup HostGroup
					query := map[string]interface{}{
						"condition": []interface{}{
							map[string]interface{}{
								"bk_obj_id": "module",
								"fields": []string{},
								"condition": []interface{}{
									map[string]interface{}{
										"field":    "bk_module_id",
										"operator": "$eq",
										"value":    module.BkModuleId,
									},
								},
							},
						},
					}
					hosts, err := cmdbApi.SearchHosts(query)
					if err != nil {
						return ret, err
					}
					if !hosts.Result {
						return ret, errors.New(hosts.Message)
					}
					hostIps := []string{}
					for _, host := range hosts.Data.Info {
						hostIps = append(hostIps, host.Host.BkHostName + "/" + host.Host.BkHostInnerip)
					}

					var hp []Proc
					if _, ok := procs[module.BkModuleName]; ok {
						hp = procs[module.BkModuleName]
					}

					hostGroup.Groupname = hostGroupName
					hostGroup.Hosts = hostIps
					hostGroup.Procs = hp
					hostGroup.Operator = handleOperators(biz.BkBizDeveloper, biz.BkBizMaintainer, module.Operator, module.BkBakOperator)
					ret[hostGroupName] = hostGroup
				}
			}
		}
	}

	return ret, nil
}

func GetProcsByBiz(ownerID string, bizID int) (map[string][]Proc, error) {
	cmdbApi := api.New()
	ret := make(map[string][]Proc)

	query := map[string]interface{}{
		"condition": map[string]interface{}{
			"bk_biz_id":    bizID,
		},
		"fields": []string{"bk_biz_id", "bk_process_id", "bk_process_name", "port"},
		"page": map[string]interface{}{
			"start":    0,
			"limit": 9999,
			"sort":    "bk_process_id",
		},
	}
	procs, err := cmdbApi.SearchProc(ownerID, bizID, query)
	if err != nil {
		return ret, err
	}

	for _, proc := range procs.Data.Info {
		procModules, err := cmdbApi.GetProcessBindModule(ownerID, bizID, proc.BkProcessId)
		if err != nil {
			return ret, err
		}
		for _, module := range procModules.Data {
			if module.IsBind == 1 {
				if _, ok := ret[module.BkModuleName]; !ok {
					var moduleProcs []Proc
					var moduleProc Proc
					moduleProc.BkProcessName = proc.BkProcessName
					moduleProc.Port = proc.Port
					moduleProcs = append(moduleProcs, moduleProc)
					ret[module.BkModuleName] = moduleProcs
				} else {
					ret[module.BkModuleName] = append(ret[module.BkModuleName], Proc{proc.BkProcessName, proc.Port})
				}
			}
		}
	}

	return ret, nil
}

func handleOperators(bd, bm, mo, mbo string) []string{
	ret := []string{}
	bds := strings.Split(bd, ",")
	bms := strings.Split(bm, ",")
	mos := strings.Split(mo, ",")
	mbos := strings.Split(mbo, ",")

	for _, bd := range bds {
		if bd != "admin" && bd != "" && !common.InList(bd, ret) {
			ret = append(ret, bd)
		}
	}
	for _, bm := range bms {
		if bm != "admin" && bm != "" && !common.InList(bm, ret) {
			ret = append(ret, bm)
		}
	}
	for _, mo := range mos {
		if mo != "admin" && mo != "" && !common.InList(mo, ret) {
			ret = append(ret, mo)
		}
	}
	for _, mbo := range mbos {
		if mbo != "admin" && mbo != "" && !common.InList(mbo, ret) {
			ret = append(ret, mbo)
		}
	}

	return ret
}


type Nodes struct {
	NodeName 	string
	Operator 	string
	BkOperator	string
	Hosts    	map[string]Host
	Children	[]Node
}

type Node struct {
	NodeName 	string
	NodeKey		string
}

type Host struct {
	BkHostInnerip 	string
	BkHostName 	string
}

func GetAllNodes() ([]map[string]Nodes, error){
	cmdbApi := api.New()
	ret := []map[string]Nodes{}
	ret_root := make(map[string]Nodes)
	ret_biz := make(map[string]Nodes)
	ret_set := make(map[string]Nodes)
	ret_module := make(map[string]Nodes)

	_, ok := ret_root["ROOT"]
	if !ok {
		var rootNodes Nodes
		rootNodes.NodeName = "ROOT"
		var rootChilds []Node

		bizs, err := cmdbApi.SearchBusiness("0", map[string]interface{}{})
		if err != nil {
			return ret, err
		}
		if !bizs.Result {
			return ret, errors.New(bizs.Message)
		}
		for _, biz := range bizs.Data.Info {
			keyBiz := "0:" + common.GetString(biz.BkBizId)
			var bizNode Node
			bizNode.NodeName = biz.BkBizName
			bizNode.NodeKey = keyBiz
			rootChilds = append(rootChilds, bizNode)

			if _, ok := ret_biz[keyBiz]; !ok {
				var bizNodes Nodes
				bizNodes.NodeName = biz.BkBizName
				var bizChilds []Node
				sets, err := cmdbApi.SearchSet("0", biz.BkBizId, map[string]interface{}{})
				if err != nil {
					return ret, err
				}
				if !sets.Result {
					return ret, errors.New(sets.Message)
				}
				for _, set := range  sets.Data.Info {
					keySet := keyBiz + ":" + common.GetString(set.BkSetId)
					var setNode Node
					setNode.NodeName = set.BkSetName
					setNode.NodeKey = keySet
					bizChilds = append(bizChilds, setNode)

					if _, ok := ret_set[keySet]; !ok {
						var setNodes Nodes
						setNodes.NodeName = set.BkSetName
						var setChilds []Node
						modules, err := cmdbApi.SearchModule("0", biz.BkBizId, set.BkSetId, map[string]interface{}{})
						if err != nil {
							return ret, err
						}
						if !modules.Result {
							return ret, errors.New(modules.Message)
						}
						for _, module := range modules.Data.Info {
							keyModule := keySet + ":" + common.GetString(module.BkModuleId)
							var moduleNode Node
							moduleNode.NodeName = module.BkModuleName
							moduleNode.NodeKey = keyModule
							setChilds = append(setChilds, moduleNode)

							if _, ok := ret_module[keyModule]; !ok {
								var moduleNodes Nodes
								moduleNodes.NodeName = module.BkModuleName
								moduleNodes.BkOperator = module.BkBakOperator
								moduleNodes.Operator = module.Operator
								query := map[string]interface{}{
									"condition": []interface{}{
										map[string]interface{}{
											"bk_obj_id": "module",
											"fields": []string{},
											"condition": []interface{}{
												map[string]interface{}{
													"field":    "bk_module_id",
													"operator": "$eq",
													"value":    module.BkModuleId,
												},
											},
										},
									},
								}
								hosts, err := cmdbApi.SearchHosts(query)
								if err != nil {
									return ret, err
								}
								if !hosts.Result {
									return ret, errors.New(hosts.Message)
								}
								hostIps := make(map[string]Host)
								for _, host := range hosts.Data.Info {
									if _, ok := hostIps[host.Host.BkHostInnerip]; !ok {
										hostIps[host.Host.BkHostInnerip] = Host{host.Host.BkHostInnerip,host.Host.BkHostName}
									}
								}
								moduleNodes.Hosts = hostIps
								ret_module[keyModule] = moduleNodes
							}
						}
						setNodes.Children = setChilds
						ret_set[keySet] = setNodes
					}
				}
				bizNodes.Children = bizChilds
				ret_biz[keyBiz] = bizNodes
			}
		}
		rootNodes.Children = rootChilds
		ret_root["ROOT"] = rootNodes
	}

	ret = append(ret, ret_root, ret_biz, ret_set, ret_module)

	return ret, nil
}

func GetCmdbHosts(query string, limit int64, offset int64) (map[string]map[string]string, int, error) {
	match, _ := regexp.MatchString("[1-9]\\d*", query)
	beego.Info("match===>", match)
	var Total map[string]interface{}
	var Query  map[string]interface{}
	if match == true {
		Query = map[string]interface{}{
			"ip": map[string]interface{}{
				"data": []string{query},
				"exact":0,
				"flag": "bk_host_innerip|bk_host_outerip",
			},
			"condition": []interface{}{
				map[string]interface{}{
					"bk_obj_id": "host",
					"fields": []string{},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "biz",
					"fields": []string{"bk_biz_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "set",
					"fields": []string{"bk_set_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "module",
					"fields": []string{"bk_module_name"},
					"condition": []interface{}{
					},
				},
			},
			"page": map[string]interface{}{
				"start":offset,
				"limit":limit,
				"sort": "bk_host_name",
			},
			"pattern": "",
		}
		Total = map[string]interface{}{
			"ip": map[string]interface{}{
				"data": []string{query},
				"exact":0,
				"flag": "bk_host_innerip",
			},
			"condition": []interface{}{
				map[string]interface{}{
					"bk_obj_id": "host",
					"fields": []string{},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "biz",
					"fields": []string{"bk_biz_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "set",
					"fields": []string{"bk_set_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "module",
					"fields": []string{"bk_module_name"},
					"condition": []interface{}{
					},
				},
			},
			"pattern": "",
		}
	} else {
		//$regex模糊匹配
		Query = map[string]interface{}{
			"condition": []interface{}{
				map[string]interface{}{
					"bk_obj_id": "host",
					"fields": []string{},
					"condition": []interface{}{
						map[string]interface{}{
							"field": "bk_host_name",
							"operator": "$regex",
							"value": query,
						},
					},
				},
				map[string]interface{}{
					"bk_obj_id": "biz",
					"fields": []string{"bk_biz_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "set",
					"fields": []string{"bk_set_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "module",
					"fields": []string{"bk_module_name"},
					"condition": []interface{}{
					},
				},
			},
			"page": map[string]interface{}{
				"start":offset,
				"limit":limit,
				"sort": "bk_host_id",
			},
			"pattern": "",
		}
		Total = map[string]interface{}{
			"condition": []interface{}{
				map[string]interface{}{
					"bk_obj_id": "host",
					"fields": []string{},
					"condition": []interface{}{
						map[string]interface{}{
							"field": "bk_host_name",
							"operator": "$regex",
							"value": query,
						},
					},
				},
				map[string]interface{}{
					"bk_obj_id": "biz",
					"fields": []string{"bk_biz_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "set",
					"fields": []string{"bk_set_name"},
					"condition": []interface{}{
					},
				},
				map[string]interface{}{
					"bk_obj_id": "module",
					"fields": []string{"bk_module_name"},
					"condition": []interface{}{
					},
				},
			},
			"pattern": "",
		}
	}
	beego.Info("query==>", Query)
	cmdbApi := api.New()
	hostsAll, err := cmdbApi.SearchHosts(Total)
	var host_total []string
	for _, h := range hostsAll.Data.Info {
		temp := common.GetString(h.Host.BkHostName)
		host_total = append(host_total, temp)
	}
	if err != nil {
		return nil, len(host_total), err
	}
	hosts, err := cmdbApi.SearchHosts(Query)
	if err != nil {
		return nil, len(host_total), err
	}
	hostIps := make(map[string]map[string]string)
	for _, host := range hosts.Data.Info {
		temp := make(map[string]string)
		//var temp map[string]string
		temp["title"] = common.GetString(host.Host.BkHostName) + "/" + common.GetString(host.Host.BkHostInnerip)
		temp["BkHostName"] = common.GetString(host.Host.BkHostName)
		temp["BkHostInnerip"] = common.GetString(host.Host.BkHostInnerip)
		temp["BkHostOuterip"] = common.GetString(host.Host.BkHostOuterip)
		temp["BkHostType"] = common.GetString(host.Host.BkHostType)
		temp["Operator"] = common.GetString(host.Host.Operator)
		temp["BkBakOperator"] = common.GetString(host.Host.BkBakOperator)
		temp["BkManufacturer"] = common.GetString(host.Host.BkManufacturer)
		temp["BkOsName"] = common.GetString(host.Host.BkOsName)
		temp["BkOsBit"] = common.GetString(host.Host.BkOsBit)
		temp["BkCpuModule"] = common.GetString(host.Host.BkCpuModule)
		temp["BkIspName"] = common.GetString(host.Host.BkIspName)
		temp["BkStatus"] = common.GetString(host.Host.BkStatus)
		temp["BkCurrentStatus"] = common.GetString(host.Host.BkCurrentStatus)
		temp["BkProductName"] = common.GetString(host.Host.BkProductName)
		s := ""
		for i, k := range host.Module {
			if i == 0 {
				s = strings.Replace(common.GetString(k.TopModuleName), "##", ">", -1)
			} else {
				s = s + "    |    " + strings.Replace(common.GetString(k.TopModuleName), "##", ">", -1)
			}
		}
		temp["Module"] = s

		//var temp_t interface{}
		//temp_t = temp
		//temp_tt := temp_t.(map[string]string)
		hostIps[common.GetString(host.Host.BkHostName) + "/" + common.GetString(host.Host.BkHostInnerip)] = temp

		//hostIps = append(hostIps, temp)
	}
	return hostIps, len(host_total),nil
}