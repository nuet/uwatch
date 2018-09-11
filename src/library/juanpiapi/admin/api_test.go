//卷皮后台相关接口
package admin

import "testing"

// 测试添加用户方法
func TestAddUserMes(t *testing.T) {
	userData := map[string]string{
		"username":  "衣然",
		"email":     "yiran@juanpi.com",
		"truename":  "胡金梅",
		"telephone": "13545163557",
		"password":  "juanpi.com",
		"groupid":   "3",
		"roleid":    "2151",
	}
	err := AddUserMes(userData)
	t.Log(err)
}

// 测试更新用户
func TestUpdateUserMes(t *testing.T) {
	userData := map[string]string{
		"username": "test",
		"password": "123456",
		"groupid":  "1",
		"roleid":   "1",
	}
	err := UpdateUserMes(userData)
	t.Log(err)
}
