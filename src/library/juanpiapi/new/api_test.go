package new

import (
	"fmt"
	"testing"
)

func TestGetStation(t *testing.T) {
	result, err := GetStations()
	if err != nil {
		print(err)
	}

	t.Logf("%v", result)
}

type user struct {
	username string
	userid   string
}

// 更新用户状态
func TestUpdateUserMes(t *testing.T) {
	var users = []user{{"超哥", "254"}, {"熊猫", "1040"}, {"黛西", "1084"}, {"小念", "420"}, {"老振", "143"}, {"叮当", "1357"}, {"橘子", "1358"}, {"寒冰", "1383"}, {"无双", "339"}, {"西瓜", "1422"}, {"小苏", "1468"}, {"暖暖", "1551"}, {"娟子", "1614"}, {"欧克", "927"}, {"雪莉", "1488"}, {"小丁", "2614"}, {"李子", "2707"}, {"蕾仔", "161"}}

	// 黛西_1084

	updateData := make(map[string]string)
	updateData["status"] = "1"
	for _, user := range users {
		updateData["username"] = fmt.Sprintf("%s_%s", user.username, user.userid)
		updateData["username_id"] = user.userid
		err := UpdateUserMes(updateData)
		if err != nil {
			t.Log(fmt.Sprintf("%s_%s", user.username, user.userid))
			t.Log(err)
		}
	}
}
