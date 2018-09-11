package api

import (
	"encoding/json"
	"fmt"
)

const BASE_ASSET_API_URL  = "/api/assets/v1"
const BASE_PERMS_API_URL = "/api/perms/v1"
type Node struct {
	Id 		string			`json:"id"`
	Key 		string			`json:"key"`
	Value		string			`json:"value"`
	Parent		string			`json:"parent"`
	AssetsAmount	float64			`json:"assets_amount"`
	IsNode		bool			`json:"is_node"`
}

type Asset struct {
	Id 		string			`json:"id"`
	Ip 		string			`json:"ip"`
	Hostname	string			`json:"hostname"`
	AdminUser       string			`json:"admin_user"`
}

type Assets struct {
	Assets 		[]string		`json:"assets"`
}

type AdminUser struct {
	Id 		string			`json:"id"`
	Name 		string			`json:"name"`
	Username	string			`json:"username"`
	BecomeUser	string			`json:"become_user"`
	BecomeMethod	string			`json:"become_method"`
	Become		bool			`json:"become"`
}

type SystemUser struct {
	Id 		string			`json:"id"`
	Name 		string			`json:"name"`
	Username	string			`json:"username"`
}

type Permissons struct {
	Id string `json:"id"`
	User []string 	`json:"user"`
	User_groups []string 	`json:"user_groups"`
	Assets []string 	`json:"assets"`
	Nodes []string `json:"nodes"`
	System_users []string `json:"system_users"`
	Inherit string `json:"inherit"`
	Name string `json:"name"`
	Is_active bool 	`json:"is_active"`
	Date_start string `json:"date_start"`
	Date_expired string `json:"date_expired"`
	Created_by string `json:"created_by"`
	Date_created string `json:"date_created"`
	Comment string 	`json:"comment"`
}

func (c *Api) GetNodeList() ([]Node, error) {
	c.SetRequestType(BASE_ASSET_API_URL + "/nodes/")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Node{}

	if nil != rstErr {
		fmt.Println("failed to search the NodeList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetNodeChildren(parentId string) ([]Node, error) {
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/children", parentId))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Node{}

	if nil != rstErr {
		fmt.Println("failed to search the NodeList, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetNodeAssets(nodeId string) ([]Asset, error) {
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/assets", nodeId))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Asset{}

	if nil != rstErr {
		fmt.Println("failed to search the NodeAssets, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateNode(data map[string]interface{}) (Node, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + "/nodes/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := Node{}

	if nil != rstErr {
		fmt.Println("failed to create the Node, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) UpdateNode(nodeId string, data map[string]interface{}) (Node, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/", nodeId))
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := Node{}

	if nil != rstErr {
		fmt.Println("failed to update the Node, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteNode(nodeId string) error {
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/", nodeId))
	_, rstErr := c.Send("DELETE", []byte{})
	if nil != rstErr {
		fmt.Println("failed to delete the Node , error info is ", rstErr.Error())
		return rstErr
	}
	return nil
}

func (c *Api) CreateAsset(data map[string]interface{}) (string, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + "/assets/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := Asset{}

	if nil != rstErr {
		fmt.Println("failed to create the asset, error info is ", rstErr.Error())
		return "", rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return "", jsErr
	}

	return rstObj.Id, nil
}

func (c *Api) SearchAssets(hostName, hostIp string) ([]Asset, error) {
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/assets/?hostname=%s&ip=%s", hostName, hostIp))
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Asset{}

	if nil != rstErr {
		fmt.Println("failed to search the Assets, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) NodesAssetsAdd(nodeId string, data map[string]interface{}) (Assets, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/assets/add/", nodeId))
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := Assets{}

	if nil != rstErr {
		fmt.Println("failed to add the NodeAssets, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) NodesAssetsRemove(nodeId string, data map[string]interface{}) (Assets, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/assets/remove/", nodeId))
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := Assets{}

	if nil != rstErr {
		fmt.Println("failed to add the NodeAssets, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetAdminUsers() ([]AdminUser, error) {
	c.SetRequestType(BASE_ASSET_API_URL + "/admin-user/")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []AdminUser{}

	if nil != rstErr {
		fmt.Println("failed to search the Admin-users, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateAdminUser(data map[string]interface{}) (AdminUser, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + "/admin-user/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := AdminUser{}

	if nil != rstErr {
		fmt.Println("failed to create the Admin-user, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) GetSystemUsers() ([]SystemUser, error) {
	c.SetRequestType(BASE_ASSET_API_URL + "/system-user/")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []SystemUser{}

	if nil != rstErr {
		fmt.Println("failed to search the System-users, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateSystemUser(data map[string]interface{}) (SystemUser, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + "/system-user/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := SystemUser{}

	if nil != rstErr {
		fmt.Println("failed to create the System-user, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) AddNodeChildren(nodeId string, data map[string]interface{}) error {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_ASSET_API_URL + fmt.Sprintf("/nodes/%s/children/add/", nodeId))
	_, rstErr := c.Send("PUT", dataStr)
	if nil != rstErr {
		fmt.Println("failed to add the Node childs, error info is ", rstErr.Error())
		return rstErr
	}
	return nil
}

func (c *Api) GetAssetPermissionsList() ([]Permissons, error) {
	c.SetRequestType(BASE_PERMS_API_URL + "/asset-permissions/")
	rst, rstErr := c.Send("GET", []byte{})
	rstObj := []Permissons{}

	if nil != rstErr {
		fmt.Println("failed to search the asset-permissions, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) CreateAssetPermissions(data map[string]interface{}) (Permissons, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_PERMS_API_URL + "/asset-permissions/")
	rst, rstErr := c.Send("POST", dataStr)
	rstObj := Permissons{}

	if nil != rstErr {
		fmt.Println("failed to create the asset-permissions, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) UpdateAssetPermissions(Id string, data map[string]interface{}) (Permissons, error) {
	dataStr, _ := json.Marshal(data)
	c.SetRequestType(BASE_PERMS_API_URL + fmt.Sprintf("/asset-permissions/%s/", Id))
	rst, rstErr := c.Send("PUT", dataStr)
	rstObj := Permissons{}

	if nil != rstErr {
		fmt.Println("failed to update the asset-permissions, error info is ", rstErr.Error())
		return rstObj, rstErr
	}

	if jsErr := json.Unmarshal([]byte(rst), &rstObj); nil != jsErr {
		fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
		return rstObj, jsErr
	}

	return rstObj, nil
}

func (c *Api) DeleteAssetPremissions(Id string) (error) {
	c.SetRequestType(BASE_PERMS_API_URL + fmt.Sprintf("/asset-permissions/%s/", Id))
	_, rstErr := c.Send("DELETE", []byte{})
	if nil != rstErr {
		fmt.Println("failed to delete the DeleteAssetPremissions , error info is ", rstErr.Error())
		return rstErr
	}
	return nil
}