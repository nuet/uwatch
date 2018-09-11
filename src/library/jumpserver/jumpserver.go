package jumpserver

import (
	"library/jumpserver/api"
	"strings"
	"github.com/astaxie/beego"
)

type Nodes struct {
	api.Node
	Assets map[string]api.Asset
}

func GetAllNodes() (map[string]Nodes, error) {
	ret := make(map[string]Nodes)
	assetApi := api.New()
	nodes, err := assetApi.GetNodeList()
	if err != nil {
		return ret, err
	}

	for _, node := range nodes {
		value := strings.Split(node.Value, "-")
		key := value[0]
		if _, ok := ret[key]; !ok {
			nodeAssets := make(map[string]api.Asset)
			assets, err := assetApi.GetNodeAssets(node.Id)
			if err != nil {
				return ret, err
			}
			if len(assets) > 0 {
				for _, asset := range assets {
					if _, ok := nodeAssets[asset.Ip]; !ok {
						nodeAssets[asset.Ip] = asset
					}
				}
			}

			nodes := Nodes{node, nodeAssets}
			ret[key] = nodes
		}
	}

	return ret, nil
}

func GetFirstNodes() (map[string]api.Node, error){
	ret := make(map[string]api.Node)
	assetApi := api.New()
	nodes, err := assetApi.GetNodeList()
	if err != nil {
		return ret, err
	}
	var rootId string
	for _, node := range nodes {
		if node.Value == "ROOT" {
			rootId = node.Id
			break
		}
	}

	fNodes, err := assetApi.GetNodeChildren(rootId)
	if err != nil {
		return ret, err
	}

	for _, fnode := range fNodes {
		if fnode.Value != "ROOT" {
			fname := fnode.Value
			list := strings.Split(fname, "-")
			if len(list) > 1 {
				key := list[0]
				_, ok := ret[key]
				if !ok {
					ret[key] = fnode
				}
			}
		}
	}

	return ret, nil
}

func GetChildrenNodes(parentId string) (map[string]api.Node, error){
	ret := make(map[string]api.Node)
	assetApi := api.New()

	fNodes, err := assetApi.GetNodeChildren(parentId)
	if err != nil {
		return ret, err
	}

	for _, fnode := range fNodes {
		_, ok := ret[fnode.Key]
		if !ok {
			ret[fnode.Key] = fnode
		}
	}

	return ret, nil
}

func GetAdminUsers() ([]api.AdminUser, error) {
	assetApi := api.New()

	users, err := assetApi.GetAdminUsers()
	return users, err
}

func CreateAdminUser(name, username, password string) (api.AdminUser, error) {
	param := map[string]interface{}{
		"name" : name,
		"username" : username,
	}
	assetApi := api.New()
	user, err := assetApi.CreateAdminUser(param)
	return user, err
}

func GetSystemUsers() ([]api.SystemUser, error) {
	assetApi := api.New()

	users, err := assetApi.GetSystemUsers()
	return users, err
}

func CreateSystemUser(name, username string) (api.SystemUser, error) {
	param := map[string]interface{}{
		"name" : name,
		"username" : username,
	}
	assetApi := api.New()
	user, err := assetApi.CreateSystemUser(param)
	return user, err
}

func CreateNode(value string) (Nodes, error) {
	param := map[string]interface{}{
		"value" : value,
	}

	assetApi := api.New()

	node, err := assetApi.CreateNode(param)
	nodeAssets := make(map[string]api.Asset)
	return Nodes{node, nodeAssets}, err
}

func UpdateNode(value string, nodesId string) (Nodes, error) {
	param := map[string]interface{}{
		"id" : nodesId,
		"value" : value,
	}

	assetApi := api.New()

	node, err := assetApi.UpdateNode(nodesId, param)
	nodeAssets := make(map[string]api.Asset)
	return Nodes{node, nodeAssets}, err
}

func DeleteNode(nodeId string) error {
	assetApi := api.New()
	err := assetApi.DeleteNode(nodeId)
	return err
}

