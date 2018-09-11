package qywechat

import (
	"testing"
	"fmt"
	"encoding/json"
)

var userApp = NewApp(34)

//测试数据
var testUser = User{
	Email:"testtest@juanpi.com",
	Name : "testtest.测试",
	Mobile : "",
	Userid : "testtest",
	Department : []int{296},
}

// 测试添加企业号成员
func TestCreateUser(t *testing.T) {
	j, _ := json.Marshal(testUser)
	fmt.Println(string(j))
	err := userApp.CreateUser(testUser)
	if err==nil{
		fmt.Println("添加成功")
	}else{
		fmt.Println("添加失败")
	}
}

// 测试更新企业号成员
func TestUpdateUser(t *testing.T) {
	testUser.Mobile = "222222222"
	err := userApp.UpdateUser(testUser)
	if err==nil{
		fmt.Println("更新成功")
	}else{
		fmt.Println("更新失败")
	}
}


// 测试删除企业号成员
func TestDeleteUser(t *testing.T) {
	err := userApp.DeleteUser(testUser.Userid)
	if err==nil{
		fmt.Println("删除成功")
	}else{
		fmt.Println("删除失败")
	}
}

// 测试获取企业号成员
func TestGetUser(t *testing.T) {
	user,err := userApp.GetUser(testUser.Userid)
	if err==nil{
		fmt.Println("用户信息：",user)
	}else{
		fmt.Println("获取失败")
	}
}