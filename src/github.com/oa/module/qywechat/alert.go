package qywechat

import (
	"fmt"
	"time"
)

type alertInfo struct {
	users    string    // 报警接收人
	alertNum int       // 报警数量
	content  string    // 报警内容
	update   time.Time // 最后一条报警的时间
	send     time.Time // 最后一次报警发送时间
	sendNum  int       // 报警发送次数
}

const (
	SEND = iota
	STORE
	EXPIRED
)

// 获取告警状态
func (alert *alertInfo) getStatus() int {
	now := time.Now()
	updateSub := int(now.Sub(alert.update).Seconds())
	println("update:", updateSub)
	if updateSub > 3600 {
		return EXPIRED
	}

	sendSub := int(now.Sub(alert.send).Seconds())
	println("send:", sendSub)
	if sendSub >= ((alert.sendNum) * 3) {
		return SEND
	}

	return STORE
}

// 发送告警
func (alert *alertInfo) sendMsg() {
	alertApp.SendText(alert.users, fmt.Sprintf("%s 【告警数：%d】", alert.content, alert.alertNum))
	alert.send = time.Now()
	alert.sendNum += 1
}

var (
	alertLogs = make(map[string]*alertInfo)
	alertApp  = NewApp(34)
)

/**
 * 发送告警信息
 */
func Alert(users, name, content string) {
	// 首先判断该类型的报警是否已经发出过
	alert, ok := alertLogs[name]
	if !ok {
		alert = &alertInfo{}
		alertLogs[name] = alert
	}

	alert.users = users
	alert.alertNum += 1
	alert.content = content
	alert.update = time.Now()

	for name, alert := range alertLogs {
		switch alert.getStatus() {
		case SEND:
			alert.sendMsg()
		case EXPIRED:
			delete(alertLogs, name)
		}
	}
}