func GetAssetInfo(hostName, hostIp string) ([]api.Asset, error) {
	assetApi := api.New()

	assets, err := assetApi.SearchAssets(hostName, hostIp)
	return assets, err
}

func CreateAsset(hostName, hostIp, adminUser string) (string, error) {
	param := map[string]interface{}{
		"hostname" : hostName,
		"ip" : hostIp,
		"admin_user" : adminUser,
	}
	assetApi := api.New()

	id, err := assetApi.CreateAsset(param)
	return id, err
}

func NodesAssetsAdd(nodeId string, assets []string) (api.Assets, error) {
	param := map[string]interface{}{
		"assets" : assets,
	}

	assetApi := api.New()
	ret, err := assetApi.NodesAssetsAdd(nodeId, param)
	return ret, err
}

func NodesAssetsRemove(nodeId string, assets []string) (api.Assets, error) {
	param := map[string]interface{}{
		"assets" : assets,
	}

	assetApi := api.New()
	ret, err := assetApi.NodesAssetsRemove(nodeId, param)
	return ret, err
}

func AddNodeChildren(nodeId string, nodes []string) error {
	param := map[string]interface{}{
		"nodes" : nodes,
	}

	assetApi := api.New()
	err := assetApi.AddNodeChildren(nodeId, param)
	return err
}

func GetAllUsers() ([]api.Users, error) {
	usersApi := api.New()
	ret, err := usersApi.GetUsersList()
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func GetAllUserList() (map[string]api.Users, error){
	ret := make(map[string]api.Users)
	usersApi := api.New()
	users, err := usersApi.GetUsersList()
	if err != nil {
		return ret, err
	}

	for _, user := range users {
		_, ok := ret[user.Username]
		if !ok {
			ret[user.Username] = user
		}
	}

	return ret, nil
}

func GetUsersName(username string) ([]api.Users, error) {
	usersApi := api.New()
	user, err := usersApi.SearchUsers(username)
	return user, err
}


func CreateUser(cnname string, username string, email string, phone string, comment string, wechat string, groups_display string, groups []string) (api.Users, error) {
	param := map[string]interface{}{
		"groups_display" : groups_display,
		"groups" : groups,
		"name" : cnname,
		"username" : username,
		"email" : email,
		"comment": comment,
		"phone": phone,
		"wechat": wechat,
	}
	usersApi := api.New()
	user, err := usersApi.CreateUsers(param)
	return user, err
}


func UpdateUser(Id string, cnname string, username string, email string, phone string, comment string, wechat string, groups_display string, groups []string) (api.Users, error) {
	param := map[string]interface{}{
		"id" : Id,
		"groups_display" : groups_display,
		"groups" : groups,
		"name" : cnname,
		"username" : username,
		"comment": comment,
		"phone": phone,
		"wechat": wechat,
		"email" : email,
	}
	usersApi := api.New()
	user, err := usersApi.UpdateUsers(Id, param)
	return user, err
}

func DeleteUser(userId string) error {
	assetApi := api.New()
	err := assetApi.DeleteUsers(userId)
	return err
}


func GetGroupsName(groupname string) ([]api.Groups, error) {
	groupsApi := api.New()
	groups, err := groupsApi.SearchGroups(groupname)
	return groups, err
}

func GetGroupsList() ([]api.Groups, error) {
	groupsApi := api.New()
	groups, err := groupsApi.GetGroupsList()
	return groups, err
}

func GetAllGroups() (map[string]api.Groups, error){
	ret := make(map[string]api.Groups)
	groupsApi := api.New()
	userGroups, err := groupsApi.GetGroupsList()
	if err != nil {
		return ret, err
	}

	for _, userGroup := range userGroups {
		_, ok := ret[userGroup.Name]
		if !ok {
			ret[userGroup.Name] = userGroup
		}
	}

	return ret, nil
}

func CreateGroup(name string, comment string, users []string) (api.Groups, error) {
	param := map[string]interface{}{
		"name" : name,
		"comment" : comment,
		"users" : users,
	}
	usersApi := api.New()
	user, err := usersApi.CreateGroups(param)
	return user, err
}

func UpdateGroup(Id string,name string, comment string, users []string) (api.Groups, error) {
	param := map[string]interface{}{
		"name" : name,
		"comment" : comment,
		"users" : users,
	}
	usersApi := api.New()

	user, err := usersApi.UpdateGroups(Id,param)
	return user, err
}

func DeleteGroup(groupId string) error {
	assetApi := api.New()
	err := assetApi.DeleteGroups(groupId)
	return err
}

func GetAssetPermissions() (map[string]api.Permissons, error){
	ret := make(map[string]api.Permissons)
	perApi := api.New()
	AssetPermissions, err := perApi.GetAssetPermissionsList()
	if err != nil {
		return ret, err
	}

	for _, asset := range AssetPermissions {
		_, ok := ret[asset.Name]
		if !ok {
			ret[asset.Name] = asset
		}
	}

	return ret, nil
}

func CreateAssetPermissions(name string, users []string, assets []string, system_user []string, is_active bool, comment string) (api.Permissons, error){
		param := map[string]interface{}{
			"name" : name,
			"comment" : comment,
			"system_users" : system_user,
			"users" : users,
			//"user_groups" : temp_group,
			"assets" : assets,
			//"nodes" : temp_node,
			"is_active" : is_active,
		}
		assetPreApi := api.New()
		pre, err := assetPreApi.CreateAssetPermissions(param)
		return pre, err
}

func CreateAssetPermission(name string, temp_group []string, temp_node []string, system_user []string, is_active bool, comment string) (api.Permissons, error){
	param := map[string]interface{}{
		"name" : name,
		"comment" : comment,
		"system_users" : system_user,
		"user_groups" : temp_group,
		"nodes" : temp_node,
		"is_active" : is_active,
	}
	assetPreApi := api.New()
	pre, err := assetPreApi.CreateAssetPermissions(param)
	return pre, err
}

func UpdateAssetPermissions(id string, name string, temp_group []string, temp_node []string, system_user []string, is_active bool, comment string) (api.Permissons, error){
	param := map[string]interface{}{
		"name" : name,
		"comment" : comment,
		"system_users" : system_user,
		"user_groups" : temp_group,
		"nodes" : temp_node,
		"is_active" : is_active,
	}
	assetPreApi := api.New()
	pre, err := assetPreApi.UpdateAssetPermissions(id,param)
	return pre, err
}

func UpdateAssetPermission(id string,name string, users []string,  assets []string, system_user []string, is_active bool, comment string) (api.Permissons, error){
	param := map[string]interface{}{
		"name" : name,
		"comment" : comment,
		"users" : users,
		"assets" : assets,
		"system_users" : system_user,
		//"user_groups" : temp_group,
		//"nodes" : temp_node,
		"is_active" : is_active,
	}
	beego.Info("UpdateAssetPermission===>", param)
	assetPreApi := api.New()
	pre, err := assetPreApi.UpdateAssetPermissions(id,param)
	return pre, err
}

func DeleteAssetPermissions(id string) error {
	assetPreApi := api.New()
	err := assetPreApi.DeleteAssetPremissions(id)
	return err
}

func GetNodeAssets() (map[string]api.Asset, error) {
	ret := make(map[string]api.Asset)
	nodeAssets := make(map[string]api.Asset)
	assetApi := api.New()
	nodes, err := assetApi.GetNodeList()
	if err != nil {
		return ret, err
	}

	for _, node := range nodes {
		value := strings.Split(node.Value, "-")
		key := value[0]
		if _, ok := ret[key]; !ok {
			assets, err := assetApi.GetNodeAssets(node.Id)
			if err != nil {
				return ret, err
			}
			if len(assets) > 0 {
				for _, asset := range assets {
					if _, ok := nodeAssets[asset.Ip]; !ok {
						nodeAssets[node.Id] = asset
					}
				}
			}
		}
	}
	return nodeAssets, nil
}